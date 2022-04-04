package Routes

import (
	"projects-template/Controllers/Visitors"
)

func (app RouterApp) VisitorRoutes() {
	app.Gin.GET("/create-user", Visitors.CreateUser)
}
