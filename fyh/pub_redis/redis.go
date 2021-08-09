package pub_redis

import (
	"context"
	"fmt"
	"github.com/gomodule/redigo/redis"
	"log"
	"time"
)

// 创建连接池
type RedisPool struct {
	pool *redis.Pool
}

type RedisConf struct {
	Host       string
	Port       string
	Password   string
	DB         int
	PoolConfig *PoolConfig
}

type PoolConfig struct {
	MaxIdle int
	// Maximum number of connections allocated by the pool at a given time.
	// When zero, there is no limit on the number of connections in the pool.
	MaxActive int

	// Close connections after remaining idle for this duration. If the value
	// is zero, then idle connections are not closed. Applications should set
	// the timeout to a value less than the server's timeout.
	IdleTimeout time.Duration

	// If Wait is true and the pool is at the MaxActive limit, then Get() waits
	// for a connection to be returned to the pool before returning.
	Wait bool

	// Close connections older than this duration. If the value is zero, then
	// the pool does not close connections based on age.
	MaxConnLifetime time.Duration
}

func dial(conf *RedisConf) *redis.Pool {
	addr := fmt.Sprintf("%s:%s", conf.Host, conf.Port)

	p := &PoolConfig{
		MaxIdle:     10,
		MaxActive:   100,
		IdleTimeout: 10 * time.Second,
	}
	if conf.PoolConfig != nil {
		p = conf.PoolConfig
	}
	return &redis.Pool{
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", addr,
				redis.DialPassword(conf.Password),
				redis.DialDatabase(conf.DB))
			if err != nil {
				return nil, err
			}
			return c, nil
		},
		MaxIdle:         p.MaxIdle,
		MaxActive:       p.MaxActive,
		IdleTimeout:     p.IdleTimeout,
		Wait:            p.Wait,
		MaxConnLifetime: p.MaxConnLifetime,
	}
}

func (rp *RedisPool) GetConn(ctx context.Context) redis.Conn {
	return rp.pool.Get()
}

func (rp *RedisPool) do(ctx context.Context, commandName string, args ...interface{}) (reply interface{}, err error) {
	conn := rp.GetConn(ctx)
	defer conn.Close()

	start := time.Now()
	reply, err = conn.Do(commandName, args...)
	if err != nil {
		log.Print(
			map[string]interface{}{
				"command": commandName,
				"args":    args,
				"since":   time.Since(start).Milliseconds(),
				"error":   err,
			})
		return reply, err
	}
	log.Print(
		map[string]interface{}{
			"command": commandName,
			"args":    args,
			"since":   time.Since(start).Milliseconds(),
		})
	return reply, err
}

func (rp *RedisPool) Do(ctx context.Context, command string, filed ...interface{}) (interface{}, error) {
	reply, err := rp.do(ctx, command, filed...)
	return reply, err
}

func NewRedisPool(conf *RedisConf) *RedisPool {
	if conf == nil {
		conf = &RedisConf{
			Host:     "127.0.0.1",
			Port:     "6379",
			Password: "",
			DB:       0,
		}
	}
	pool := &RedisPool{dial(conf)}
	r, err := pool.Do(context.Background(), "ping")
	if err != nil {
		panic(err)
	}

	respString, ok := r.(string)
	if respString != "PONG" || !ok {
		panic("InitRedis authRedis err" + respString)
	}
	return pool
}

func (rp *RedisPool) Del(ctx context.Context, key string) error {
	_, err := rp.do(ctx, "DEL", key)
	return err
}

func (rp *RedisPool) Get(ctx context.Context, key string) (string, error) {
	return redis.String(rp.do(ctx, "GET", key))
}
func (rp *RedisPool) Set(ctx context.Context, key string, val interface{}) (string, error) {
	return redis.String(rp.do(ctx, "SET", key, val))
}

func (rp *RedisPool) SetNx(ctx context.Context, key string, val interface{}) (int, error) {
	return redis.Int(rp.do(ctx, "SETNX", key, val))
}

func (rp *RedisPool) SetEx(ctx context.Context, key string, second int, val interface{}) (string, error) {
	return redis.String(rp.do(ctx, "SETEX", key, second, val))
}

func (rp *RedisPool) Expire(ctx context.Context, key string, seconds int) error {
	_, err := rp.do(ctx, "EXPIRE", key, seconds)
	return err
}
