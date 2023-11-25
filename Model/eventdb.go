package Model

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
	"time"
)

type Event struct {
	gorm.Model
	StartTime time.Time
	EventType string
	EndTime   time.Time
	Local     string

	MsgNum int
}

func CreateEventDB() {
	db, err := gorm.Open(sqlite.Open("data.db"), &gorm.Config{})
	if err != nil {
		log.Println(err)
	}
	err = db.AutoMigrate(&Event{})
	if err != nil {
		log.Println(err)
	}
}

// 查询该灾害是否存在
func SearchEvent(etype string) (bool, Event) {
	db, err := gorm.Open(sqlite.Open("data.db"), &gorm.Config{})
	if err != nil {
		log.Println(err)
	}
	e := Event{}
	db.Where("EventType = ?", etype).Find(e)
	if e.EventType == "" {
		return false, e
	} else {
		return true, e
	}
}

func CreateEvent(event Event) {
	db, err := gorm.Open(sqlite.Open("data.db"), &gorm.Config{})
	if err != nil {
		log.Println(err)
	}
	db.Create(&event)
}

func GetMsgNum(Etype string) {
	db, err := gorm.Open(sqlite.Open("data.db"), &gorm.Config{})
	if err != nil {
		log.Println(err.Error())
		return
	}
	msg := []Message{}
	db.Where("EventType=?, ", Etype).Find(msg)
}
