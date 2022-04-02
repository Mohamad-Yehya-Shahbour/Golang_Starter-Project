package main

import (
	"fmt"
	_ "net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("hellooo")
	app := app()
	application := app()
	application.Gin.GET("/posts", Posts)
	application.Gin.Run()
}

func Posts(c *gin.Context) {
	r := newRequest(c)
	// to check connection status
	//fmt.Println(request.Connection.Ping())
	r.NotAuth()
	// r.Context.JSON(http.StatusCreated, gin.H{
	// 	"msg": "post has been added",
	// })
}
