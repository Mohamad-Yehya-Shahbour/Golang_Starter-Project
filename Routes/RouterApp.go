package Routes

import "projects-template/Application"

type RouterApp struct {
	*Application.Application
}

func (app RouterApp) Routing() {
	app.VisitorRoutes()
	app.authRoutes()
	app.adminRoutes()
}
