package main

import (
	"fmt"
	_ "net/http"
	"projects-template/Application"

	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("hellooo")
	app := Application.NewApp()
	app.Gin.GET("/posts", Posts)
	app.Gin.Run()
}

func Posts(c *gin.Context) {
	r := Application.NewRequest(c)
	// to check connection status
	//fmt.Println(request.Connection.Ping())
	r.NotAuth()
	// r.Context.JSON(http.StatusCreated, gin.H{
	// 	"msg": "post has been added",
	// })
}
