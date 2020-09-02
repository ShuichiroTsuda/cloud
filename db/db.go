package db

import (
    "github.com/jinzhu/gorm"
    "github.com/jinzhu/gorm/dialects/postgres" // Use PostgreSQL in gorm
)

ver (
	db *gorm.DB
	err error
)

func Init() {
	db, err = gorm.Open("mysql", "host=0.0.0.0 port=5432 user=gorm dbname=gorm password=gorm sslmode=disable")
	if err != nil {
        panic(err)
    }
}

func GetDB() *gorm.DB {
	return db
}

func Close() {
	if err := db.Close(); err != nil {
		panic(err)
	}
}