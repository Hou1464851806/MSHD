package Service

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
	"main/Model"
	"time"
)

func CreateEvent(msg Model.Message, local string) {
	db, err := gorm.Open(sqlite.Open("data.db"), &gorm.Config{})
	if err != nil {
		log.Println(err)
		return
	}
	m := Model.Event{}
	db.Where("EventType = ?", msg.EventType).Find(m)
	if m.EventType == "" {
		Model.CreateEvent(Model.Event{
			Model:     gorm.Model{},
			StartTime: msg.Time,
			EventType: msg.EventType,
			EndTime:   time.Time{},
			Local:     local,
			MsgNum:    0,
		})
	} else {
		m.MsgNum = m.MsgNum + 1
	}
}
