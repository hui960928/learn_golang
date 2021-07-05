package resource

import (
	"context"
	"github.com/qiniu/qmgo"
	"github.com/sirupsen/logrus"
)

var (
	dbClient       *qmgo.Client
	GetMongoClient = func() *qmgo.Client {
		if dbClient == nil {
			initMongo()
		}
		return dbClient
	}
)

//初始化mongo数据库
func initMongo() {
	ctx := context.Background()
	client, err := qmgo.NewClient(ctx, &qmgo.Config{Uri: "mongodb://" + "Hik706706" + ":" + "Alibaba&Cetiti" + "@" + "10.0.40.30" + ":" + "20000"})
	//client, err := qmgo.NewClient(ctx, &qmgo.Config{Uri: "mongodb://" + MongoInfo.Username + ":" + MongoInfo.Password + "@" + MongoInfo.Cluster})
	if err != nil {
		logrus.WithField("topic", "初始化").Info("mongodb初始化失败!")
		return
	}
	logrus.WithField("topic", "初始化").Info("mongodb初始化成功!")
	dbClient = client
}

func init() {
	GetMongoClient()
}
