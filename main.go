package main

import (
	"fmt"
	_ "net/http"
	"projects-template/Application"
	"projects-template/Models"
	"projects-template/Routes"

	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("hellooo")
	app := Application.NewApp()
	// migrate models
	app.DB.AutoMigrate(&Models.User{})
	routerApp := Routes.RouterApp(&app)
	routerApp.VisitorRoutes()
	// close conection
	Application.CloseConnection(&app)

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
