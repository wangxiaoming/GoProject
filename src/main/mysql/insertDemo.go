package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type Persons struct {
	UserId   int    `db:"user_id"`
	UserName string `db:"username"`
	Sex      string `db:"sex"`
	Email    string `db:"email"`
}

var db *gorm.DB

func init() {
	database, err := gorm.Open("mysql", "root:root@tcp(127.0.0.1:3306)/test?charset=utf8")
	if err != nil {
		fmt.Println("连接数据库失败", err)
		return
	}
	db = database
	//defer database.Close()
}

func main() {
	// 创建表
	r1 := db.CreateTable(&Persons{})

	if r1.Error != nil {
		fmt.Println("建表失败", r1.Error)
	} else {
		fmt.Println("建表成功")
	}

	var person Persons
	person = Persons{UserId: 1105202018, UserName: "xiaoming", Sex: "MAN", Email: "123456@qq.com"}
	// 插入数据
	db.Create(&person)
	fmt.Println("插入数据成功")

}
