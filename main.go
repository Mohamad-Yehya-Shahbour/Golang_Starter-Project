package main

import (
	"fmt"
	"net/http"

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
	request := newRequest(c)
	fmt.Println(request.Connection.Ping()) // to check connection status
	request.closeConnection()
	fmt.Println(request.Connection.Ping())
	request.Context.JSON(http.StatusCreated, gin.H{
		"msg": "post has been added",
	})
}
