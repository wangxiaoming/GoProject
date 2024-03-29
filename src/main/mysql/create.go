package main

import (
	"database/sql"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"time"
)

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

type User struct {
	gorm.Model
	UserId      int64 `gorm:"index"`
	Birtheday   time.Time
	Age         int           `gorm:"column:age"`
	Name        string        `gorm:"size:255;index:idx_name_add_id"`
	Num         int           `gorm:"AUTO_INCREMENT"`
	Email       string        `gorm:"type:varchar(100);unique_index"`
	AddressID   sql.NullInt64 `gorm:"index:idx_name_add_id"`
	IgnoreMe    int           `gorm:"_"`
	Description string        `gorm:"size:2019;comment:'用户描述字段'"`
	Status      string        `gorm:"type:enum('published', 'pending', 'deleted');default:'pending'"`
}

//设置表名，默认是结构体的名的复数形式
func (User) TableName() string {
	return "VIP_USER"
}

func main() {

	defer db.Close()

	if db.HasTable(&User{}) { //判断表是否存在
		db.AutoMigrate(&User{}) //存在就自动适配表，也就说原先没字段的就增加字段
	} else {
		db.CreateTable(&User{}) //不存在就创建新表
	}
}
