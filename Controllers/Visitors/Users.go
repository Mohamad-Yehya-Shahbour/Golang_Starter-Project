package Visitors

import (
	"projects-template/Application"
	"projects-template/Models"

	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {
	r := Application.NewRequest(c).Auth()
	if !r.IsAdmin {
		r.NotAuth()
		return
	}
	user := Models.User{
		UserName: "mohamad yehya",
		Email:    "yehya@gmail.com",
		Password: "123456789",
	}
	r.DB.Create(&user)
	r.Created(user)
}
