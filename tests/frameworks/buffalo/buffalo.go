package buffalo

import (
	// add buffalo adapter
	_ "github.com/davidlcarrascal/go-admin/adapter/buffalo"
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

	"github.com/davidlcarrascal/go-admin/template"
	"github.com/davidlcarrascal/go-admin/template/chartjs"

	"github.com/davidlcarrascal/go-admin/engine"
	"github.com/davidlcarrascal/go-admin/plugins/admin"
	"github.com/davidlcarrascal/go-admin/plugins/example"
	"github.com/davidlcarrascal/go-admin/tests/tables"
	"github.com/gobuffalo/buffalo"
	"net/http"
	"os"
)

func newBuffaloHandler() http.Handler {
	bu := buffalo.New(buffalo.Options{
		Env:  "test",
		Addr: "127.0.0.1:9033",
	})

	eng := engine.Default()

	adminPlugin := admin.NewAdmin(tables.Generators)
	adminPlugin.AddGenerator("user", tables.GetUserTable)

	examplePlugin := example.NewExample()

	template.AddComp(chartjs.NewChart())

	if err := eng.AddConfigFromJSON(os.Args[len(os.Args)-1]).
		AddPlugins(adminPlugin, examplePlugin).Use(bu); err != nil {
		panic(err)
	}

	eng.HTML("GET", "/admin", tables.GetContent)

	bu.ServeFiles("/uploads", http.Dir("./uploads"))

	return bu
}
