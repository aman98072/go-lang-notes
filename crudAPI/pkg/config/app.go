package config

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var (
	db *gorm.DB
)

func Connect() {
	dbDriver := "mysql"
	dbUser := "aman"
	dbPass := "Aman@98072"
	dbName := "go_crud"
	query := "?charset=utf8mb4&parseTime=True&loc=Local"
	d, err := gorm.Open(dbDriver, dbUser+":"+dbPass+"@/"+dbName+query)
	if err != nil {
		panic(err.Error())
	}
	db = d
}

func GetDB() *gorm.DB {
	return db
}
