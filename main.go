package main

import (
	_"fmt"
	_ "net/http"
	"projects-template/Application"
	"projects-template/Models"
	"projects-template/Routes"
	_ "github.com/gin-gonic/gin"
)

func main() {
	
	app := Application.NewApp()

	// migrate models
	app.DB.AutoMigrate(&Models.User{})

	// add routes
	routerApp := Routes.RouterApp{app}
	routerApp.Routing()

	// close conection
	Application.CloseConnection(app)
	// start app server
	app.Gin.Run()
}
