package Service

import (
	"fmt"
	"log"
	"main/Model"
	"os"
	"strconv"
)

func AddMsg(message *Model.Message) {
	time := message.Time.Format("0601021504")
	lat := strconv.Itoa(int(message.Lat * 10000))
	lng := strconv.Itoa(int(message.Lng * 10000))
	et := message.EventType
	id := strconv.Itoa(int(message.ID))
	code := time + lat + lng + et + id
	message.Code = code
	// 将msg插入数据库
	log.Printf("code: %v", code)
	tmp := Model.InsertMsg(*message)
	log.Printf("msg.code: %v", message.Code)
	path := "data/" + message.EventType + "/" + message.Code + "/"

	message.File = path
	// 创建文件夹存储文件
	errdir := os.MkdirAll(path, os.ModePerm)
	if errdir != nil {
		log.Printf("mkdir %v", errdir)
		return // 如果无法创建目录，则不继续执行
	}
	fmt.Println("目录已创建")

	// 更新数据库中的文件路径
	Model.UploadFile(tmp, path)
	// 将msg插入数据库
	//tmp := Model.InsertMsg(*message)
	//i := strconv.Itoa(int(tmp))
	//path := "data/" + message.EventType + "/" + i + "/"
	//
	//// 设置文件路径
	//message.File = path
	//
	//// 创建文件夹存储文件
	//errdir := os.MkdirAll(path, os.ModePerm)
	//if errdir != nil {
	//	log.Printf("mkdir %v", errdir)
	//	return // 如果无法创建目录，则不继续执行
	//}
	//fmt.Println("目录已创建")
	//
	//// 更新数据库中的文件路径
	//Model.UploadFile(tmp, path)
}

func GetMsg(id int) Model.Message {
	msg := Model.SearchMessage(id)
	return msg
}

func GetAllMsg() {

}
