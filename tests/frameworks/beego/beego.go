package beego

import (
	// add beego adapter
	_ "github.com/davidlcarrascal/go-admin/adapter/beego"
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
	"github.com/astaxie/beego"
	"net/http"
	"os"
)

func newBeegoHandler() http.Handler {

	app := beego.NewApp()

	eng := engine.Default()
	adminPlugin := admin.NewAdmin(tables.Generators)
	adminPlugin.AddGenerator("user", tables.GetUserTable)

	examplePlugin := example.NewExample()

	if err := eng.AddConfigFromJSON(os.Args[len(os.Args)-1]).
		AddPlugins(adminPlugin, examplePlugin).Use(app); err != nil {
		panic(err)
	}

	template.AddComp(chartjs.NewChart())

	eng.HTML("GET", "/admin", tables.GetContent)

	beego.BConfig.Listen.HTTPAddr = "127.0.0.1"
	beego.BConfig.Listen.HTTPPort = 9087

	return app.Handlers
}
