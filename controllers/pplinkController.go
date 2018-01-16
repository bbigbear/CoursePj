package controllers

import (
	"CoursePj/models"
	"fmt"

	_ "github.com/GO-SQL-Driver/MySQL"
	"github.com/astaxie/beego/orm"
)

type PPLinkController struct {
	BaseController
}

func (this *PPLinkController) Get() {

	o := orm.NewOrm()
	var maps []orm.Params
	pm := new(models.Pm)
	practice := new(models.Practice)
	query := o.QueryTable(practice).Filter("Status", "可用")
	query1 := o.QueryTable(pm).Filter("Status", "可用")

	//实践环节
	num, err := query.Values(&maps)
	if err == nil {
		fmt.Printf("Result Nums: %d\n", num)
		this.Data["m"] = maps
		this.Data["num"] = num
		for _, m := range maps {
			fmt.Println(m["Year"], m["Faculty"], m["Status"])
		}
	}
	//专业
	num1, err := query1.Values(&maps)
	if err == nil {
		fmt.Printf("Result Nums: %d\n", num)
		this.Data["m1"] = maps
		this.Data["num1"] = num1
		for _, m1 := range maps {
			fmt.Println(m1["Year"], m1["Faculty"], m1["Status"])
		}
	}

	this.TplName = "pplink.tpl"
}
