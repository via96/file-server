package routers

import (
	"file-server/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.LoginController{}, "get")
    //beego.Router("/login", &controllers.LoginController{}, "post:Login")
    //beego.Router("/logged", &controllers.UserController{})

	beego.Router("/register", &controllers.LoginController{}, "get,post:Register")
	beego.Router("/login", &controllers.LoginController{}, "get,post:Login")
	beego.Router("/logout", &controllers.LoginController{}, "get:Logout")
}
