package echo

import (
	// add echo adapter
	_ "github.com/davidlcarrascal/go-admin/adapter/echo"
	// add mysql driver
	_ "github.com/davidlcarrascal/go-admin/modules/db/drivers/mysql"
	// add postgresql driver
	_ "github.com/davidlcarrascal/go-admin/modules/db/drivers/postgres"
	// add sqlite driver
	_ "github.com/davidlcarrascal/go-admin/modules/db/drivers/sqlite"
	// add mssql driver
	_ "github.com/davidlcarrascal/go-admin/modules/db/drivers/mssql"
	// add adminlte ui theme
	_ "github.com/GoAdminGroup/themes/adminlte"

	"github.com/davidlcarrascal/go-admin/engine"
	"github.com/davidlcarrascal/go-admin/plugins/admin"
	"github.com/davidlcarrascal/go-admin/plugins/example"
	"github.com/davidlcarrascal/go-admin/template"
	"github.com/davidlcarrascal/go-admin/template/chartjs"
	"github.com/davidlcarrascal/go-admin/tests/tables"
	"github.com/labstack/echo/v4"
	"net/http"
	"os"
)

func newEchoHandler() http.Handler {
	e := echo.New()

	eng := engine.Default()

	adminPlugin := admin.NewAdmin(tables.Generators)
	adminPlugin.AddGenerator("user", tables.GetUserTable)
	template.AddComp(chartjs.NewChart())

	examplePlugin := example.NewExample()

	if err := eng.AddConfigFromJSON(os.Args[len(os.Args)-1]).
		AddPlugins(adminPlugin, examplePlugin).Use(e); err != nil {
		panic(err)
	}

	eng.HTML("GET", "/admin", tables.GetContent)

	return e
}
