package controllers

import (
	"CoursePj/models"
	"fmt"
	"strings"

	_ "github.com/GO-SQL-Driver/MySQL"
	"github.com/astaxie/beego/orm"
)

type PmController struct {
	BaseController
}

func (this *PmController) Get() {

	this.TplName = "pm.tpl"
}

//展示
func (this *PmController) PmAdd() {
	this.TplName = "addPm.tpl"

}

//新建
func (this *PmController) PmAddAction() {
	fmt.Println("post新建")
	pm := models.Pm{}

	if err := this.ParseForm(&pm); err != nil {
		fmt.Println("获取表单数据失败")
	}
	fmt.Println("获取表单数据成功")
	if _, err := orm.NewOrm().Insert(&pm); err != nil {
		fmt.Println("插入失败")
		this.ajaxMsg("", MSG_ERR)
	}
	this.ajaxMsg("", MSG_OK)
	return
}

//检索
func (this *PmController) PmSearch() {
	fmt.Println("post查询")
	//	Cunit := strings.TrimSpace(this.GetString("s_Cunit"))
	//	Cname := strings.TrimSpace(this.GetString("s_Cname"))
	Pmyear := this.Input().Get("s_Pmyear")
	Pmfaculty := this.Input().Get("s_Pmfaculty")
	fmt.Println(Pmyear)
	fmt.Println(Pmfaculty)
	o := orm.NewOrm()
	var maps []orm.Params
	pm := new(models.Pm)
	num, err := o.QueryTable(pm).Filter("Year", Pmyear).Filter("Faculty", Pmfaculty).Filter("Status", "可用").Values(&maps)
	if err == nil {
		fmt.Printf("Result Nums: %d\n", num)
		this.Data["m"] = maps
		this.Data["num"] = num
		for _, m := range maps {
			fmt.Println(m["Year"], m["Faculty"], m["Status"])
		}
	}
	num1, err := o.QueryTable(pm).Filter("Year", Pmyear).Filter("Faculty", Pmfaculty).Filter("Status", "停用").Values(&maps)
	if err == nil {
		fmt.Printf("Result Nums: %d\n", num1)
		this.Data["m1"] = maps
		this.Data["num1"] = num1
		for _, m1 := range maps {
			fmt.Println(m1["Year"], m1["Faculty"], m1["Status"])
		}
	}
	this.TplName = "pm.tpl"

}

//编辑
func (this *PmController) PmEdit() {
	fmt.Println("编辑")
	pmid := this.Input().Get("pmid")
	fmt.Println(pmid)
	o := orm.NewOrm()
	var maps []orm.Params
	pm := new(models.Pm)
	num, err := o.QueryTable(pm).Filter("Pmid", pmid).Values(&maps)
	if err == nil {
		fmt.Printf("Result Nums: %d\n", num)
		this.Data["m"] = maps
		for _, m := range maps {
			fmt.Println(m["Pmid"])
		}
	}
	this.TplName = "editPm.tpl"
	return
}

//更新
func (this *PmController) PmUpdata() {
	fmt.Println("更新")
	pm := models.Pm{}

	if err := this.ParseForm(&pm); err != nil {
		fmt.Println("获取表单数据失败")
	}

	if _, err := orm.NewOrm().Update(&pm); err != nil {
		fmt.Println("更新失败")
		this.ajaxMsg("", MSG_ERR)
	}
	this.ajaxMsg("", MSG_OK)
	return
}

//删除
func (this *PmController) PmDelete() {
	fmt.Println("删除")
	o := orm.NewOrm()
	pm := new(models.Pm)
	pmid := strings.TrimSpace(this.GetString("pmid"))
	fmt.Println(pmid)
	num, err := o.QueryTable(pm).Filter("Pmid", pmid).Delete()
	if err == nil {
		fmt.Printf("删除成功")
		fmt.Printf("Result Nums: %d\n", num)
	}
	this.ajaxMsg("", MSG_OK)
	return
}

//状态改变
func (this *PmController) PmStautsChange() {
	fmt.Println("状态改变")
	pmid := this.Input().Get("pmid")
	fmt.Println(pmid)

}
