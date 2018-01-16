package controllers

import (
	"CoursePj/models"
	"fmt"
	"regexp"
	"strings"

	_ "github.com/GO-SQL-Driver/MySQL"
	"github.com/astaxie/beego/orm"
)

type PTCourseController struct {
	BaseController
}

func (this *PTCourseController) Get() {

	o := orm.NewOrm()
	var maps []orm.Params
	pm := new(models.Pm)
	theoryCourse := new(models.TheoryCourse)
	query := o.QueryTable(theoryCourse).Filter("Status", "可用")
	query1 := o.QueryTable(pm).Filter("Status", "可用")

	//理论
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

	this.TplName = "ptcourse.tpl"
}

//专业设置课程
func (this *PTCourseController) Setcourse() {
	fmt.Println("接收到课程")
	cid := this.Input().Get("cid")
	pmid := this.Input().Get("pmid")
	cidlist := strings.Split(cid, ",")
	reg := regexp.MustCompile(`\["(.*?)"\]`)
	pmidlist := reg.FindAllString(pmid, -1)
	fmt.Println(pmid)
	fmt.Println(cidlist)
	fmt.Println(pmidlist)
}
