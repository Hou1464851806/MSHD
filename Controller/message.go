package Controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"log"
	"main/Model"
	"main/Service"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

func UploadMsg(c *gin.Context) {
	var msgReqBody Model.Message
	log.Println(c.PostForm("datetime"))
	tmpTime := strings.Replace(c.PostForm("datetime"), "T", " ", 1)
	tmpTime = tmpTime + ":00"
	log.Println(tmpTime)
	msgReqBody.Lat, _ = strconv.ParseFloat(c.PostForm("latitude"), 64)
	msgReqBody.Lng, _ = strconv.ParseFloat(c.PostForm("longitude"), 64)
	msgReqBody.EventType = c.PostForm("disaster-type")
	msgReqBody.Source = c.PostForm("id-number")
	msgReqBody.Password = c.PostForm("password")
	msgReqBody.Description = c.PostForm("description")
	msgReqBody.Time, _ = time.Parse(time.DateTime, tmpTime)

	msg := Model.Message{
		Model:       gorm.Model{},
		Time:        msgReqBody.Time,
		Lng:         msgReqBody.Lng,
		Lat:         msgReqBody.Lat,
		Description: msgReqBody.Description,
		File:        "",
		EventType:   msgReqBody.EventType,
		Source:      msgReqBody.Source,
		IsStar:      false,
		Password:    msgReqBody.Password,
	}
	log.Printf("上传的Msg：%v", msg)
	Service.AddMsg(&msg)
	//保存文件
	form, _ := c.MultipartForm()
	log.Println(form)
	files, ok := form.File["file"]
	log.Println(files)
	if ok {
		for _, f := range files {
			filePath := msg.File + f.Filename
			err := c.SaveUploadedFile(f, filePath)
			if err != nil {
				log.Println(err)
				return
			}
			//c.JSON(200, gin.H{
			//	"msg":  "success",
			//	"name": f.Filename,
			//	"size": f.Size,
			//})
		}
	}
	num := Model.GetMsgNumBySourceThisDay(msgReqBody.Source)
	c.JSON(http.StatusOK, gin.H{"msg": "success", "num": num})
}

func GetMsg(c *gin.Context) {
	id, _ := strconv.Atoi(c.Query("id"))
	m := Service.GetMsg(id)
	url := "http://" + c.Request.Host + "/"
	fileNames := listFile(m.File)
	var urls []string
	for _, fileName := range fileNames {
		urls = append(urls, url+m.File+fileName)
	}
	response := getAllMsgResponse{
		ID:           m.ID,
		Datetime:     m.Time,
		Latitude:     m.Lat,
		Longitude:    m.Lng,
		Description:  m.Description,
		File:         urls,
		DisasterType: m.EventType,
		IsStar:       m.IsStar,
	}

	c.JSON(http.StatusOK, response)
}

type getAllMsgResponse struct {
	ID           uint      `json:"id"`
	Datetime     time.Time `json:"datetime,string"`
	Latitude     float64   `json:"latitude"`
	Longitude    float64   `json:"longitude"`
	Description  string    `json:"description"`
	File         []string  `json:"file"`
	DisasterType string    `json:"disaster-type"`
	IsStar       bool      `json:"is-star"`
}

func listFile(path string) []string {
	var fileNames []string
	dir, err := os.Open(path)
	if err != nil {
		fmt.Println("Error:", err)
		return nil
	}
	defer dir.Close()
	files, err := dir.Readdir(-1)
	if err != nil {
		fmt.Println("Error:", err)
		return nil
	}
	for _, file := range files {
		fileNames = append(fileNames, file.Name())
	}
	return fileNames
}

func GetAllMsg(c *gin.Context) {
	url := "http://" + c.Request.Host + "/"
	response := []getAllMsgResponse{}
	msg := Model.GetAllMsg()
	for _, m := range msg {
		fileNames := listFile(m.File)
		var urls []string
		for _, fileName := range fileNames {
			urls = append(urls, url+m.File+fileName)
		}
		tmp := getAllMsgResponse{
			ID:           m.ID,
			Datetime:     m.Time,
			Latitude:     m.Lat,
			Longitude:    m.Lng,
			Description:  m.Description,
			File:         urls,
			DisasterType: m.EventType,
			IsStar:       m.IsStar,
		}
		response = append(response, tmp)
	}
	log.Printf("%v", response)
	c.JSON(http.StatusOK, response)
	log.Println(c)
}

func GetMsgFile(c *gin.Context) {
	id, _ := strconv.Atoi(c.Query("Id"))
	msg := Model.SearchMessage(id)
	c.File(msg.File)
}

type SearchMsgReq struct {
	Match string `json:"match"`
}

func SearchMsg(c *gin.Context) {
	//json := SearchMsgReq{}
	//c.BindJSON(&json)
	//match := json.Match
	match := c.Query("match")
	url := "http://" + c.Request.Host + "/"
	response := []getAllMsgResponse{}
	msg := Model.SearchMsgAny(match)
	for _, m := range msg {
		fileNames := listFile(m.File)
		var urls []string
		for _, fileName := range fileNames {
			urls = append(urls, url+m.File+fileName)
		}
		tmp := getAllMsgResponse{
			ID:           m.ID,
			Datetime:     m.Time,
			Latitude:     m.Lat,
			Longitude:    m.Lng,
			Description:  m.Description,
			File:         urls,
			DisasterType: m.EventType,
			IsStar:       m.IsStar,
		}
		response = append(response, tmp)
	}
	log.Printf("%v", response)
	c.JSON(http.StatusOK, response)
}

type starMsgReq struct {
	ID int `json:"id"`
}

func StarMsg(c *gin.Context) {
	json := starMsgReq{}
	c.Bind(&json)
	log.Println(json.ID)
	msg := Model.SearchMessage(json.ID)
	Model.UpdateMsgStar(msg)
	c.JSON(http.StatusOK, gin.H{"msg": "success"})
}

func GetAllStar(c *gin.Context) {
	url := "http://" + c.Request.Host + "/"
	response := []getAllMsgResponse{}
	msg := Model.GetAllStarMsg()
	for _, m := range msg {
		fileNames := listFile(m.File)
		var urls []string
		for _, fileName := range fileNames {
			urls = append(urls, url+m.File+fileName)
		}
		tmp := getAllMsgResponse{
			ID:           m.ID,
			Datetime:     m.Time,
			Latitude:     m.Lat,
			Longitude:    m.Lng,
			Description:  m.Description,
			File:         urls,
			DisasterType: m.EventType,
			IsStar:       m.IsStar,
		}
		response = append(response, tmp)
	}
	log.Printf("%v", response)
	c.JSON(http.StatusOK, response)
	log.Println(c)
}
