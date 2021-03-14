package main

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type Animal struct {
	ID   int64
	Name string
	Age  int64
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

	defer db.Close()

	db.Debug().Where("id = ?", 13).Delete(&Animal{})
	// DELETE FROM `animals`  WHERE (id = 13)

	db.Debug().Delete(&Animal{}, "id = ? AND age = ?", 14, 10)
	//DELETE FROM `animals`  WHERE (id = 14 AND age = 10)

}
