package main

import (
	"github.com/gin-gonic/gin"
	"main/Model"
	"net/http"
)

func main() {
	Model.CreateMessageDB()
	r := gin.Default()
	r.StaticFS("/data", http.Dir("./data"))

	initRouter(r)

	r.Run()
	//Model.CreateMessageDB()
	//Model.CreateEventDB()
	//Model.CreateLocationDB()

}
