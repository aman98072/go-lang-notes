## Go Configurations

## Basic Commands 
- go mod init github.com/aman98072/crudAPI : first create init path like root path
- go get -u github.com/gorilla/mux  : this is for routing
- go get -u github.com/jinzhu/gorm  : for gorm 
- go get -u gorm.io/gorm            : this is ORM for database query  
- go get -u gorm.io/driver/mysql    : this is using for mysql driver  

## create users table 
- create table users ( id INT NOT NULL AUTO_INCREMENT, fname VARCHAR(30) NOT NULL, lname VARCHAR(30) NOT NULL, date DATE, PRIMARY KEY ( id ) );

## Print query for debugging
- db.Debug().Where("ID = ?", Id).Delete(user) 

## Run raw query 
- db.Exec("DELETE FROM users where id = 2")

## Run project command 
- go run cmd/main/main.go