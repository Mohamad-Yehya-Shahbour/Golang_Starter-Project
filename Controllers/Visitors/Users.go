package Visitors

import (
	"projects-template/Application"
	_"projects-template/Models"

	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {
	r, auth := Application.AuthRequest(c)
	if !auth {
		return
	}
	r.Ok(r.User)

	// user := Models.User{
	// 	UserName: "mohamad yehya",
	// 	Email:    "yehya@gmail.com",
	// 	Password: "123456789",
	// }
	// r.DB.Create(&user)
	
}
