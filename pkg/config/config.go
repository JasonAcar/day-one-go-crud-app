package config

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/joho/godotenv/autoload"
	"os"
)

var (
	db *gorm.DB
)

func Connect() {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true&charset=utf8&loc=Local", os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_NAME"))
	// akhil:Axlesharma@12@/simplerest?charset=utf8&parseTime=True&loc=Local
	d, err := gorm.Open("mysql", dsn)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	db = d
}

func GetDB() *gorm.DB {
	return db
}
