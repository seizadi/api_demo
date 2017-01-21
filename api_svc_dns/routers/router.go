// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"github.com/Infoblox-CTO/api-gateway/demo/api_svc_dns/controllers"

	"github.com/astaxie/beego"
)

func init() {
	ns := beego.NewNamespace("/v1",

		beego.NSNamespace("/migrations",
			beego.NSInclude(
				&controllers.MigrationsController{},
			),
		),

		beego.NSNamespace("/host",
			beego.NSInclude(
				&controllers.HostController{},
			),
		),

		beego.NSNamespace("/client_name",
			beego.NSInclude(
				&controllers.ClientNameController{},
			),
		),

		beego.NSNamespace("/zone",
			beego.NSInclude(
				&controllers.ZoneController{},
			),
		),
	)
	beego.AddNamespace(ns)
}
