package routers

import (
	"file-server/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.LoginController{})

	beego.Router("/register", &controllers.LoginController{}, "get,post:Register")
	beego.Router("/login", &controllers.LoginController{}, "get,post:Login")
	beego.Router("/logout", &controllers.UserController{}, "get,post:Logout")
	beego.Router("/user/:user:string", &controllers.UserController{}, "get:List")
	beego.Router("/upload", &controllers.UserController{}, "get,post:Upload")
	beego.Router("/download/:linkId([A-Za-z0-9\\:\\+_-]+)", &controllers.UserController{}, "get,post:Download")
	beego.Router("/remove/:linkId([A-Za-z0-9\\:\\+_-]+)", &controllers.UserController{}, "get,post:Remove")
}
