package Visitors

import (
	"projects-template/Application"
	_"projects-template/Models"
	"github.com/bykovme/gotrans"
	"github.com/gin-gonic/gin"
)

func Register(c *gin.Context) {
	r := Application.NewRequestWithAuth(c)
	// binding Request
	var user Models.User
	r.Context.ShouldBindJSON(&user)

	// validate request
	err := validation.Errors{
		"username": validation.Validate(user.UserName, validation.Required.Error(gotrans.T("required")), validation.Length(5, 50).Error(gotrans.T("min-max"))),
		"email": validation.Validate(user.Email, validation.Required.Error(gotrans.T("required")), is.Email.Error(gotrans.T("email")), validation.Length(5, 50).Error(gotrans.T("min-max")) ),
		"password": validation.Validate(user.Password, validation.Required.Error(gotrans.T("required")), validation.Length(5, 50)),
	}.Filter()

	if err != nil {
		r.BadRequest(err)
		return
	}
	r.Ok(user)

	

	// user := Models.User{
	// 	UserName: "mohamad yehya",
	// 	Email:    "yehya@gmail.com",
	// 	Password: "123456789",
	// }
	// r.DB.Create(&user)
	
}
