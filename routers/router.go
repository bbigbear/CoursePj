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
	beego.Router("/home/add/action", &controllers.HomeController{}, "*:TheoryCourseAddAction")
	beego.Router("/home/search", &controllers.HomeController{}, "*:TheoryCourseSearch")
	beego.Router("/home/updata", &controllers.HomeController{}, "*:TheoryCourseUpdata")
	beego.Router("/home/delete", &controllers.HomeController{}, "*:TheoryCourseDelete")
	beego.AutoRouter(&controllers.HomeController{})

	//实践环节
	beego.Router("/practice", &controllers.PracticeController{})
	beego.Router("/practice/edit", &controllers.PracticeController{}, "*:PracticeEdit")
	beego.Router("/practice/add", &controllers.PracticeController{}, "*:PracticeAdd")
	beego.Router("/practice/add/action", &controllers.PracticeController{}, "*:PracticeAddAction")
	beego.Router("/practice/search", &controllers.PracticeController{}, "*:PracticeSearch")
	beego.Router("/practice/updata", &controllers.PracticeController{}, "*:PracticeUpdata")
	beego.Router("/practice/delete", &controllers.PracticeController{}, "*:PracticeDelete")
	beego.AutoRouter(&controllers.PracticeController{})

	//专业信息
	beego.Router("/pm", &controllers.PmController{})
	beego.Router("/pm/edit", &controllers.PmController{}, "*:PmEdit")
	beego.Router("/pm/add", &controllers.PmController{}, "*:PmAdd")
	beego.Router("/pm/add/action", &controllers.PmController{}, "*:PmAddAction")
	beego.Router("/pm/search", &controllers.PmController{}, "*:PmSearch")
	beego.Router("/pm/updata", &controllers.PmController{}, "*:PmUpdata")
	beego.Router("/pm/delete", &controllers.PmController{}, "*:PmDelete")
	beego.AutoRouter(&controllers.PmController{})

	//专业理论
	beego.Router("/ptcourse", &controllers.PTCourseController{})
	beego.Router("/ptcourse/edit", &controllers.PTCourseController{}, "*:PTCourseEdit")
	beego.Router("/ptcourse/search", &controllers.PTCourseController{}, "*:PTCourseSearch")
	beego.Router("/ptcourse/setcourse", &controllers.PTCourseController{}, "*:Setcourse")
	beego.Router("/ptcourse/delete", &controllers.PTCourseController{}, "*:PTCourseDelete")
	beego.AutoRouter(&controllers.PTCourseController{})

	//专业实践环节
	beego.Router("/pplink", &controllers.PPLinkController{})
	beego.Router("/pplink/edit", &controllers.PPLinkController{}, "*:PPLinkEdit")
	beego.Router("/pplink/search", &controllers.PPLinkController{}, "*:PPLinkSearch")
	beego.Router("/pplink/setcourse", &controllers.PPLinkController{}, "*:Setcourse")
	beego.Router("/pplink/delete", &controllers.PPLinkController{}, "*:PPLinkDelete")
	beego.AutoRouter(&controllers.PPLinkController{})

	//专业学分
	beego.Router("/pgcredit", &controllers.PGCreditController{})
	beego.Router("/pgcredit/add", &controllers.PGCreditController{}, "*:PgcAdd")
	beego.Router("/pgcredit/edit", &controllers.PGCreditController{}, "*:PgcEdit")
	beego.Router("/pgcredit/search", &controllers.PGCreditController{}, "*:PgcSearch")
	beego.Router("/pgcredit/save", &controllers.PGCreditController{}, "*:PgcSave")
	beego.Router("/pgcredit/updata", &controllers.PGCreditController{}, "*:PgcUpdate")
	beego.Router("/pgcredit/delete", &controllers.PGCreditController{}, "*:PgcDel")
	beego.AutoRouter(&controllers.PGCreditController{})

	//复制计划
	beego.Router("/copyplan/profession", &controllers.CopyPlanController{}, "*:GetProfessionPlan")
	beego.Router("/copyplan/year", &controllers.CopyPlanController{}, "*:GetYearPlan")
	beego.Router("/copyplan/profession/search", &controllers.CopyPlanController{}, "*:PPSearch")
	beego.Router("/copyplan/profession/copy", &controllers.CopyPlanController{}, "*:PPCopy")
	beego.Router("/copyplan/profession/remove", &controllers.CopyPlanController{}, "*:PPRemove")
	beego.Router("/copyplan/year/search", &controllers.CopyPlanController{}, "*:GYSearch")
	beego.Router("/copyplan/year/copy", &controllers.CopyPlanController{}, "*:GYCopy")
	beego.Router("/copyplan/year/remove", &controllers.CopyPlanController{}, "*:GYRemove")
	beego.AutoRouter(&controllers.CopyPlanController{})
}
