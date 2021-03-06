package controllers

import (
	"CoursePj/models"
	"fmt"
	"regexp"
	"strconv"
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
	//	reg := regexp.MustCompile(`^\[.*\]`)
	reg := regexp.MustCompile(`[[:digit:]]+`)
	pmidlist := reg.FindAllString(pmid, -1)
	fmt.Println(pmid)
	fmt.Println(cidlist)
	fmt.Println(pmidlist)

	o := orm.NewOrm()
	cid_count := len(cidlist) - 1
	pmid_count := len(pmidlist)
	for j := 0; j < pmid_count; j++ {
		for i := 0; i < cid_count; i++ {

			var ptcourse models.Ptcourse
			pt := new(models.Ptcourse)
			ci, err := strconv.ParseInt(cidlist[i], 10, 64)
			if err == nil {
				ptcourse.Cid = ci
			}
			pmi, err := strconv.ParseInt(pmidlist[j], 10, 64)
			if err == nil {
				ptcourse.Pmid = pmi
			}
			//先查询再建立
			num, err := o.QueryTable(pt).Filter("Cid", ci).Filter("Pmid", pmi).Count()
			if err != nil {
				this.ajaxMsg("", MSG_ERR)
				fmt.Println("查询失败")
			}
			fmt.Println("query num:", num)
			if num == 0 {
				id, err := o.Insert(&ptcourse)
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
func (this *PTCourseController) PTCourseEdit() {
	fmt.Println("查看专业课程")
	pmid := this.Input().Get("pmid")
	fmt.Println(pmid)

	//取其中的pmid
	reg := regexp.MustCompile(`[[:digit:]]+`)
	pmidlist := reg.FindAllString(pmid, -1)

	o := orm.NewOrm()
	var maps []orm.Params
	ptc := new(models.Ptcourse)
	tc := new(models.TheoryCourse)
	var tc_info models.TheoryCourse

	num, err := o.QueryTable(ptc).Filter("Pmid", pmidlist[0]).Values(&maps)
	if err == nil {
		fmt.Printf("Result Nums: %d\n", num)
		//		slice := make([]string, num)
		var slice []string
		for _, m := range maps {
			fmt.Println("map_cid:", m["Cid"])
			err := o.QueryTable(tc).Filter("Cid", m["Cid"]).One(&tc_info)
			if err == nil {
				fmt.Println("tc_info:", tc_info.Cname)
				slice = append(slice, tc_info.Cname)
			}

		}
		fmt.Println("slice:", slice)
		fmt.Println("len slice:", len(slice))
		this.Data["m"] = maps
		this.Data["pmid"] = pmid
		this.Data["s"] = slice
	}
	this.TplName = "editPtcourse.tpl"
	return
}

//删除课程
func (this *PTCourseController) PTCourseDelete() {
	fmt.Println("删除")
	cname := this.Input().Get("cname")
	cnamelist := strings.Split(cname, ",")
	cname_count := len(cnamelist) - 1
	fmt.Println("cname del:", cname, cnamelist)

	pmid := this.Input().Get("pmid")
	fmt.Println("pmid:", pmid)

	//取其中的pmid
	reg := regexp.MustCompile(`[[:digit:]]+`)
	pmidlist := reg.FindAllString(pmid, -1)

	o := orm.NewOrm()
	tc := new(models.TheoryCourse)
	ptc := new(models.Ptcourse)
	var tc_info models.TheoryCourse

	for i := 0; i < cname_count; i++ {
		err := o.QueryTable(tc).Filter("Cname", cnamelist[i]).One(&tc_info)
		if err == nil {
			num, err := o.QueryTable(ptc).Filter("Cid", tc_info.Cid).Filter("Pmid", pmidlist[0]).Delete()
			if err == nil {
				fmt.Printf("删除成功")
				fmt.Printf("Result delCid Nums: %d\n", num)
			}
		}
	}

	this.ajaxMsg("", MSG_OK)
	return

}
