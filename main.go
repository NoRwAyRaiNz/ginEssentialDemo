package main

import (
	"ginEssential/common"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"os"
)

func main() {
	InitConfig()

	db := common.InitDB()
	db = db

	r := gin.Default()
	r = CollectRoute(r)
	//panic(r.Run()) // listen and serve on 0.0.0.0:8080
	//监听端口
	port := viper.GetString("server.port")

	if port != "" {
		panic(r.Run(":" + port))
	}

}

func InitConfig() {
	//获取工作目录
	workDir, _ := os.Getwd()
	//设置要读取的配置文件名
	viper.SetConfigName("application")
	//设置要读取的配置文件类型
	viper.SetConfigType("yml")
	//读取要配置的文件路径
	viper.AddConfigPath(workDir + "/config")
	//
	err := viper.ReadInConfig()
	if err != nil {

	}

}

//func InitDB() *gorm.DB {
//	dirverName := "mysql"
//	host := "localhost"
//	port := "3306"
//	database := "ginessential"
//	username := "root"
//	password := "root"
//	charset := "utf-8"
//	args := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime = true",
//		username,
//		password,
//		host,
//		port,
//		database,
//		charset)
//
//}
