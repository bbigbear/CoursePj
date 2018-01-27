package controllers

import (
	"CoursePj/models"
	"fmt"
	"strconv"

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
	query := o.QueryTable(pm).Filter("Status", "可用")
	var slice []string
	num, err := query.Values(&maps)
	if err == nil {
		fmt.Printf("Result Nums: %d\n", num)

		for _, m := range maps {
			//获取院系
			faculty := fmt.Sprint(m["Faculty"])
			slice = append(slice, faculty)

		}
		//		this.Data["m"] = Pmslice
		//		this.Data["num"] = num
	}

	this.Data["m"] = this.RemoveRepBySlice(slice)

	this.TplName = "pgcredit.tpl"
}

//检索
func (this *PGCreditController) PgcSearch() {
	o := orm.NewOrm()
	var maps []orm.Params
	pm := new(models.Pm)
	pgc := new(models.Pgcredits)
	year := this.Input().Get("year")
	this.Data["y"] = year
	faculty := this.Input().Get("faculty")
	this.Data["f"] = faculty
	query := o.QueryTable(pm).Filter("Status", "可用").Filter("Year", year).Filter("Faculty", faculty)
	query_pgc := o.QueryTable(pgc).Filter("Year", year).Filter("Faculty", faculty)

	//院系下拉框
	query1 := o.QueryTable(pm).Filter("Status", "可用")
	var slice []string
	num, err1 := query1.Values(&maps)
	if err1 == nil {
		fmt.Printf("Result Nums: %d\n", num)
		for _, m := range maps {
			//获取院系
			faculty := fmt.Sprint(m["Faculty"])
			slice = append(slice, faculty)

		}
		this.Data["m"] = this.RemoveRepBySlice(slice)
	}

	//专业列表
	var Pmslice []string
	//mac
	Pmmap := make(map[int64]string)
	//未设置专业列表
	var Pmslice_NotSet []string
	//已设置专业列表
	var Pmslice_Set []string
	//合并列表
	//	var slice_merge []string
	//var slice []string
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
	//	slice_merge = append(Pmslice, Pmslice_Set...)
	//	fmt.Println("slice_merge:", slice_merge)

	Pmslice_NotSet = RemoveRepByLoop(Pmslice, Pmslice_Set)

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
	pmname := this.Input().Get("pmname")
	year := this.Input().Get("year")
	faculty := this.Input().Get("faculty")
	this.Data["pmname"] = pmname
	this.Data["y"] = year
	this.Data["f"] = faculty
	fmt.Println("pmname:", pmname)

	this.TplName = "addPgcredit.tpl"
}

//保存专业学分
func (this *PGCreditController) PgcSave() {
	//
	fmt.Println("保存pgc")
	pmname := this.Input().Get("pmname")
	year, err := strconv.ParseInt(this.Input().Get("year"), 10, 64)
	if err != nil {
		fmt.Println("year error!")
	}
	faculty := this.Input().Get("faculty")
	ggbx, err := strconv.ParseFloat(this.Input().Get("ggbx"), 64)
	if err != nil {
		fmt.Println("ggbx error!")
	}
	ggrx, err := strconv.ParseFloat(this.Input().Get("ggrx"), 64)
	if err != nil {
		fmt.Println("ggrx error!")
	}
	zybx, err := strconv.ParseFloat(this.Input().Get("zybx"), 64)
	if err != nil {
		fmt.Println("zybx error!")
	}
	zyxx, err := strconv.ParseFloat(this.Input().Get("zyxx"), 64)
	if err != nil {
		fmt.Println("zyxx error!")
	}
	zyrx, err := strconv.ParseFloat(this.Input().Get("zyrx"), 64)
	if err != nil {
		fmt.Println("zyrx error!")
	}
	sjxf, err := strconv.ParseFloat(this.Input().Get("sjxf"), 64)
	if err != nil {
		fmt.Println("sjxf error!")
	}
	zxf, err := strconv.ParseFloat(this.Input().Get("zxf"), 64)
	if err != nil {
		fmt.Println("zxf error!")
	}
	fmt.Println("pgc_info:", pmname, ggbx, ggrx, zybx, zyrx, zyxx, sjxf, zxf)

	o := orm.NewOrm()
	pm := new(models.Pm)
	pgc := new(models.Pgcredits)
	var pm_info models.Pm
	err1 := o.QueryTable(pm).Filter("Pmname", pmname).Filter("Year", year).Filter("Faculty", faculty).One(&pm_info)
	if err1 == nil {
		pmid := pm_info.Pmid
		//插入专业学分
		fmt.Println("pmid:", pmid)
		pgc.Year = year
		pgc.Pgcid = pmid
		pgc.Faculty = faculty
		pgc.Open_require_credit = ggbx
		pgc.Open_option_credit = ggrx
		pgc.Professional_require_credit = zybx
		pgc.Professional_limit_credit = zyxx
		pgc.Professional_option_credit = zyrx
		pgc.Practice_credit = sjxf
		pgc.Total_credit = zxf

		num, err := o.Insert(pgc)
		if err != nil {
			fmt.Println("插入失败")
			this.ajaxMsg("", MSG_ERR)
		}
		fmt.Println("成功插入num:", num)
		this.ajaxMsg("", MSG_OK)
		return
	}
}

//跟新
func (this *PGCreditController) PgcUpdate() {
	fmt.Println("更新")
	id, err := strconv.ParseInt(this.Input().Get("id"), 10, 64)
	pmname := this.Input().Get("pmname")
	year, err := strconv.ParseInt(this.Input().Get("year"), 10, 64)
	if err != nil {
		fmt.Println("year error!")
	}
	faculty := this.Input().Get("faculty")
	if err != nil {
		fmt.Println("id error!")
	}
	ggbx, err := strconv.ParseFloat(this.Input().Get("ggbx"), 64)
	if err != nil {
		fmt.Println("ggbx error!")
	}
	ggrx, err := strconv.ParseFloat(this.Input().Get("ggrx"), 64)
	if err != nil {
		fmt.Println("ggrx error!")
	}
	zybx, err := strconv.ParseFloat(this.Input().Get("zybx"), 64)
	if err != nil {
		fmt.Println("zybx error!")
	}
	zyxx, err := strconv.ParseFloat(this.Input().Get("zyxx"), 64)
	if err != nil {
		fmt.Println("zyxx error!")
	}
	zyrx, err := strconv.ParseFloat(this.Input().Get("zyrx"), 64)
	if err != nil {
		fmt.Println("zyrx error!")
	}
	sjxf, err := strconv.ParseFloat(this.Input().Get("sjxf"), 64)
	if err != nil {
		fmt.Println("sjxf error!")
	}
	zxf, err := strconv.ParseFloat(this.Input().Get("zxf"), 64)
	if err != nil {
		fmt.Println("zxf error!")
	}
	fmt.Println("pgc_info:", year, faculty, pmname, id, ggbx, ggrx, zybx, zyrx, zyxx, sjxf, zxf)

	o := orm.NewOrm()
	pm := new(models.Pm)
	pgc := new(models.Pgcredits)
	var pm_info models.Pm
	err1 := o.QueryTable(pm).Filter("Pmname", pmname).Filter("Year", year).Filter("Faculty", faculty).One(&pm_info)
	if err1 == nil {
		//更新专业学分
		pmid := pm_info.Pmid
		fmt.Println("pmid:", pmid)
		pgc.Pgcid = pmid
		pgc.Id = id
		pgc.Faculty = faculty
		pgc.Year = year
		pgc.Open_require_credit = ggbx
		pgc.Open_option_credit = ggrx
		pgc.Professional_require_credit = zybx
		pgc.Professional_limit_credit = zyxx
		pgc.Professional_option_credit = zyrx
		pgc.Practice_credit = sjxf
		pgc.Total_credit = zxf

		if _, err := o.Update(pgc); err != nil {
			fmt.Println("插入失败")
			this.ajaxMsg("", MSG_ERR)
		}
	}

	this.ajaxMsg("", MSG_OK)
	return

}

//删除专业学分
func (this *PGCreditController) PgcDel() {
	fmt.Println("删除")
	id := this.Input().Get("id")
	o := orm.NewOrm()
	//pm := new(models.Pm)
	pgc := new(models.Pgcredits)
	//var pm_info models.Pm
	//err := o.QueryTable(pm).Filter("id", id).One(&pm_info)
	//if err == nil {
	num, err := o.QueryTable(pgc).Filter("id", id).Delete()
	if err == nil {
		fmt.Println("num:", num)
	}
	//}
	this.ajaxMsg("", MSG_OK)
	return
}

//编辑专业学分
func (this *PGCreditController) PgcEdit() {
	fmt.Println("编辑专业学分")
	pmname := this.Input().Get("pmname")
	this.Data["pmname"] = pmname
	fmt.Println("pmname:", pmname)
	year, err := strconv.ParseInt(this.Input().Get("year"), 10, 64)
	if err != nil {
		fmt.Println("year error!")
	}
	faculty := this.Input().Get("faculty")
	this.Data["y"] = year
	this.Data["f"] = faculty
	o := orm.NewOrm()
	var maps []orm.Params
	pm := new(models.Pm)
	pgc := new(models.Pgcredits)
	var pm_info models.Pm
	err1 := o.QueryTable(pm).Filter("Pmname", pmname).Filter("Year", year).Filter("Faculty", faculty).One(&pm_info)
	if err1 == nil {
		num, err := o.QueryTable(pgc).Filter("Pgcid", pm_info.Pmid).Filter("Year", year).Filter("Faculty", faculty).Values(&maps)
		if err == nil {
			fmt.Println("num:", num)
			this.Data["pgc_info"] = maps
		}
	}
	this.TplName = "editPgcredit.tpl"
}

// 通过两重循环挑选不同元素
func RemoveRepByLoop(slc []string, slc1 []string) []string {
	result := []string{} // 存放结果
	if len(slc1) == 0 {
		result = append(result, slc...)
	}
	for i := range slc {
		for j := range slc1 {
			if slc[i] == slc1[j] {
				break
			}
			//存储slice中不同的值
			if j == len(slc1)-1 {
				result = append(result, slc[i])
			}
		}
	}
	return result
}
