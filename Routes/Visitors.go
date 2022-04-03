package Routes

import (
	"projects-template/Application"
	"projects-template/Models"


	"github.com/gin-gonic/gin"
)

func (app RouterApp) VisitorRoutes() {
	app.Gin.GET("/createuser", func(c *gin.Context){
		r := Application.NewRequest(c)
		user := Models.User{
			UserName: "mohamad yehya",
			Email: "yehya@gmail.com",
			Password: "123456",
		}
		r.DB.Create(&user)
		r.Created(user)
	})
}
