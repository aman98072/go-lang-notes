package models

import (
	"fmt"

	"github.com/aman98072/crudAPI/pkg/config"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB

type User struct {
	gorm.Model
	Fname string `gorm:""json:"fname"`
	Lname string `json:"lname"`
}

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&User{})
}

func (b *User) CreateUser() *User {
	db.NewRecord(b)
	db.Create(&b)
	return b
}

func GetAllUsers() []User {
	var Users []User
	db.Find(&Users)
	return Users
}

func GetUserById(Id int64) (*User, *gorm.DB) {
	var getUser User
	db := db.First(&getUser, Id)
	return &getUser, db
}

func DeleteUser(Id int64) User {
	var user User
	fmt.Println("id is : ", Id)
	// db.Delete(&user{}, Id)
	// db.Delete(&User{}, 10)
	// DELETE FROM users WHERE id = 10;

	// db.Delete(user, Id)
	// db.Debug().Delete(user, Id)
	// db.Delete(Id)
	db.Where("ID = ?", Id).Delete(User{})
	db.Debug().Where("ID = ?", Id).Delete(User{})
	// db.Debug().Where("ID = ?", Id).Delete(user) : for print query for debugging

	// db.Exec("DELETE FROM users where id = 2") : run raw query
	return user
}
