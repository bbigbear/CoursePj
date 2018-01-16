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

	o := orm.NewOrm()
	var maps []orm.Params
	pm := new(models.Pm)
	query := o.QueryTable(pm).Filter("Status", "可用")
	query1 := o.QueryTable(pm).Filter("Status", "停用")

	//可用
	num, err := query.Values(&maps)
	if err == nil {
		fmt.Printf("Result Nums: %d\n", num)
		this.Data["m"] = maps
		this.Data["num"] = num
		for _, m := range maps {
			fmt.Println(m["Year"], m["Faculty"], m["Status"])
		}
	}
	//停用
	num1, err := query1.Values(&maps)
	if err == nil {
		fmt.Printf("Result Nums: %d\n", num)
		this.Data["m1"] = maps
		this.Data["num1"] = num1
		for _, m1 := range maps {
			fmt.Println(m1["Year"], m1["Faculty"], m1["Status"])
		}
	}

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

	filters := make([]interface{}, 0)
	Pmyear := this.Input().Get("s_Pmyear")
	if Pmyear != "" {
		filters = append(filters, "Year", Pmyear)
	}
	Pmfaculty := this.Input().Get("s_Pmfaculty")
	if Pmfaculty != "" {
		filters = append(filters, "Faculty", Pmfaculty)
	}

	fmt.Println(Pmyear)
	fmt.Println(len(Pmfaculty))
	o := orm.NewOrm()
	var maps []orm.Params
	pm := new(models.Pm)
	query := o.QueryTable(pm).Filter("Status", "可用")
	query1 := o.QueryTable(pm).Filter("Status", "停用")

	if len(filters) > 0 {
		l := len(filters)
		for k := 0; k < l; k += 2 {
			query = query.Filter(filters[k].(string), filters[k+1])
			query1 = query1.Filter(filters[k].(string), filters[k+1])
		}
	}
	//可用
	num, err := query.Values(&maps)
	if err == nil {
		fmt.Printf("Result Nums: %d\n", num)
		this.Data["m"] = maps
		this.Data["num"] = num
		for _, m := range maps {
			fmt.Println(m["Year"], m["Faculty"], m["Status"])
		}
	}
	//停用
	num1, err := query1.Values(&maps)
	if err == nil {
		fmt.Printf("Result Nums: %d\n", num)
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
	pmid := this.Input().Get("pm1id")
	fmt.Println(pmid)
	//	var req []byte = this.Ctx.Input.RequestBody
	//	fmt.Println(string(req[:]))
}
