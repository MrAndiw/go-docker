package database

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func Init() {
	dsn := "host=pg1 user=postgres password=postgres dbname=andi port=5432 sslmode=disable TimeZone=Asia/Jakarta"
	dbs, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	db = dbs
	fmt.Println("Connected To Postgres")
}

func Connect() *gorm.DB {
	return db
}
