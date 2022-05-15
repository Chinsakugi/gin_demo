package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Student struct {
	Id     int `gorm:"primarykey"`
	Name   string
	Number string
	Age    int
}

type Author struct {
	Name  string
	Email string
}

type Blog struct {
	ID      int
	Author  Author `gorm:"embedded"`
	Upvotes int
	Deleted gorm.DeletedAt
}

type Result struct {
	Name  string
	Total int
}

func main() {
	db, err := gorm.Open(mysql.Open("root:123456@tcp(127.0.0.1:3306)/test?charset=utf8mb4&parseTime=True&loc=Local"), &gorm.Config{})
	if err != nil {
		fmt.Println("database connect fail")
	}

	if err1 := db.AutoMigrate(&Student{}); err1 != nil {
		fmt.Println("database autoMigrate error")
	}

	db.AutoMigrate(&Student{})

	//stu := Student{
	//	Name:   "czy",
	//	Number: "10185101283",
	//	Age:    23,
	//}
	//db.Create(&stu)
	//fmt.Println(stu)

	stu := Student{}
	db.First(&stu, "age = ?", 23)
	fmt.Println(stu)

	stus := []Student{}
	db.Find(&stus)
	fmt.Println(stus)

	res1 := Student{}
	db.Where("id = ?", 4).First(&res1)
	fmt.Println(res1)

	res2 := []Result{}
	db.Model(&Student{}).Select("name, sum(age) as total").Group("name").Find(&res2)
	fmt.Println(res2)
}
