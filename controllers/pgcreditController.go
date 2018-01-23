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

type PGCreditController struct {
	BaseController
}

func (this *PGCreditController) Get() {

	o := orm.NewOrm()
	var maps []orm.Params
	pm := new(models.Pm)
	pgc := new(models.Pgcredits)
	query := o.QueryTable(pm).Filter("Status", "可用")
	query_pgc := o.QueryTable(pgc)
	//专业列表
	var Pmslice []string
	//mac
	Pmmap := make(map[int64]string)
	//未设置专业列表
	var Pmslice_NotSet []string
	//已设置专业列表
	var Pmslice_Set []string
	//合并列表
	var slice_merge []string

	//专业
	num, err := query.Values(&maps)
	if err == nil {
		fmt.Printf("Result Nums: %d\n", num)

		for _, m := range maps {
			pmname := fmt.Sprint(m["Pmname"])
			pmid := fmt.Sprint(m["Pmid"])
			id, err := strconv.ParseInt(pmid, 10, 64)
			if err == nil {
				Pmmap[id] = pmname
			}
			Pmslice = append(Pmslice, pmname)

		}
		//		this.Data["m"] = Pmslice
		//		this.Data["num"] = num
	}
	//专业毕业学分
	pgc_num, err := query_pgc.Values(&maps)
	if err == nil {
		fmt.Printf("Result pgc_num:\n", pgc_num)
		for _, pgc_m := range maps {
			pgcid := fmt.Sprint(pgc_m["Pgcid"])
			id, err := strconv.ParseInt(pgcid, 10, 64)
			if err == nil {
				pgname := Pmmap[id]
				Pmslice_Set = append(Pmslice_Set, pgname)
			}

		}
		fmt.Println("Pmslice_set:", Pmslice_Set)
	}

	//合并
	slice_merge = append(Pmslice, Pmslice_Set...)
	Pmslice_NotSet = RemoveRepByLoop(slice_merge)

	fmt.Println("Pmslice_NotSet:", Pmslice_NotSet)

	this.Data["Pmslice_Set"] = Pmslice_Set
	this.Data["Pmslice_Set_count"] = len(Pmslice_Set)
	this.Data["Pmslice_NotSet"] = Pmslice_NotSet
	this.Data["Pmslice_NotSet_count"] = len(Pmslice_NotSet)

	this.TplName = "pgcredit.tpl"
}

//设置专业学分
func (this *PGCreditController) PgcAdd() {
	fmt.Println("设置专业学分")
	this.TplName = "addPgcredit.tpl"
}

//专业设置课程
func (this *PGCreditController) Setcourse() {
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
func (this *PGCreditController) PgcEdit() {
	//	fmt.Println("查看专业课程")
	//	pmid := this.Input().Get("pmid")
	//	fmt.Println(pmid)

	//	//取其中的pmid
	//	reg := regexp.MustCompile(`[[:digit:]]+`)
	//	pmidlist := reg.FindAllString(pmid, -1)

	//	o := orm.NewOrm()
	//	var maps []orm.Params
	//	ptc := new(models.Ptcourse)
	//	tc := new(models.TheoryCourse)
	//	var tc_info models.TheoryCourse

	//	num, err := o.QueryTable(ptc).Filter("Pmid", pmidlist[0]).Values(&maps)
	//	if err == nil {
	//		fmt.Printf("Result Nums: %d\n", num)
	//		//		slice := make([]string, num)
	//		var slice []string
	//		for _, m := range maps {
	//			fmt.Println("map_cid:", m["Cid"])
	//			err := o.QueryTable(tc).Filter("Cid", m["Cid"]).One(&tc_info)
	//			if err == nil {
	//				fmt.Println("tc_info:", tc_info.Cname)
	//				slice = append(slice, tc_info.Cname)
	//			}

	//		}
	//		fmt.Println("slice:", slice)
	//		fmt.Println("len slice:", len(slice))
	//		this.Data["m"] = maps
	//		this.Data["pmid"] = pmid
	//		this.Data["s"] = slice
	//	}
	this.TplName = "editPgcredit.tpl"
	return
}

//删除课程
func (this *PGCreditController) PTCourseDelete() {
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

// 通过两重循环过滤重复元素
func RemoveRepByLoop(slc []string) []string {
	result := []string{} // 存放结果
	for i := range slc {
		flag := true
		for j := range result {
			if slc[i] == result[j] {
				flag = false // 存在重复元素，标识为false
				break
			}
		}
		if flag { // 标识为false，不添加进结果
			result = append(result, slc[i])
		}
	}
	return result
}
