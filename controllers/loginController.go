package controllers

import (
	"github.com/astaxie/beego"
)

type LoginController struct {
	BaseController
}

func (this *LoginController) Get() {

	this.TplName = "login.tpl"
}

func (this *LoginController) Post() {

	uaccount := this.Input().Get("inputAccount")
	pwd := this.Input().Get("inputPassword")

	if beego.AppConfig.String("uaccount") == uaccount &&
		beego.AppConfig.String("pwd") == pwd {
		//		c.TplName = "studentCenter.tpl"
		this.Redirect("/home", 301)
	}

	this.Redirect("/login", 301)
	return

}
