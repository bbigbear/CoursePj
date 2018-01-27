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

	var slice []string
	//可用
	num, err := query.Values(&maps)
	if err == nil {
		fmt.Printf("Result Nums: %d\n", num)
		this.Data["m"] = maps
		this.Data["num"] = num
		for _, m := range maps {
			faculty := fmt.Sprint(m["Faculty"])
			slice = append(slice, faculty)
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
	this.Data["f"] = this.RemoveRepBySlice(slice)

	this.TplName = "pm.tpl"
}

//展示
func (this *PmController) PmAdd() {
	this.TplName = "addPm.tpl"

}

//新建
func (this *PmController) PmAddAction() {
	fmt.Println("post新建")
	m := new(models.Pm)
	pm := models.Pm{}

	if err := this.ParseForm(&pm); err != nil {
		fmt.Println("获取表单数据失败")
	}
	//获取year和Pmid
	year := pm.Year
	pmid := pm.Pmid
	fmt.Println("获取pmid_info", year, pmid)
	fmt.Println("获取表单数据成功")
	o := orm.NewOrm()
	exist := o.QueryTable(m).Filter("Pmid", pmid).Filter("Year", year).Exist()
	if exist {
		fmt.Println("插入失败,已存在该专业")
		this.ajaxMsg("", MSG_ERR)
	} else {
		if _, err := o.Insert(&pm); err != nil {
			fmt.Println("插入失败")
			this.ajaxMsg("", MSG_ERR)
		}
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

	//院系下拉框
	query2 := o.QueryTable(pm).Filter("Status", "可用")
	var slice []string
	num, err1 := query2.Values(&maps)
	if err1 == nil {
		fmt.Printf("Result Nums: %d\n", num)
		for _, m := range maps {
			//获取院系
			faculty := fmt.Sprint(m["Faculty"])
			slice = append(slice, faculty)

		}
		this.Data["f"] = this.RemoveRepBySlice(slice)
	}

	if len(filters) > 0 {
		l := len(filters)
		for k := 0; k < l; k += 2 {
			query = query.Filter(filters[k].(string), filters[k+1])
			query1 = query1.Filter(filters[k].(string), filters[k+1])
		}
	}
	//var slice []string
	//可用
	num, err := query.Values(&maps)
	if err == nil {
		fmt.Printf("Result Nums: %d\n", num)
		this.Data["m"] = maps
		this.Data["num"] = num
		for _, m := range maps {
			//faculty := fmt.Sprint(m["Faculty"])
			//slice = append(slice, faculty)
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
	//this.Data["f"] = this.RemoveRepBySlice(slice)

	this.TplName = "pm.tpl"

}

//编辑
func (this *PmController) PmEdit() {
	fmt.Println("编辑")
	pmid := this.Input().Get("pmid")
	fmt.Println(pmid)
	year := this.Input().Get("year")
	fmt.Println(year)
	o := orm.NewOrm()
	var maps []orm.Params
	pm := new(models.Pm)
	num, err := o.QueryTable(pm).Filter("Pmid", pmid).Filter("Year", year).Values(&maps)
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
	year := this.Input().Get("year")
	fmt.Println(year)
	num, err := o.QueryTable(pm).Filter("Pmid", pmid).Filter("Year", year).Delete()
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
