package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type Person struct {
	ID    int    `gorm:"primary_key"`
	Code  string `gorm:"type:varchar(20);"`
	Price int    `gorm:type:in`
	Name  string `gorm:type:varchar(64)`
}

var db *gorm.DB

func main() {

	db, err := gorm.Open("mysql", "root:root@tcp(127.0.0.1:3306)/test?charset=utf8")
	if err != nil {
		fmt.Println("连接数据库失败", err)
	}
	fmt.Println("连接数据库成功")
	var person Person
	// 创建表
	db.CreateTable(&Person{})
	// 插入数据
	person = Person{Code: " 111111", Price: 1100}
	db.Create(&person)

	// 读取数据
	var person1 Person
	//db.First(&person1, 2) // 查询 id 为 1的 Product
	//fmt.Println(person1.Price, person1.Code, person1.Name)
	db.First(&person1, "Name = ?", "xiaohua")
	fmt.Println(person1.Price, person1.Code, person1.Name)

	// 更新数据
	db.Model(&person1).Update("Price", 2000)
	db.Model(&person).Update("Name", "xiaohong")

	// 删除 person
	//db.Delete(&person)
}
