package db

import (
	"fmt"
	"gorm.io/gorm"
)
import "gorm.io/driver/mysql"

var Client *gorm.DB

func init() {
	Client, err := gorm.Open(mysql.Open("root:ltinyho@tcp(localhost:3306)/ltinyho?charset=utf8&parseTime=True&loc=Local"), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	fmt.Println(Client.Name())
}

func Db()*gorm.DB{
	return Client
}
