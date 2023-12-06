package utils

import (
	"fmt"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"manageSystem/config"
	"manageSystem/model"
)

var (
	DB *gorm.DB
)

func InitViper() {
	if err := config.Init(""); err != nil {
		panic(err)
	}
}

func InitDB() {
	fmt.Println("数据库 init")
	var err error
	conf := &model.Database{
		Host:     viper.GetString("database.host"),
		Username: viper.GetString("database.username"),
		Password: viper.GetString("database.password"),
		DBName:   viper.GetString("database.db_name"),
	}

	config := fmt.Sprintf("%s:%s@tcp(%s)/%s?parseTime=true&charset=utf8&parseTime=%t&loc=%s",
		conf.Username,
		conf.Password,
		conf.Host,
		conf.DBName,
		true,
		"Local")

	DB, err = gorm.Open(mysql.Open(config), &gorm.Config{})
	if err != nil {
		log.Fatalf("connect error: %v\n", err)
	}
	fmt.Println("数据库 init 结束...")
}
