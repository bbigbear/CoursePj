package routers

import (
	"CoursePj/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/login", &controllers.LoginController{})

	//首页-理论课程
	beego.Router("/home", &controllers.HomeController{})
	beego.Router("/home/edit", &controllers.HomeController{}, "*:TheoryCourseEdit")
	beego.Router("/home/add", &controllers.HomeController{}, "*:TheoryCourseAdd")
	beego.Router("/home/search", &controllers.HomeController{}, "*:TheoryCourseSearch")
	beego.AutoRouter(&controllers.HomeController{})

	//实践环节
	beego.Router("/practice", &controllers.PracticeController{})
	beego.Router("/practice/edit", &controllers.PracticeController{}, "*:PracticeEdit")
	beego.Router("/practice/add", &controllers.PracticeController{}, "*:PracticeAdd")
	beego.Router("/practice/search", &controllers.PracticeController{}, "*:PracticeSearch")
	beego.AutoRouter(&controllers.PracticeController{})

	//专业信息
	beego.Router("/pm", &controllers.PmController{})
	beego.Router("/pm/edit", &controllers.PmController{}, "*:PmEdit")
	beego.Router("/pm/add", &controllers.PmController{}, "*:PmAdd")
	beego.Router("/pm/search", &controllers.PmController{}, "*:PmSearch")
	beego.AutoRouter(&controllers.PmController{})

	//专业理论
	beego.Router("/ptcourse", &controllers.PTCourseController{})
	beego.Router("/ptcourse/edit", &controllers.PTCourseController{}, "*:PTCourseEdit")
	beego.AutoRouter(&controllers.PTCourseController{})

	//专业实践环节
	beego.Router("/pplink", &controllers.PPLinkController{})
	beego.Router("/pplink/edit", &controllers.PPLinkController{}, "*:PPLinkEdit")
	beego.AutoRouter(&controllers.PPLinkController{})

	//专业学分
	beego.Router("/pgcredit", &controllers.PGCreditController{})
	beego.Router("/pgcredit/add", &controllers.PGCreditController{}, "*:PgcAdd")
	beego.Router("/pgcredit/edit", &controllers.PGCreditController{}, "*:PgcEdit")
	beego.AutoRouter(&controllers.PGCreditController{})

	//复制计划
	beego.Router("/copyplan/profession", &controllers.CopyPlanController{}, "*:GetProfessionPlan")
	beego.Router("/copyplan/year", &controllers.CopyPlanController{}, "*:GetYearPlan")
	beego.Router("/copyplan/profession/search", &controllers.CopyPlanController{}, "*:PPSearch")
	beego.Router("/copyplan/year/search", &controllers.CopyPlanController{}, "*:GYSearch")
	beego.AutoRouter(&controllers.CopyPlanController{})
}
