package example

import (
	"github.com/davidlcarrascal/go-admin/context"
	"github.com/davidlcarrascal/go-admin/modules/auth"
	"github.com/davidlcarrascal/go-admin/modules/db"
	"github.com/davidlcarrascal/go-admin/modules/service"
)

func InitRouter(prefix string, srv service.List) *context.App {

	app := context.NewApp()
	route := app.Group(prefix)
	route.GET("/example", auth.Middleware(db.GetConnection(srv)), TestHandler)

	return app
}
