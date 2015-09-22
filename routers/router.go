package routers

import (
	"dockerwebconsole/config"
	"dockerwebconsole/controllers"
	"github.com/astaxie/beego"
)

func init() {

	var defaultConfig = config.GetConfig()

	var mainController = &controllers.MainController{ConfigData: defaultConfig}

	beego.Router("/", mainController, "get:Get")
	beego.Router("/home", mainController)
	beego.Router("/notice", mainController, "get:Notice")
	beego.Router("/about", mainController, "get:About")
	beego.Router("/about", mainController, "post:ContactHandler")

	beego.Router("/hosts", mainController, "get:HomeHandler")
	beego.Router("/hosts/console/:action", mainController, "get:ConsoleHandler")

	beego.Router("/user/login/:back", mainController, "get,post:Login")
	beego.Router("/user/logout", mainController, "get:Logout")
	beego.Router("/user/register", mainController, "get,post:Register")
	beego.Router("/user/profile", mainController, "get,post:Profile")
	beego.Router("/user/verify/:uuid({[0-9A-F]{8}-[0-9A-F]{4}-4[0-9A-F]{3}-[89AB][0-9A-F]{3}-[0-9A-F]{12}})", mainController, "get:Verify")
	beego.Router("/user/remove", mainController, "get,post:Remove")
	beego.Router("/user/forgot", mainController, "get,post:Forgot")
	beego.Router("/user/reset/:uuid({[0-9A-F]{8}-[0-9A-F]{4}-4[0-9A-F]{3}-[89AB][0-9A-F]{3}-[0-9A-F]{12}})", mainController, "get,post:Reset")

	beego.Router("/appadmin/index", &controllers.AdminController{}, "get,post:Index")
	beego.Router("/appadmin/index/:parms", &controllers.AdminController{}, "get,post:Index")
	beego.Router("/appadmin/add/:parms", &controllers.AdminController{}, "get,post:Add")
	beego.Router("/appadmin/update/:username", &controllers.AdminController{}, "get,post:Update")
}
