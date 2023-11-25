package Model

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
	"time"
)

type Message struct {
	gorm.Model
	//Title       string
	Time        time.Time
	Lng         float64
	Lat         float64
	Description string
	File        string
	EventType   string
	Source      string
	Password    string
	IsStar      bool
	Code        string
	//EventId     string
}

func CreateMessageDB() {
	db, err := gorm.Open(sqlite.Open("data.db"), &gorm.Config{})
	if err != nil {
		log.Println(err)
		return
	}
	err = db.AutoMigrate(&Message{})
	if err != nil {
		log.Println(err)
		return
	}
}

func SearchMessage(id int) Message {
	db, err := gorm.Open(sqlite.Open("data.db"), &gorm.Config{})
	msg := Message{}
	if err != nil {
		log.Println(err)
		return msg
	}
	db.Where("id = ?", id).Find(&msg)
	return msg
}

func SearchMessageByEvent(etype string) []Message {
	db, err := gorm.Open(sqlite.Open("data.db"), &gorm.Config{})
	if err != nil {
		log.Println(err)
		return nil
	}
	m := []Message{}
	db.Where("EventId = ?", etype).Find(m)
	return m
}

func InsertMsg(msg Message) uint {
	db, err := gorm.Open(sqlite.Open("data.db"), &gorm.Config{})
	if err != nil {
		log.Println(err)
		return 0
	}
	result := db.Create(&msg)
	log.Println(result.Error)
	return msg.ID
}

func GetAllMsg() []Message {
	db, err := gorm.Open(sqlite.Open("data.db"), &gorm.Config{})
	var msg []Message
	if err != nil {
		log.Println(err)
		return msg
	}

	//db.Order("Time").Find(&msg)
	db.Order("created_at desc").Find(&msg)
	return msg
}

func UploadFile(id uint, path string) {
	db, err := gorm.Open(sqlite.Open("data.db"), &gorm.Config{})
	if err != nil {
		log.Println(err)
		return
	}

	// 更新文件路径
	result := db.Model(&Message{}).Where("id = ?", id).Update("file", path)
	if result.Error != nil {
		log.Println(result.Error) // 打印更新错误
		return
	}

	if result.RowsAffected == 0 {
		log.Println("No rows affected") // 如果没有更新任何行
	}
}

//func CreateCode(msg Message, code string) {
//	db, err := gorm.Open(sqlite.Open("data.db"), &gorm.Config{})
//	if err != nil {
//		log.Println(err)
//		return
//	}
//	result := db.Model(&Message{}).Where("id = ?", msg.ID).Update("file", code)
//	if result.Error != nil {
//		log.Println(result.Error) // 打印更新错误
//		return
//	}
//}

func GetMsgNumBySourceThisDay(source string) int {
	now := time.Now()
	dayStart := now.Format("2006-01-02")
	dayStart = dayStart + " 00:00:00"
	dayEnd := now.Format("2006-01-02")
	dayEnd = dayEnd + " 23:59:59"

	log.Println(dayStart, dayEnd)
	db, err := gorm.Open(sqlite.Open("data.db"), &gorm.Config{})
	if err != nil {
		log.Println(err)
		return -1
	}
	var msgs []Message
	db.Where("source = ? AND time >= ? AND time <= ?", source, dayStart, dayEnd).Find(&msgs)
	return len(msgs)
}

func SearchMsgAny(match string) []Message {
	db, err := gorm.Open(sqlite.Open("data.db"), &gorm.Config{})
	if err != nil {
		log.Println(err)
		return nil
	}
	var msgs []Message
	match = "%" + match + "%"
	log.Println(match)
	db.Where("time LIKE ?", match).Or("lat LIKE ?", match).Or("lng LIKE ?", match).Or("description LIKE ?", match).Or("event_type LIKE ?", match).Or("source LIKE ?", match).Find(&msgs)
	return msgs
}

func UpdateMsgStar(msg Message) {
	db, err := gorm.Open(sqlite.Open("data.db"), &gorm.Config{})
	if err != nil {
		log.Println(err)
		return
	}
	msg.IsStar = !msg.IsStar
	log.Println(msg)
	log.Println(msg.IsStar)
	db.Save(&msg)
}

func GetAllStarMsg() []Message {
	db, err := gorm.Open(sqlite.Open("data.db"), &gorm.Config{})
	if err != nil {
		log.Println(err)
		return nil
	}
	var msgs []Message
	db.Where("is_star = ?", true).Find(&msgs)
	return msgs
}
