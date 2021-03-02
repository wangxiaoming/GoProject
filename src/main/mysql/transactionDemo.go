package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type Persons struct {
	UserId   int    `gorm:"primary_key"`
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

	db.Set("gorm:table_options", "ENGINE-InnoDB").AutoMigrate(&Persons{})
	// 开启事务
	tx := db.Begin()
	// 插入数据
	var person1 Persons
	person1 = Persons{110520219, "xiaoming", "Feman", "123@qq.com"}
	r1 := tx.Create(&person1)
	var person2 Persons
	person2 = Persons{110520220, "xiaohua", "MAN", "123323@qq.com"}
	r2 := tx.Create(&person2)

	// 两个有一个异常，则回滚
	if r1.Error != nil || r2.Error != nil {
		fmt.Println("出现系统异常", r1.Error, r2.Error)
		tx.Rollback()
	} else {
		tx.Commit()
		fmt.Println("系统提交")
	}

}
