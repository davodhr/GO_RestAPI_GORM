package models

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	// dsn := "postgres://dave:dave_password@postgres:5432/rect_db"
	dsn := "host=postgres user=dave password=dave_password dbname=rect_db port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{DisableForeignKeyConstraintWhenMigrating: true})
	if err != nil {
		panic(err)
	}
	err = db.AutoMigrate(&Rectangle{})
	if err != nil {
		fmt.Println(err)
	}
	DB = db
}
