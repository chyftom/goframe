package model

import (
	"awesomeProject/model/db"
	"fmt"
	"github.com/jinzhu/gorm"
	"time"
)

type User struct {
	gorm.Model
	UserName string
	Email string
	CreatedBy time.Time
	UpdatedBy string
}

func init() {
	db.GetDB().AutoMigrate(&User{})

	// 创建
	db.GetDB().Create(&User{UserName: "Kaifeng", Email: "yanfeng@qq.com",UpdatedBy: "yanking"})
}

func GetUserByID(id int32)  {
	// 读取
	var user User
	db.GetDB().First(&user, 2) // 查询id为1的product
	fmt.Println(user)
}