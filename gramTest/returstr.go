package main

import (
	"github.com/sirupsen/logrus"
	"os"
	"runtime"
)

func main() {

}

type MongodbInfo struct {
	MongodbCluster  string
	MongodbUsername string
	MongodbPassword string
}

func GetMongoInfo() MongodbInfo {
	var mongodbInfo MongodbInfo
	mongodbInfo.MongodbCluster = os.Getenv("mongodbCluster")
	mongodbInfo.MongodbUsername = os.Getenv("mongodbUsername")
	mongodbInfo.MongodbPassword = os.Getenv("mongodbPassword")

	if len(mongodbInfo.MongodbCluster) == 0 {
		if runtime.GOOS == "windows" || runtime.GOOS == "darwin" {
			logrus.WithField("topic", "初始化").Warn("没有获取到mongodbCluster 环境变量,采用本地配置变量")
			mongodbInfo.MongodbCluster = "1"
		} else {
			logrus.WithField("topic", "初始化").Fatal("没有获取到mongodbCluster 环境变量,结束程序")
		}
	}
	if len(mongodbInfo.MongodbUsername) == 0 {
		if runtime.GOOS == "windows" || runtime.GOOS == "darwin" {
			logrus.WithField("topic", "初始化").Warn("没有获取到mongodbUsername 环境变量,采用本地配置变量")
			mongodbInfo.MongodbUsername = "2"
		} else {
			logrus.WithField("topic", "初始化").Fatal("没有获取到mongodbUsername 环境变量,结束程序")
		}
	}
	if len(mongodbInfo.MongodbPassword) == 0 {
		if runtime.GOOS == "windows" || runtime.GOOS == "darwin" {
			logrus.WithField("topic", "初始化").Warn("没有获取到mongodbPassword 环境变量,采用本地配置变量")
			mongodbInfo.MongodbPassword = "3"
		} else {
			logrus.WithField("topic", "初始化").Fatal("没有获取到mongodbPassword 环境变量,结束程序")
		}
	}
	return mongodbInfo
}
