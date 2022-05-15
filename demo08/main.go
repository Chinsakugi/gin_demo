package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type UserInfo struct {
	gorm.Model
	Name  string `gorm:"column:name"`
	Age   int
	Email string
}

func main() {
	db, err := gorm.Open(mysql.Open("root:123456@tcp(127.0.0.1:3306)/test?charset=utf8mb4&parseTime=True&loc=Local"), &gorm.Config{})

	if err != nil {
		fmt.Println("database connect fail")
	}

	db.AutoMigrate(&UserInfo{})
	//新增
	//user := UserInfo{
	//	Name:  "ccc",
	//	Age:   22,
	//	Email: "1771611869@qq.com",
	//}
	//db.Create(&user)
	//fmt.Println(user.ID)

	//查询
	user := UserInfo{}
	db.First(&user, 2)
	fmt.Println(user)

	users := make([]UserInfo, 10)
	db.Debug().Find(&users)
	fmt.Println(users)

	user1 := UserInfo{}
	db.Where("name = ?", "aaa").First(&user1)
	fmt.Println(user1)

	//修改
	//user.Name = "bbb"
	//db.Model(&user).Update("name", "ccc")

	//删除
	//var u = UserInfo{}
	//u.ID = 3
	//db.Delete(&u)

}
