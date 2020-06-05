package db

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
	"log"
)
const (
	host     = "192.168.9.128"
	port     = 5432
	user     = "postgres"
	password = "postgrespassword"
	dbname   = "postgres"
)

var db *gorm.DB

func init() {

	var err error
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err = gorm.Open("postgres", psqlInfo)

	if err != nil {
		log.Fatal("Failed to init db:", err)
	}

	db.SingularTable(true)
	//// 自动迁移模式
	db.LogMode(true)
}



func GetDB() *gorm.DB{
	return db
}