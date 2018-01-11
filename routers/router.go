package routers

import (
	"CoursePj/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/login", &controllers.LoginController{})
	beego.Router("/home", &controllers.HomeController{})
	beego.Router("/home/edit", &controllers.HomeController{}, "*:TheoryCourseEdit")
	beego.Router("/home/add", &controllers.HomeController{}, "*:TheoryCourseAdd")
	beego.AutoRouter(&controllers.HomeController{})
}
