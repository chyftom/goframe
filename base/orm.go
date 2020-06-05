package db

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)


type Product struct {
	gorm.Model
	Code string
	Price uint
	CreatedBy string
	UpdatedBy string
}

type book struct {

}

func GetProduct() {

	psqlInfo = fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := gorm.Open("postgres", psqlInfo)
	if err != nil {
		panic("连接数据库失败")
	}
	defer db.Close()
	db.SingularTable(true)

	db.AutoMigrate(&Product{},&book{})

	//// 自动迁移模式
	db.LogMode(true)
	// 创建
	db.Create(&Product{Code: "L1212", Price: 1000})

	// 读取
	var product Product
	db.First(&product, 1) // 查询id为1的product
	db.First(&product, "code = ?", "L1212") // 查询code为l1212的product

	// 更新 - 更新product的price为2000
	db.Model(&product).Update("Price", 2000)

	// 删除 - 删除product
	db.Delete(&product)
}