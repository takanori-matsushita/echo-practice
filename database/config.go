package database

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Connect() *gorm.DB {
	db, dbErr := gorm.Open(mysql.Open("root@tcp(localhost:3306)/go_practice"))
	if dbErr != nil {
		panic(dbErr)
	}
	// db.AutoMigrate(&Models.User{})
	return db
}
