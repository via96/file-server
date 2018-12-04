package routers

import (
	"file-server/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
    beego.Router("/login", &controllers.MainController{}, "post:Login")
    beego.Router("/logged", &controllers.UserController{})
}
