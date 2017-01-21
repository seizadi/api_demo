package main

import (
	_ "github.com/Infoblox-CTO/api-gateway/demo/api_svc_threats/routers"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/lib/pq"
)

func init() {
	orm.RegisterDataBase("default", "postgres", "postgres://seizadi:seizadi@127.0.0.1:5432/demo?sslmode=disable")
}

func main() {
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}
	beego.Run()
}

