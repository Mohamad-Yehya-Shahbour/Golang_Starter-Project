package Application

import (
	"database/sql"
	"fmt"
	"github.com/bykovme/gotrans"
	"github.com/gin-gonic/gin"
	_ "gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Application struct {
	Gin 		*gin.Engine
	DB         	*gorm.DB
	Connection 	*sql.DB
}

func (app *Application) Share(){}

func App() func() *Application {
	return func() *Application {
		var application Application
		application.Gin = gin.Default()
		connectToDataBase(&application)
		err := gotrans.InitLocales("./Public/Lang")
		if err != nil {
			fmt.Println("Error while loading translation file", err.Error())
		}
		return &application
	}
}

func NewApp() *Application {
	app := App()
	return app()
}
