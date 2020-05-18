package gin

import (
	// add gin adapter
	_ "github.com/davidlcarrascal/go-admin/adapter/gin"
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
	"github.com/davidlcarrascal/go-admin/tests/tables"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
	"os"
)

func newHandler() http.Handler {
	r := gin.Default()

	gin.SetMode(gin.ReleaseMode)
	gin.DefaultWriter = ioutil.Discard

	eng := engine.Default()

	adminPlugin := admin.NewAdmin(tables.Generators)
	adminPlugin.AddGenerator("user", tables.GetUserTable)

	template.AddComp(chartjs.NewChart())

	if err := eng.AddConfigFromJSON(os.Args[len(os.Args)-1]).
		AddPlugins(adminPlugin).
		Use(r); err != nil {
		panic(err)
	}

	eng.HTML("GET", "/admin", tables.GetContent)

	r.Static("/uploads", "./uploads")

	return r
}
