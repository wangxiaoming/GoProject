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

	///根据一个条件更新
	//根据条件更新字段值,
	//后面加Debug()，运行时，可以打印出sql
	db.Debug().Model(&Animal{}).Where("id = ? ", 4).Update("name", "jimupdate")
	//UPDATE `animals` SET `name` = 'jimupdate'  WHERE (id = 4)

	//另外一种写法： 根据条件更新
	var animal Animal
	animal = Animal{ID: 3}
	db.Debug().Model(animal).Update("name", "demotest2update")
	// db.Debug().Model(&animal).Update("name", "demotest2update") // 这种写法也可以
	//UPDATE `animals` SET `name` = 'demotest2update'  WHERE `animals`.`id` = 3

	/// 多个条件更新
	db.Model(&Animal{}).Where("id = ? AND age = ?", 4, 45).Update("name", "jimupdate3")
	//UPDATE `animals` SET `name` = 'jimupdate2'  WHERE (id = 4 AND age = 45)

	/// 更新多个值
	db.Debug().Model(&Animal{}).Where("id = ?", 4).Update(Animal{Name: "jim", Age: 90})
	// UPDATE `animals` SET `age` = 90, `name` = 'jim'  WHERE (id = 4)

	animal2 := Animal{ID: 5}
	db.Debug().Model(&animal2).Update(map[string]interface{}{"name": "jimm", "age": 100})
	//UPDATE `animals` SET `age` = 100, `name` = 'jimm'  WHERE `animals`.`id` = 5
}
