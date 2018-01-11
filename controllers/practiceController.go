package controllers

import (
	"CoursePj/models"
	"fmt"
	"strings"

	_ "github.com/GO-SQL-Driver/MySQL"
	"github.com/astaxie/beego/orm"
)

type PracticeController struct {
	BaseController
}

func (this *PracticeController) Get() {

	o := orm.NewOrm()
	var maps []orm.Params
	practice := new(models.Practice)
	num, err := o.QueryTable(practice).Values(&maps)
	if err == nil {
		fmt.Printf("Result Nums: %d\n", num)
		this.Data["m"] = maps
		this.Data["num"] = num
	}
	this.TplName = "practice.tpl"
}

//展示
func (this *PracticeController) PracticeAdd() {
	this.TplName = "addPractice.tpl"

}

//新建
func (this *PracticeController) PracticeAddAction() {
	fmt.Println("post新建")
	practice := models.Practice{}

	if err := this.ParseForm(&practice); err != nil {
		fmt.Println("获取表单数据失败")
	}
	fmt.Println("获取表单数据成功")
	if _, err := orm.NewOrm().Insert(&practice); err != nil {
		fmt.Println("插入失败")
		this.ajaxMsg("", MSG_ERR)
	}
	this.ajaxMsg("", MSG_OK)
	return
}

//检索
func (this *PracticeController) PracticeSearch() {
	fmt.Println("post查询")
	Punit := this.Input().Get("s_Punit")
	Pname := this.Input().Get("s_Pname")
	fmt.Println(Punit)
	fmt.Println(Pname)
	o := orm.NewOrm()
	var maps []orm.Params
	practice := new(models.Practice)
	num, err := o.QueryTable(practice).Filter("Punit", Punit).Filter("Pname", Pname).Values(&maps)
	if err == nil {
		fmt.Printf("Result Nums: %d\n", num)
		this.Data["m"] = maps
		this.Data["num"] = num
		for _, m := range maps {
			fmt.Println(m["Punit"], m["Pid"])
		}
	}
	this.TplName = "practice.tpl"

}

//编辑
func (this *PracticeController) PracticeEdit() {
	fmt.Println("编辑")
	pid := this.Input().Get("pid")
	fmt.Println(pid)
	o := orm.NewOrm()
	var maps []orm.Params
	practice := new(models.Practice)
	num, err := o.QueryTable(practice).Filter("Pid", pid).Values(&maps)
	if err == nil {
		fmt.Printf("Result Nums: %d\n", num)
		this.Data["m"] = maps
		for _, m := range maps {
			fmt.Println(m["Pid"])
		}
	}
	this.TplName = "editPractice.tpl"
	return
}

//更新
func (this *PracticeController) PracticeUpdata() {
	fmt.Println("更新")
	practice := models.Practice{}

	if err := this.ParseForm(&practice); err != nil {
		fmt.Println("获取表单数据失败")
	}

	if _, err := orm.NewOrm().Update(&practice); err != nil {
		fmt.Println("更新失败")
		this.ajaxMsg("", MSG_ERR)
	}
	this.ajaxMsg("", MSG_OK)
	return
}

//删除
func (this *PracticeController) PracticeDelete() {
	fmt.Println("删除")
	o := orm.NewOrm()
	practice := new(models.Practice)
	pid := strings.TrimSpace(this.GetString("pid"))
	fmt.Println(pid)
	num, err := o.QueryTable(practice).Filter("Pid", pid).Delete()
	if err == nil {
		fmt.Printf("删除成功")
		fmt.Printf("Result Nums: %d\n", num)
	}
	this.ajaxMsg("", MSG_OK)
	return
}
