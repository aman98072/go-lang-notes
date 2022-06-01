package config

import (
	"github.com/jinzhu/gorm"
	// "github.com/jinzhu/gorm/dialects/mysql"  //  ask
)

var (
	db *gorm.DB
)

func Connect() {
	d, err := gorm.Open("mysql", "root:profsnapeisdon@tcp(dev1.droom.in)/cscart_new?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}

	db = d
}

func GetDB() *gorm.DB {
	return db
}
