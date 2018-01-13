package controllers

import (
	"CoursePj/models"
	"fmt"
	"strings"

	_ "github.com/GO-SQL-Driver/MySQL"
	"github.com/astaxie/beego/orm"
)

type HomeController struct {
	BaseController
}

func (this *HomeController) Get() {

	o := orm.NewOrm()
	var maps []orm.Params
	theoryCourse := new(models.TheoryCourse)
	num, err := o.QueryTable(theoryCourse).Values(&maps)
	if err == nil {
		fmt.Printf("Result Nums: %d\n", num)
		this.Data["m"] = maps
		this.Data["num"] = num
	}
	this.TplName = "home.tpl"
}

//展示
func (this *HomeController) TheoryCourseAdd() {
	this.TplName = "addTheoryCourse.tpl"

}

//新建
func (this *HomeController) TheoryCourseAddAction() {
	fmt.Println("post新建")
	theoryCourse := models.TheoryCourse{}

	if err := this.ParseForm(&theoryCourse); err != nil {
		fmt.Println("获取表单数据失败")
	}
	fmt.Println("获取表单数据成功")
	if _, err := orm.NewOrm().Insert(&theoryCourse); err != nil {
		fmt.Println("插入失败")
		this.ajaxMsg("", MSG_ERR)
	}
	this.ajaxMsg("", MSG_OK)
	return
}

//检索
func (this *HomeController) TheoryCourseSearch() {
	fmt.Println("post查询")
	//	Cunit := strings.TrimSpace(this.GetString("s_Cunit"))
	//	Cname := strings.TrimSpace(this.GetString("s_Cname"))
	Cunit := this.Input().Get("s_Cunit")
	Cname := this.Input().Get("s_Cname")
	fmt.Println(Cunit)
	fmt.Println(Cname)
	o := orm.NewOrm()
	var maps []orm.Params
	theoryCourse := new(models.TheoryCourse)
	if (Cunit != "") && (Cname == "") {
		num, err := o.QueryTable(theoryCourse).Filter("Cunit", Cunit).Values(&maps)
		if err == nil {
			fmt.Printf("Result Nums: %d\n", num)
			this.Data["m"] = maps
			this.Data["num"] = num
			for _, m := range maps {
				fmt.Println(m["Cunit"], m["Cid"])
			}
		}
	} else if (Cunit == "") && (Cname != "") {
		num, err := o.QueryTable(theoryCourse).Filter("Cname", Cname).Values(&maps)
		if err == nil {
			fmt.Printf("Result Nums: %d\n", num)
			this.Data["m"] = maps
			this.Data["num"] = num
			for _, m := range maps {
				fmt.Println(m["Cunit"], m["Cid"])
			}
		}
	} else {
		num, err := o.QueryTable(theoryCourse).Filter("Cunit", Cunit).Filter("Cname", Cname).Values(&maps)
		if err == nil {
			fmt.Printf("Result Nums: %d\n", num)
			this.Data["m"] = maps
			this.Data["num"] = num
			for _, m := range maps {
				fmt.Println(m["Cunit"], m["Cid"])
			}
		}
	}
	this.TplName = "home.tpl"
}

//编辑
func (this *HomeController) TheoryCourseEdit() {
	fmt.Println("编辑")
	cid := this.Input().Get("cid")
	fmt.Println(cid)
	o := orm.NewOrm()
	var maps []orm.Params
	theoryCourse := new(models.TheoryCourse)
	num, err := o.QueryTable(theoryCourse).Filter("Cid", cid).Values(&maps)
	if err == nil {
		fmt.Printf("Result Nums: %d\n", num)
		this.Data["m"] = maps
		for _, m := range maps {
			fmt.Println(m["Cid"])
		}
	}
	this.TplName = "editTheoryCourse.tpl"
	return
}

//更新
func (this *HomeController) TheoryCourseUpdata() {
	fmt.Println("更新")
	theoryCourse := models.TheoryCourse{}

	if err := this.ParseForm(&theoryCourse); err != nil {
		fmt.Println("获取表单数据失败")
	}

	if _, err := orm.NewOrm().Update(&theoryCourse); err != nil {
		fmt.Println("更新失败")
		this.ajaxMsg("", MSG_ERR)
	}
	this.ajaxMsg("", MSG_OK)
	return
}

//删除
func (this *HomeController) TheoryCourseDelete() {
	fmt.Println("删除")
	o := orm.NewOrm()
	theoryCourse := new(models.TheoryCourse)
	cid := strings.TrimSpace(this.GetString("cid"))
	fmt.Println(cid)
	num, err := o.QueryTable(theoryCourse).Filter("Cid", cid).Delete()
	if err == nil {
		fmt.Printf("删除成功")
		fmt.Printf("Result Nums: %d\n", num)
	}
	this.ajaxMsg("", MSG_OK)
	return
}
