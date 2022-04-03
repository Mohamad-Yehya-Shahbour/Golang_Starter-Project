package Application

import (
	"database/sql"
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)


func makeConnection()*gorm.DB{

	dsn := "root:@tcp(127.0.0.1:3306)/golangdb?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		fmt.Println("Error while connecting the database", err.Error())
	}

	return db
}

func returnConnection (db *gorm.DB) *sql.DB{
	connection, err := db.DB()
	if err != nil {
		fmt.Println("error while getting the connection", err.Error())
	}
	return connection
}

// connect to db
func connectToDataBase(share shareResources) {
	switch share.(type){
	case *Application:
		app := share.(*Application)
		app.DB = makeConnection()
		app.Connection = returnConnection(app.DB)
	case *Request:
		req := share.(*Request)
		req.DB = makeConnection()
		req.Connection = returnConnection(req.DB)
	}
}

func CloseConnection(share shareResources) {
	switch share.(type){
	case *Application:
		app := share.(*Application)
		app.Connection.Close()
	case *Request:
		req := share.(*Request)
		req.Connection.Close()
	}
}