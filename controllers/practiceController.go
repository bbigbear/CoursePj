package controllers

import (
	"CoursePj/models"
	"fmt"

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
	p := new(models.Practice)
	practice := models.Practice{}

	if err := this.ParseForm(&practice); err != nil {
		fmt.Println("获取表单数据失败")
		this.ajaxMsg("获取表单数据失败", MSG_ERR_Param)
	}
	fmt.Println("获取表单数据成功")
	o := orm.NewOrm()
	pid := practice.Pid
	year := practice.Year
	fmt.Println("获取pid_info", year, pid)
	if year == 0 || pid == "" {
		this.ajaxMsg("year，pid不能为空", MSG_ERR_Param)
	}
	exist := o.QueryTable(p).Filter("Pid", pid).Filter("Year", year).Exist()
	if exist {
		fmt.Println("插入失败,存在相同项")
		this.ajaxMsg("新增失败,存在相同项", MSG_ERR_Resources)
	} else {
		if _, err := o.Insert(&practice); err != nil {
			fmt.Println("插入失败")
			this.ajaxMsg("新增失败", MSG_ERR_Resources)
		}
	}

	list := make(map[string]interface{})
	list["id"] = practice.Id
	this.ajaxList("新增成功", MSG_OK, 1, list)
	//this.ajaxMsg("新增成功", MSG_OK)
	return
}

//检索
func (this *PracticeController) PracticeSearch() {
	fmt.Println("post查询")

	filters := make([]interface{}, 0)
	Punit := this.Input().Get("s_Punit")
	if Punit != "" {
		filters = append(filters, "Punit", Punit)
	}
	Pname := this.Input().Get("s_Pname")
	if Pname != "" {
		filters = append(filters, "Pname", Pname)
	}
	Pcg1 := this.Input().Get("s_Pcg1")
	if Pcg1 != "" {
		filters = append(filters, "Pcg1", Pcg1)
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

	fmt.Println(Punit)
	fmt.Println(len(Pname))
	o := orm.NewOrm()
	var maps []orm.Params
	practice := new(models.Practice)
	query := o.QueryTable(practice)

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

	this.TplName = "practice.tpl"

}

//编辑
func (this *PracticeController) PracticeEdit() {
	fmt.Println("编辑")
	pid := this.Input().Get("pid")
	fmt.Println(pid)
	year := this.Input().Get("year")
	fmt.Println(year)
	o := orm.NewOrm()
	var maps []orm.Params
	practice := new(models.Practice)
	num, err := o.QueryTable(practice).Filter("Pid", pid).Filter("Year", year).Values(&maps)
	if err == nil {
		fmt.Printf("Result Nums: %d\n", num)
		this.Data["m"] = maps
		for _, m := range maps {
			fmt.Println(m["Pid"])
			this.Data["y"] = year
			this.Data["s"] = m["Status"]
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
		this.ajaxMsg("获取表单数据失败", MSG_ERR_Param)
	}
	id := practice.Id
	if id == 0 {
		this.ajaxMsg("id 不能为空", MSG_ERR_Param)
	}
	num, err := orm.NewOrm().Update(&practice)
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
func (this *PracticeController) PracticeDelete() {
	fmt.Println("删除")
	o := orm.NewOrm()
	practice := new(models.Practice)
	pid := this.Input().Get("pid")
	//var p models.Practice
	//json.Unmarshal(this.Ctx.Input.RequestBody, &p)
	//pid := p.Pid
	fmt.Println(pid)
	year := this.Input().Get("year")
	//year := p.Year
	fmt.Println(year)
	if year == "" || pid == "" {
		this.ajaxMsg("pid,year不能为空", MSG_ERR_Param)
	}
	num, err := o.QueryTable(practice).Filter("Pid", pid).Filter("Year", year).Delete()
	if err == nil {
		fmt.Printf("删除成功")
		fmt.Printf("Result Nums: %d\n", num)
		if num == 0 {
			this.ajaxMsg("删除失败,找不到对应的环节", MSG_ERR_Resources)
		}
	} else {
		this.ajaxMsg("删除失败", MSG_ERR_Resources)
	}
	this.ajaxMsg("删除成功", MSG_OK)
	return
}
