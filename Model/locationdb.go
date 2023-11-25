package Model

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
)

type Location struct {
	LocationID uint `gorm:"primaryKey;autoIncrement"`
	Number     string
	Province   string
	City       string
	County     string
	Town       string
	Village    string
}

func CreateLocationDB() {
	db, err := gorm.Open(sqlite.Open("data.db"), &gorm.Config{})
	if err != nil {
		log.Println(err)
	}
	err = db.AutoMigrate(&Location{})
	if err != nil {
		log.Println(err)
	}

}
