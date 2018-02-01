package controllers

import (
	"CoursePj/models"
	"fmt"

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
	tc := new(models.TheoryCourse)
	theoryCourse := models.TheoryCourse{}
	if err := this.ParseForm(&theoryCourse); err != nil {
		fmt.Println("获取表单数据失败")
		this.ajaxMsg("获取表单数据失败", MSG_ERR_Param)
	}
	//获取year和cid
	year := theoryCourse.Year
	cid := theoryCourse.Cid
	fmt.Println("获取cid_info", year, cid)
	fmt.Println("获取表单数据成功")
	if year == 0 || cid == "" {
		this.ajaxMsg("year，cid不能为空", MSG_ERR_Param)
	}
	o := orm.NewOrm()
	exist := o.QueryTable(tc).Filter("Cid", cid).Filter("Year", year).Exist()
	if exist {
		fmt.Println("插入失败,存在相同项")
		this.ajaxMsg("插入失败,存在相同项", MSG_ERR_Resources)
	} else {
		if _, err := o.Insert(&theoryCourse); err != nil {
			fmt.Println("插入失败")
			this.ajaxMsg("新增失败", MSG_ERR_Resources)
		}
	}
	list := make(map[string]interface{})
	list["id"] = theoryCourse.Id
	this.ajaxList("新增成功", MSG_OK, 1, list)
	//this.ajaxMsg("新增成功", MSG_OK)
	return
}

//检索
func (this *HomeController) TheoryCourseSearch() {
	fmt.Println("post查询")

	filters := make([]interface{}, 0)
	Cunit := this.Input().Get("s_Cunit")
	if Cunit != "" {
		filters = append(filters, "Cunit", Cunit)
	}
	Cname := this.Input().Get("s_Cname")
	if Cname != "" {
		filters = append(filters, "Cname", Cname)
	}

	Ccg1 := this.Input().Get("s_Ccg1")
	if Ccg1 != "" {
		filters = append(filters, "Ccg1", Ccg1)
	}
	Ccg2 := this.Input().Get("s_Ccg2")
	if Ccg2 != "" {
		filters = append(filters, "Ccg2", Ccg2)
	}
	Status := this.Input().Get("s_Status")
	if Status != "" {
		filters = append(filters, "Status", Status)
	}
	this.Data["s"] = Status
	Year := this.Input().Get("s_Year")
	if Year != "" {
		filters = append(filters, "Year", Year)
	}
	this.Data["y"] = Year

	fmt.Println(Cunit)
	fmt.Println(len(Cname))
	o := orm.NewOrm()
	var maps []orm.Params
	theoryCourse := new(models.TheoryCourse)
	query := o.QueryTable(theoryCourse)

	if len(filters) > 0 {
		l := len(filters)
		for k := 0; k < l; k += 2 {
			query = query.Filter(filters[k].(string), filters[k+1])
		}
	}
	num, err := query.Values(&maps)
	if err == nil {
		fmt.Printf("Result Nums: %d\n", num)
		this.Data["m"] = maps
		this.Data["num"] = num
		for _, m := range maps {
			fmt.Println(m["Punit"], m["Pid"])
		}
	}
	this.TplName = "home.tpl"
}

//编辑
func (this *HomeController) TheoryCourseEdit() {
	fmt.Println("编辑")
	cid := this.Input().Get("cid")
	fmt.Println(cid)
	year := this.Input().Get("year")
	fmt.Println(year)
	o := orm.NewOrm()
	var maps []orm.Params
	theoryCourse := new(models.TheoryCourse)
	num, err := o.QueryTable(theoryCourse).Filter("Cid", cid).Filter("Year", year).Values(&maps)
	if err == nil {
		fmt.Printf("Result Nums: %d\n", num)
		this.Data["m"] = maps
		for _, m := range maps {
			fmt.Println(m["Cid"])
			this.Data["y"] = year
			this.Data["s"] = m["Status"]
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
		this.ajaxMsg("获取表单数据失败", MSG_ERR_Param)
	}
	id := theoryCourse.Id
	if id == 0 {
		this.ajaxMsg("id 不能为空", MSG_ERR_Param)
	}
	num, err := orm.NewOrm().Update(&theoryCourse)
	if err != nil {
		fmt.Println("更新失败")
		this.ajaxMsg("更新失败", MSG_ERR_Resources)
	}
	if num == 0 {
		this.ajaxMsg("更新失败", MSG_ERR_Resources)
	}
	this.ajaxMsg("更新成功", MSG_OK)
	return
}

//删除
func (this *HomeController) TheoryCourseDelete() {
	fmt.Println("删除")
	o := orm.NewOrm()
	theoryCourse := new(models.TheoryCourse)
	cid := this.Input().Get("cid")
	//var tc models.TheoryCourse
	//json.Unmarshal(this.Ctx.Input.RequestBody, &tc)
	//fmt.Println("cid",tc.Cid)
	//cid := this.GetString("cid")
	//cid := tc.Cid
	fmt.Println(cid)
	year := this.Input().Get("year")
	//year := tc.Year
	fmt.Println(year)
	if year == "" || cid == "" {
		this.ajaxMsg("cid，year不能为空", MSG_ERR_Param)
	}
	num, err := o.QueryTable(theoryCourse).Filter("Cid", cid).Filter("Year", year).Delete()
	if err == nil {
		fmt.Printf("删除成功")
		fmt.Printf("Result Nums: %d\n", num)
		if num == 0 {
			this.ajaxMsg("删除失败", MSG_ERR_Resources)
		}
	} else {
		this.ajaxMsg("删除失败", MSG_ERR_Resources)
	}
	this.ajaxMsg("删除成功", MSG_OK)
	return
}
