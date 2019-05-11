package models

import (
	"log"

	"github.com/jinzhu/gorm"
)

var db database

type database interface {
	Create(interface{}) *gorm.DB
	Find(out interface{}, where ...interface{}) *gorm.DB
	Where(query interface{}, args ...interface{}) *gorm.DB
	Save(value interface{}) *gorm.DB
	Delete(value interface{}, where ...interface{}) *gorm.DB
	First(out interface{}, where ...interface{}) *gorm.DB
	AutoMigrate(value interface{}) *gorm.DB
}

func dbConnect() *gorm.DB {
	host := "localhost"
	name := "yummy-api"
	ssl := "disable"

	db, err := gorm.Open("postgres", "host="+host+"name="+name+"ssl="+ssl)
	if err != nil {
		log.Fatal(err)
	}
	return db
}

func adapter() {
	db := dbConnect()
	db.AutoMigrate(User{}, Category{}, Recipe{})
}
