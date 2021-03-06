package main

import (
	_ "github.com/davidlcarrascal/go-admin/adapter/buffalo"
	_ "github.com/davidlcarrascal/go-admin/modules/db/drivers/mysql"
	"github.com/davidlcarrascal/go-admin/template"
	"github.com/davidlcarrascal/go-admin/template/chartjs"
	"log"
	"os"
	"os/signal"

	"github.com/davidlcarrascal/go-admin/engine"
	"github.com/davidlcarrascal/go-admin/examples/datamodel"
	"github.com/davidlcarrascal/go-admin/modules/config"
	"github.com/davidlcarrascal/go-admin/modules/language"
	"github.com/davidlcarrascal/go-admin/plugins/admin"
	"github.com/davidlcarrascal/go-admin/plugins/example"
	"github.com/GoAdminGroup/themes/adminlte"
	"github.com/gobuffalo/buffalo"
	"net/http"
)

func main() {
	bu := buffalo.New(buffalo.Options{
		Env:  "test",
		Addr: "127.0.0.1:9033",
	})

	eng := engine.Default()

	cfg := config.Config{
		Databases: config.DatabaseList{
			"default": {
				Host:       "127.0.0.1",
				Port:       "3306",
				User:       "root",
				Pwd:        "root",
				Name:       "godmin",
				MaxIdleCon: 50,
				MaxOpenCon: 150,
				Driver:     config.DriverMysql,
			},
		},
		UrlPrefix: "admin",
		Store: config.Store{
			Path:   "./uploads",
			Prefix: "uploads",
		},
		Language:    language.EN,
		IndexUrl:    "/",
		Debug:       true,
		ColorScheme: adminlte.ColorschemeSkinBlack,
	}

	adminPlugin := admin.NewAdmin(datamodel.Generators).AddDisplayFilterXssJsFilter()

	template.AddComp(chartjs.NewChart())

	// add generator, first parameter is the url prefix of table when visit.
	// example:
	//
	// "user" => http://localhost:9033/admin/info/user
	//
	adminPlugin.AddGenerator("user", datamodel.GetUserTable)

	// customize a plugin

	examplePlugin := example.NewExample()

	// load from golang.Plugin
	//
	// examplePlugin := plugins.LoadFromPlugin("../datamodel/example.so")

	// customize the login page
	// example: https://github.com/GoAdminGroup/demo.go-admin.cn/blob/master/main.go#L39
	//
	// template.AddComp("login", datamodel.LoginPage)

	// load config from json file
	//
	// eng.AddConfigFromJSON("../datamodel/config.json")

	if err := eng.AddConfig(cfg).AddPlugins(adminPlugin, examplePlugin).Use(bu); err != nil {
		panic(err)
	}

	bu.ServeFiles("/uploads", http.Dir("./uploads"))

	// you can custom your pages like:

	eng.HTML("GET", "/admin", datamodel.GetContent)

	go func() {
		_ = bu.Serve()
	}()

	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Print("closing database connection")
	eng.MysqlConnection().Close()
}
