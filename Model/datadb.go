package Model

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
)

type Data struct {
	ID     uint `gorm:"primaryKey;autoIncrement"`
	Code   string
	Locate Location `gorm:"foreignKey:LocationID"`
}

func CreateDataDB() {
	db, err := gorm.Open(sqlite.Open("data.db"), &gorm.Config{})
	if err != nil {
		log.Println(err)
	}
	err = db.AutoMigrate(&Data{})
	if err != nil {
		log.Println(err)
	}
}

func InsertData(data Data) {
	db, err := gorm.Open(sqlite.Open("data.db"), &gorm.Config{})
	if err != nil {
		log.Println(err)
	}
	db.Create(data)
}

func SearchByCode(code string) Data {
	db, err := gorm.Open(sqlite.Open("data.db"), &gorm.Config{})
	if err != nil {
		log.Println(err)
	}
	data := Data{}
	db.Where("Code = ?", code).Find(&data)
	return data
}

func SearchByLocation(location string) []Data {
	db, err := gorm.Open(sqlite.Open("data.db"), &gorm.Config{})
	if err != nil {
		log.Println(err)
	}
	data := []Data{}
	db.Where("Location = ?", location).Find(&data)
	return data
}

func DeleteDataById(id uint) {
	db, err := gorm.Open(sqlite.Open("data.db"), &gorm.Config{})
	if err != nil {
		log.Println(err)
	}
	db.Delete("Id = ?", id)
}
