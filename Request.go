package main

import (
	"database/sql"
	"fmt"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Request struct {
	Context    *gin.Context
	DB         *gorm.DB
	Connection *sql.DB
}

// handle request data
func req() func(c *gin.Context) Request {
	return func(c *gin.Context) Request {
		var request Request
		request.Context = c
		request.DB, request.Connection = connectToDataBase()
		return request
	}
}

// init new request
func newRequest(c *gin.Context) Request {
	request := req()
	req := request(c)
	return req
}

// connect to db
func connectToDataBase() (*gorm.DB, *sql.DB) {

	dsn := "root:@tcp(127.0.0.1:3306)/golangdb?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("Error while connecting the database", err.Error())
	}

	connection, err := db.DB()
	if err != nil {
		fmt.Println("error while getting the connection", err.Error())
	}

	return db, connection
}

func (req Request) closeConnection() {
	req.Connection.Close()
}
