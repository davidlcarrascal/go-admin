package gorilla

import (
	// add gorilla adapter
	_ "github.com/davidlcarrascal/go-admin/adapter/gorilla"
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
	"github.com/gorilla/mux"
	"net/http"
	"os"
)

func newGorillaHandler() http.Handler {
	app := mux.NewRouter()
	eng := engine.Default()

	examplePlugin := example.NewExample()
	template.AddComp(chartjs.NewChart())

	if err := eng.AddConfigFromJSON(os.Args[len(os.Args)-1]).
		AddPlugins(admin.NewAdmin(tables.Generators).
			AddGenerator("user", tables.GetUserTable), examplePlugin).
		Use(app); err != nil {
		panic(err)
	}

	eng.HTML("GET", "/admin", tables.GetContent)

	return app
}
