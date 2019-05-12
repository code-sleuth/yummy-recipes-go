package models

import (
	"log"

	"github.com/jinzhu/gorm"
	// postgres
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

// 	Create(interface{}) *gorm.DB
// 	Find(out interface{}, where ...interface{}) *gorm.DB
// 	Where(query interface{}, args ...interface{}) *gorm.DB
// 	Save(value interface{}) *gorm.DB
// 	Delete(value interface{}, where ...interface{}) *gorm.DB
// 	First(out interface{}, where ...interface{}) *gorm.DB
// 	AutoMigrate(value interface{}) *gorm.DB

func dbConnect() *gorm.DB {

	db, err := gorm.Open("postgres", "host=localhost port=5432 user=code dbname=yummy-api password=inetutils sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	return db
}

// Adapter func
func Adapter() {
	db := dbConnect()
	db.AutoMigrate(&User{}, &Category{}, &Recipe{})
	defer db.Close()
}
