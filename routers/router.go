package routers

import (
	"blog/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
    beego.Router("/auth/register", &controllers.AuthController{},"get,post:Register")
    beego.Router("/auth/login", &controllers.AuthController{},"get,post:Login")
    beego.Router("/auth/logout", &controllers.AuthController{},"get:Logout")
    beego.Router("/user/?:id", &controllers.UserController{})
}
