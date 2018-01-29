package controllers

import (
	"CoursePj/models"
	"fmt"
	"regexp"
	"strings"

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
	var slice []string
	num1, err := query1.Values(&maps)
	if err == nil {
		fmt.Printf("Result Nums: %d\n", num)
		this.Data["m1"] = maps
		this.Data["num1"] = num1
		for _, m1 := range maps {
			faculty := fmt.Sprint(m1["Faculty"])
			slice = append(slice, faculty)
			fmt.Println(m1["Year"], m1["Faculty"], m1["Status"])
		}
	}
	this.Data["f"] = this.RemoveRepBySlice(slice)

	this.TplName = "pplink.tpl"
}

//检索
func (this *PPLinkController) PPLinkSearch() {
	o := orm.NewOrm()
	var maps []orm.Params
	pm := new(models.Pm)
	year := this.Input().Get("year")
	this.Data["y"] = year
	faculty := this.Input().Get("faculty")
	this.Data["f"] = faculty
	practice := new(models.Practice)
	query := o.QueryTable(practice).Filter("Status", "可用").Filter("Year", year)
	query1 := o.QueryTable(pm).Filter("Status", "可用").Filter("Year", year).Filter("Faculty", faculty)
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

//专业设置课程
func (this *PPLinkController) Setcourse() {
	fmt.Println("接收到课程")
	pid := this.Input().Get("pid")
	pmid := this.Input().Get("pmid")
	pidlist := strings.Split(pid, ",")
	//	reg := regexp.MustCompile(`^\[.*\]`)
	reg := regexp.MustCompile(`[[:digit:]]+`)
	pmidlist := reg.FindAllString(pmid, -1)
	fmt.Println(pmid)
	fmt.Println(pidlist)
	fmt.Println(pmidlist)

	o := orm.NewOrm()
	pid_count := len(pidlist) - 1
	pmid_count := len(pmidlist)
	for j := 0; j < pmid_count; j++ {
		for i := 0; i < pid_count; i++ {

			var pplink models.Pplink
			pp := new(models.Pplink)
			//ci, err := strconv.ParseInt(pidlist[i], 10, 64)
			//if err == nil {
			ci := pidlist[i]
			pplink.Pid = ci
			//}
			//pmi, err := strconv.ParseInt(pmidlist[j], 10, 64)
			pmi := pmidlist[j]
			//if err == nil {
			pplink.Pmid = pmi
			//}
			//先查询再建立
			num, err := o.QueryTable(pp).Filter("Pid", ci).Filter("Pmid", pmi).Count()
			if err != nil {
				this.ajaxMsg("", MSG_ERR)
				fmt.Println("查询失败")
			}
			fmt.Println("query num:", num)
			if num == 0 {
				id, err := o.Insert(&pplink)
				if err != nil {
					this.ajaxMsg("", MSG_ERR)
					fmt.Println("插入失败")
				}
				fmt.Println(id)
			} else {
				this.ajaxMsg("", MSG_ERR)
			}

		}
	}
	fmt.Println("插入成功")
	this.ajaxMsg("", MSG_OK)
	return

}

//查看已设置课程
func (this *PPLinkController) PPLinkEdit() {
	fmt.Println("查看专业课程")
	pmid := this.Input().Get("pmid")
	fmt.Println(pmid)
	year := this.Input().Get("year")
	this.Data["y"] = year
	//取其中的pmid
	reg := regexp.MustCompile(`[[:digit:]]+`)
	pmidlist := reg.FindAllString(pmid, -1)

	o := orm.NewOrm()
	var maps []orm.Params
	ppl := new(models.Pplink)
	practice := new(models.Practice)
	var practice_info models.Practice

	num, err := o.QueryTable(ppl).Filter("Pmid", pmidlist[0]).Values(&maps)
	if err == nil {
		fmt.Printf("Result Nums: %d\n", num)
		//		slice := make([]string, num)
		var slice []string
		for _, m := range maps {
			fmt.Println("map_cid:", m["Pid"])
			err := o.QueryTable(practice).Filter("Pid", m["Pid"]).Filter("Year", year).One(&practice_info)
			if err == nil {
				fmt.Println("tc_info:", practice_info.Pname)
				slice = append(slice, practice_info.Pname)
			}

		}
		fmt.Println("slice:", slice)
		fmt.Println("len slice:", len(slice))
		this.Data["m"] = maps
		this.Data["pmid"] = pmid
		this.Data["s"] = slice
	}
	this.TplName = "editPplink.tpl"
	return
}

//删除课程
func (this *PPLinkController) PPLinkDelete() {
	fmt.Println("删除")
	pname := this.Input().Get("pname")
	pnamelist := strings.Split(pname, ",")
	pname_count := len(pnamelist) - 1
	fmt.Println("pname del:", pname, pnamelist)

	pmid := this.Input().Get("pmid")
	fmt.Println("pmid:", pmid)
	year := this.Input().Get("year")

	//取其中的pmid
	reg := regexp.MustCompile(`[[:digit:]]+`)
	pmidlist := reg.FindAllString(pmid, -1)

	o := orm.NewOrm()
	practice := new(models.Practice)
	ppl := new(models.Pplink)
	var practice_info models.Practice

	for i := 0; i < pname_count; i++ {
		err := o.QueryTable(practice).Filter("Pname", pnamelist[i]).Filter("Year", year).One(&practice_info)
		if err == nil {
			num, err := o.QueryTable(ppl).Filter("Pid", practice_info.Pid).Filter("Pmid", pmidlist[0]).Delete()
			if err == nil {
				fmt.Printf("删除成功")
				fmt.Printf("Result delPid Nums: %d\n", num)
			}
		}
	}

	this.ajaxMsg("", MSG_OK)
	return

}
