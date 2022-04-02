package Application

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	_"gorm.io/driver/mysql"
	"database/sql"
)

type Application struct {
	Gin *gin.Engine
	DB         *gorm.DB
	Connection *sql.DB
}

func (app *Application) Share(){}

func App() func() Application {
	return func() Application {
		var application Application
		application.Gin = gin.Default()
		connectToDataBase(&application)
		return application
	}
}

func NewApp() Application {
	app := App()
	application := app()
	return application
}