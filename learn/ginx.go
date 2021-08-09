package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

type ttt struct {
	Name []string `json:"name"`
}

func main() {
	// 创建一个默认的路由引擎
	r := gin.Default()
	// GET：请求方式；/hello：请求的路径
	// 当客户端以GET方法请求/hello路径时，会执行后面的匿名函数
	r.POST("/hello", func(c *gin.Context) {
		var tt ttt
		bb := c.Bind(&tt)
		fmt.Println(bb)
		fmt.Println(tt)
		// c.JSON：返回JSON格式的数据
		c.JSON(200, gin.H{
			"message": "Hello world!",
		})
	})
	// 启动HTTP服务，默认在0.0.0.0:8080启动服务
	r.Run()
}
