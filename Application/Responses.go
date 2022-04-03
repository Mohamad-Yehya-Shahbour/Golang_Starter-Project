package Application

import "github.com/gin-gonic/gin"

func (req Request) Ok(body interface{}) {
	CloseConnection(&req)
	req.Context.JSON(200, body)
}

func (req Request) Created(body interface{}) {
	CloseConnection(&req)
	req.Context.JSON(201, body)
}

func (req Request) NotAuth() {
	CloseConnection(&req)
	req.Context.JSON(401, gin.H{
		"message": "You are not authenticated",
	})
}
