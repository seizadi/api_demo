package routers

import (
	"github.com/astaxie/beego"
)

func init() {

	beego.GlobalControllerRouter["github.com/seizadi/api_demo/api_svc_cfw/controllers:ClientNameController"] = append(beego.GlobalControllerRouter["github.com/seizadi/api_demo/api_svc_cfw/controllers:ClientNameController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/seizadi/api_demo/api_svc_cfw/controllers:ClientNameController"] = append(beego.GlobalControllerRouter["github.com/seizadi/api_demo/api_svc_cfw/controllers:ClientNameController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/seizadi/api_demo/api_svc_cfw/controllers:ClientNameController"] = append(beego.GlobalControllerRouter["github.com/seizadi/api_demo/api_svc_cfw/controllers:ClientNameController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/seizadi/api_demo/api_svc_cfw/controllers:ClientNameController"] = append(beego.GlobalControllerRouter["github.com/seizadi/api_demo/api_svc_cfw/controllers:ClientNameController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/seizadi/api_demo/api_svc_cfw/controllers:ClientNameController"] = append(beego.GlobalControllerRouter["github.com/seizadi/api_demo/api_svc_cfw/controllers:ClientNameController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/seizadi/api_demo/api_svc_cfw/controllers:HostController"] = append(beego.GlobalControllerRouter["github.com/seizadi/api_demo/api_svc_cfw/controllers:HostController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/seizadi/api_demo/api_svc_cfw/controllers:HostController"] = append(beego.GlobalControllerRouter["github.com/seizadi/api_demo/api_svc_cfw/controllers:HostController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/seizadi/api_demo/api_svc_cfw/controllers:HostController"] = append(beego.GlobalControllerRouter["github.com/seizadi/api_demo/api_svc_cfw/controllers:HostController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/seizadi/api_demo/api_svc_cfw/controllers:HostController"] = append(beego.GlobalControllerRouter["github.com/seizadi/api_demo/api_svc_cfw/controllers:HostController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/seizadi/api_demo/api_svc_cfw/controllers:HostController"] = append(beego.GlobalControllerRouter["github.com/seizadi/api_demo/api_svc_cfw/controllers:HostController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/seizadi/api_demo/api_svc_cfw/controllers:MigrationsController"] = append(beego.GlobalControllerRouter["github.com/seizadi/api_demo/api_svc_cfw/controllers:MigrationsController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/seizadi/api_demo/api_svc_cfw/controllers:MigrationsController"] = append(beego.GlobalControllerRouter["github.com/seizadi/api_demo/api_svc_cfw/controllers:MigrationsController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/seizadi/api_demo/api_svc_cfw/controllers:MigrationsController"] = append(beego.GlobalControllerRouter["github.com/seizadi/api_demo/api_svc_cfw/controllers:MigrationsController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/seizadi/api_demo/api_svc_cfw/controllers:MigrationsController"] = append(beego.GlobalControllerRouter["github.com/seizadi/api_demo/api_svc_cfw/controllers:MigrationsController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/seizadi/api_demo/api_svc_cfw/controllers:MigrationsController"] = append(beego.GlobalControllerRouter["github.com/seizadi/api_demo/api_svc_cfw/controllers:MigrationsController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/seizadi/api_demo/api_svc_cfw/controllers:ZoneController"] = append(beego.GlobalControllerRouter["github.com/seizadi/api_demo/api_svc_cfw/controllers:ZoneController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/seizadi/api_demo/api_svc_cfw/controllers:ZoneController"] = append(beego.GlobalControllerRouter["github.com/seizadi/api_demo/api_svc_cfw/controllers:ZoneController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/seizadi/api_demo/api_svc_cfw/controllers:ZoneController"] = append(beego.GlobalControllerRouter["github.com/seizadi/api_demo/api_svc_cfw/controllers:ZoneController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/seizadi/api_demo/api_svc_cfw/controllers:ZoneController"] = append(beego.GlobalControllerRouter["github.com/seizadi/api_demo/api_svc_cfw/controllers:ZoneController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/seizadi/api_demo/api_svc_cfw/controllers:ZoneController"] = append(beego.GlobalControllerRouter["github.com/seizadi/api_demo/api_svc_cfw/controllers:ZoneController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

}
