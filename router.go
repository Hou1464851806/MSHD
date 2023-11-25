package main

import (
	"github.com/gin-gonic/gin"
	"main/Controller"
)

func initRouter(r *gin.Engine) {
	r.POST("/uploadmsg", Controller.UploadMsg)
	r.POST("/uploadevent", Controller.UploadEvent)
	r.GET("/getmsg", Controller.GetMsg)
	r.GET("/getallmsg", Controller.GetAllMsg)
	r.GET("/getfile", Controller.GetMsgFile)
	r.GET("/searchmsg", Controller.SearchMsg)
	r.GET("/getallstar", Controller.GetAllStar)
	r.POST("/starmsg", Controller.StarMsg)
}
