package main

import "github.com/gin-gonic/gin"

func (req Request) Ok(body interface{}) {
	req.closeConnection()
	req.Context.JSON(200, body)
}

func (req Request) Created(body interface{}) {
	req.closeConnection()
	req.Context.JSON(201, body)
}

func (req Request) NotAuth() {
	req.closeConnection()
	req.Context.JSON(401, gin.H{
		"message": "You are not authenticated",
	})
}
