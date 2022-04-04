package Models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	UserName 	string `json:"username" gorm:"type: varchar(50)"`
	Email 		string `json:"email" gorm:"type: varchar(50)"`
	Password 	string `json:"password" gorm:"type: varchar(50)"`
	Token 	string `json:"token" gorm:"type: varchar(200)"`
	Group 	string `json:"group" gorm:"type: varchar(50)"`
}