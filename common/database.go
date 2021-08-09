package common

import (
	"ginEssential/Model"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"net/url"

	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() *gorm.DB {

	host := viper.GetString("datasource.host")
	port := viper.GetString("datasource.port")
	database := viper.GetString("datasource.database")
	username := viper.GetString("datasource.username")
	password := viper.GetString("datasource.password")
	charset := viper.GetString("datasource.charset")
	loc := viper.GetString("detasource.loc")

	dsn := username + ":" + password + "@tcp" + "(" + host + ":" + port + ")" + "/" + database + "?" + "charset=" + charset + "&parseTime=True&loc=" + url.QueryEscape(loc)

	//dsn := "root:LondonVic222@tcp(127.0.0.1:3306)/ginessential?charset=utf8&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	db.AutoMigrate(&Model.User{})
	DB = db

	return db
}

func GetDB() *gorm.DB {
	return DB
}
