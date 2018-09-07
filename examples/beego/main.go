package main

import (
	"github.com/astaxie/beego"
	"goAdmin/adapter"
	"goAdmin/plugins/admin"
	"goAdmin/examples/datamodel"
	"goAdmin/modules/config"
	"goAdmin/engine"
)

func main() {
	app := beego.NewApp()

	eng := engine.Default()

	cfg := config.Config{
		DATABASE: []config.Database{
			{
				IP:           "127.0.0.1",
				PORT:         "3306",
				USER:         "root",
				PWD:          "root",
				NAME:         "godmin",
				MAX_IDLE_CON: 50,
				MAX_OPEN_CON: 150,
				DRIVER:       "mysql",
			},
		},
		AUTH_DOMAIN:  "localhost",
		LANGUAGE:     "cn",
		ADMIN_PREFIX: "admin",
	}

	if err := eng.AddConfig(cfg).AddPlugins(admin.NewAdmin(datamodel.TableFuncConfig)).
		AddAdapter(new(adapter.Beego)).
		Use(app); err != nil {
		panic(err)
	}

	beego.BConfig.Listen.HTTPAddr = "127.0.0.1"
	beego.BConfig.Listen.HTTPPort = 9087
	app.Run()
}
