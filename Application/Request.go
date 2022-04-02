package Application

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type shareResources interface{
	Share()
}

type Request struct {
	Context    *gin.Context
	DB         *gorm.DB
	Connection *sql.DB
}

func (req *Request) Share(){}

// handle request data
func Req() func(c *gin.Context) Request {
	return func(c *gin.Context) Request {
		var request Request
		request.Context = c
		connectToDataBase(&request)
		return request
	}
}

// init new request
func NewRequest(c *gin.Context) Request {
	request := Req()
	req := request(c)
	return req
}


