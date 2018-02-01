package controllers

import (
	"fmt"

	"CoursePj/models"

	"strconv"

	_ "github.com/GO-SQL-Driver/MySQL"
	"github.com/astaxie/beego/orm"
)

type CopyPlanController struct {
	BaseController
}

func (this *CopyPlanController) GetProfessionPlan() {

	pm := new(models.Pm)
	o := orm.NewOrm()
	var maps []orm.Params
	var slice []string
	query := o.QueryTable(pm).Filter("Status", "可用")

	num, err := query.Values(&maps)
	if err == nil {
		fmt.Printf("plan pm zhuanye num", num)
		//fmt.Printf("plan pm zhuanye map", maps)

		for _, m := range maps {
			faculty := fmt.Sprint(m["Faculty"])
			slice = append(slice, faculty)
			fmt.Println("maps_mp_info:", m["Faculty"])
		}

		this.Data["m"] = this.RemoveRepBySlice(slice)
	}

	this.TplName = "copyProfessionPlan.tpl"

}

func (this *CopyPlanController) PPSearch() {
	//获取学院列表
	o := orm.NewOrm()
	pm := new(models.Pm)
	var slice_pm []string
	var maps_pm []orm.Params
	num, err2 := o.QueryTable(pm).Filter("Status", "可用").Values(&maps_pm)
	if err2 == nil {
		fmt.Printf("plan pm zhuanye num", num)
		for _, m := range maps_pm {
			faculty := fmt.Sprint(m["Faculty"])
			slice_pm = append(slice_pm, faculty)
			fmt.Println("maps_mp_info:", m["Faculty"])
		}
		this.Data["m"] = this.RemoveRepBySlice(slice_pm)
	}

	fmt.Println("点击检索")
	filters := make([]interface{}, 0)
	faculty := this.Input().Get("faculty")
	if faculty != "" {
		filters = append(filters, "Faculty", faculty)
	}
	this.Data["f"] = faculty
	fmt.Println("Faculty:", faculty)
	year := this.Input().Get("year")
	if year != "" {
		filters = append(filters, "Year", year)
	}
	fmt.Println("Year:", year)
	this.Data["y"] = year
	var maps []orm.Params
	query := o.QueryTable(pm)
	pt := new(models.Ptcourse)
	pp := new(models.Pplink)

	var slice []string

	if len(filters) > 0 {
		l := len(filters)
		for k := 0; k < l; k += 2 {
			query = query.Filter(filters[k].(string), filters[k+1])
		}
	}

	num, err := query.Values(&maps)
	if err == nil {
		fmt.Println("get pm num:", num)
		for _, m := range maps {
			pmid := fmt.Sprint(m["Pmid"])
			pmname := fmt.Sprint(m["Pmname"])
			//查询
			exist := o.QueryTable(pt).Filter("Pmid", pmid).Exist()
			if exist {
				exist = o.QueryTable(pp).Filter("Pmid", pmid).Exist()
				if exist {
					slice = append(slice, pmname)
				}
			}
		}
		fmt.Println("slice:", slice)
		this.Data["s"] = slice
		this.Data["len"] = len(slice)
	}
	//已制定专业列表
	plan := new(models.Plan)
	var maps_plan []orm.Params
	var pm_info models.Pm
	var slice_plan []string
	num, err1 := o.QueryTable(plan).Filter("Faculty", faculty).Filter("Year", year).Values(&maps_plan)
	if err1 == nil {
		fmt.Println("plan num:", num)
		for _, m := range maps_plan {
			//plid := fmt.Sprint(maps_plan["Plid"])
			err := o.QueryTable(pm).Filter("Pmid", m["Plid"]).One(&pm_info)
			if err == nil {
				fmt.Println("om_info,pmname:", pm_info.Pmname)
				slice_plan = append(slice_plan, pm_info.Pmname)
			}
		}
		fmt.Println("slice_plan", slice_plan)
		this.Data["slice_plan"] = slice_plan
		this.Data["slice_plan_len"] = len(slice_plan)

	}

	this.TplName = "copyProfessionPlan.tpl"
	//this.ajaxMsg("", MSG_OK)
	return
}
func (this *CopyPlanController) PPCopy() {
	fmt.Println("点击复制")
	pmname := this.Input().Get("pmname")
	fmt.Println("copy pmname:", pmname)
	faculty := this.Input().Get("faculty")
	fmt.Println("copy faculty:", faculty)
	year, err := strconv.ParseInt(this.Input().Get("year"), 10, 64)
	if err != nil {
		fmt.Println("year error!")
	}
	fmt.Println("copy year:", year)
	if pmname == "" || faculty == "" || year == 0 {
		this.ajaxMsg("year,faculty,pmname参数不能为空", MSG_ERR_Param)
	}
	plan := new(models.Plan)
	pm := new(models.Pm)
	var pm_info models.Pm
	o := orm.NewOrm()

	//先查询
	exist := o.QueryTable(plan).Filter("Plname", pmname).Filter("Year", year).Filter("Faculty", faculty).Exist()
	if exist {
		fmt.Println("已存在")
		this.ajaxMsg("复制失败，已存在专业", MSG_ERR_Resources)
	} else {
		err1 := o.QueryTable(pm).Filter("Pmname", pmname).Filter("Year", year).Filter("Faculty", faculty).One(&pm_info)
		if err1 == nil {
			plan.Plid = pm_info.Pmid
			plan.Plname = pm_info.Pmname
			plan.Faculty = faculty
			plan.Year = year
			num, err := o.Insert(plan)
			if err != nil {
				fmt.Println("插入plan失败")
				this.ajaxMsg("复制失败", MSG_ERR_Resources)
			}
			if num == 0 {
				this.ajaxMsg("复制失败", MSG_ERR_Resources)
			}
			fmt.Println("成功插入num:", num)

		}

	}
	this.ajaxMsg("复制成功", MSG_OK)
	return

}

func (this *CopyPlanController) PPRemove() {
	fmt.Println("点击移除")
	pmname := this.Input().Get("pmname")
	fmt.Println("delete pmname:", pmname)
	faculty := this.Input().Get("faculty")
	fmt.Println("delete faculty:", faculty)
	year, err := strconv.ParseInt(this.Input().Get("year"), 10, 64)
	if err != nil {
		fmt.Println("year error!")
	}
	if pmname == "" || faculty == "" || year == 0 {
		this.ajaxMsg("year,faculty,pmname参数不能为空", MSG_ERR_Param)
	}
	plan := new(models.Plan)
	pm := new(models.Pm)
	var pm_info models.Pm
	o := orm.NewOrm()
	err1 := o.QueryTable(pm).Filter("Pmname", pmname).Filter("Year", year).Filter("Faculty", faculty).One(&pm_info)
	if err1 == nil {
		//plan.Plid = pm_info.Pmid
		num, err := o.QueryTable(plan).Filter("Plid", pm_info.Pmid).Filter("Year", year).Filter("Faculty", faculty).Delete()
		if err == nil {
			fmt.Printf("删除成功")
			fmt.Printf("Result Nums: %d\n", num)
			if num == 0 {
				this.ajaxMsg("删除失败", MSG_ERR_Resources)
			}
		} else {
			this.ajaxMsg("删除失败", MSG_ERR_Resources)
		}

	}
	this.ajaxMsg("删除成功", MSG_OK)
	return

}

func (this *CopyPlanController) GetYearPlan() {

	pm := new(models.Pm)
	o := orm.NewOrm()
	var maps []orm.Params
	var slice []string
	query := o.QueryTable(pm).Filter("Status", "可用")

	num, err := query.Values(&maps)
	if err == nil {
		fmt.Printf("plan pm zhuanye num", num)
		//fmt.Printf("plan pm zhuanye map", maps)

		for _, m := range maps {
			faculty := fmt.Sprint(m["Faculty"])
			slice = append(slice, faculty)
			fmt.Println("maps_mp_info:", m["Faculty"])
		}

		this.Data["m"] = this.RemoveRepBySlice(slice)
	}

	this.TplName = "copyYearPlan.tpl"
}

func (this *CopyPlanController) GYSearch() {
	//获取学院列表
	o := orm.NewOrm()
	pm := new(models.Pm)
	var slice_pm []string
	var maps_pm []orm.Params
	num, err2 := o.QueryTable(pm).Filter("Status", "可用").Values(&maps_pm)
	if err2 == nil {
		fmt.Printf("plan pm zhuanye num", num)
		for _, m := range maps_pm {
			faculty := fmt.Sprint(m["Faculty"])
			slice_pm = append(slice_pm, faculty)
			fmt.Println("maps_mp_info:", m["Faculty"])
		}
		this.Data["m"] = this.RemoveRepBySlice(slice_pm)
	}

	fmt.Println("点击GYSearch检索")
	faculty := this.Input().Get("faculty")
	this.Data["f"] = faculty
	fmt.Println("Faculty:", faculty)
	year := this.Input().Get("year")
	this.Data["y"] = year
	fmt.Println("Year:", year)

	//已制定专业列表
	plan := new(models.Plan)
	var maps []orm.Params
	query := o.QueryTable(plan)
	num, err1 := query.Filter("Faculty", faculty).Filter("Year", year).Values(&maps)
	if err1 == nil {
		fmt.Println("plan num:", num)
		this.Data["maps"] = maps
		this.Data["l"] = num
	}

	//右边列表
	year_right := this.Input().Get("year_right")
	if year_right != "" {

		var maps_right []orm.Params
		num, err := query.Filter("Faculty", faculty).Filter("Year", year_right).Values(&maps_right)
		if err == nil {
			this.Data["maps_right"] = maps_right
			this.Data["l1"] = num
			this.Data["year_right"] = year_right
		}
	}

	this.TplName = "copyYearPlan.tpl"
	//this.ajaxMsg("", MSG_OK)
	return
}
func (this *CopyPlanController) GYCopy() {
	fmt.Println("GY点击复制")
	plname := this.Input().Get("plname")
	fmt.Println("GYcopy plname:", plname)
	faculty := this.Input().Get("faculty")
	fmt.Println("GYcopy faculty:", faculty)
	right_year, err := strconv.ParseInt(this.Input().Get("right_year"), 10, 64)
	if err != nil {
		fmt.Println("year error!")
	}
	year, err := strconv.ParseInt(this.Input().Get("year"), 10, 64)
	if err != nil {
		fmt.Println("year error!")
	}
	fmt.Println("copy year:", year, right_year)
	if plname == "" || faculty == "" || right_year == 0 || year == 0 {
		this.ajaxMsg("year,faculty,plname,right_year参数不能为空", MSG_ERR_Param)
	}
	plan := new(models.Plan)
	pm := new(models.Pm)
	tc := new(models.TheoryCourse)
	p := new(models.Practice)
	pt := new(models.Ptcourse)
	pp := new(models.Pplink)
	var pm_info models.Pm
	var tc_info models.TheoryCourse
	var p_info models.Practice
	var pt_info models.Ptcourse
	var pp_info models.Pplink
	o := orm.NewOrm()

	//先查询
	exist := o.QueryTable(plan).Filter("Plname", plname).Filter("Faculty", faculty).Filter("Year", right_year).Exist()
	if exist {
		fmt.Println("已存在")
		this.ajaxMsg("复制失败，已存在专业", MSG_ERR_Resources)
	} else {
		err1 := o.QueryTable(pm).Filter("Pmname", plname).Filter("Faculty", faculty).Filter("Year", year).One(&pm_info)
		if err1 == nil {
			//插入plan数据库
			plan.Plid = pm_info.Pmid
			plan.Plname = pm_info.Pmname
			plan.Faculty = faculty
			plan.Year = right_year
			num, err := o.Insert(plan)
			if err != nil {
				fmt.Println("插入失败")
				this.ajaxMsg("复制失败", MSG_ERR_Resources)
			}
			if num == 0 {
				this.ajaxMsg("复制失败", MSG_ERR_Resources)
			}
			fmt.Println("成功插入num:", num)
			//插入pm数据库
			id, err := strconv.ParseInt("", 10, 64)
			if err != nil {
				fmt.Println("错误")
			}
			pm_info.Id = id
			pm_info.Year = right_year
			pm_num, err := o.Insert(&pm_info)
			if err != nil {
				fmt.Println("插入pm失败")
				this.ajaxMsg("复制失败", MSG_ERR_Resources)
			}
			fmt.Println("成功插入pm_num:", pm_num)

			//插入tc数据库，先查询cid
			pt_err := o.QueryTable(pt).Filter("Pmid", pm_info.Pmid).One(&pt_info)
			if pt_err == nil {
				tc_err := o.QueryTable(tc).Filter("Cid", pt_info.Cid).Filter("Year", year).One(&tc_info)
				if tc_err == nil {
					tc_info.Id = id
					tc_info.Year = right_year
					tc_num, err := o.Insert(&tc_info)
					if err != nil {
						fmt.Println("插入tc失败")
						this.ajaxMsg("复制失败", MSG_ERR_Resources)
					}
					fmt.Println("成功插入tc_num:", tc_num)
				}
			}
			//插入p数据库
			pp_err := o.QueryTable(pp).Filter("Pmid", pm_info.Pmid).One(&pp_info)
			if pp_err == nil {
				p_err := o.QueryTable(p).Filter("Pid", pp_info.Pid).Filter("Year", year).One(&p_info)
				if p_err == nil {
					p_info.Id = id
					p_info.Year = right_year
					p_num, err := o.Insert(&p_info)
					if err != nil {
						fmt.Println("插入p失败")
						this.ajaxMsg("复制失败", MSG_ERR_Resources)
					}
					fmt.Println("成功插入p_num:", p_num)
				}
			}

		}

	}
	this.ajaxMsg("复制成功", MSG_OK)
	return

}

func (this *CopyPlanController) GYRemove() {
	fmt.Println("GY点击移除")
	pmname := this.Input().Get("pmname")
	fmt.Println("copy pmname:", pmname)
	faculty := this.Input().Get("faculty")
	fmt.Println("delete faculty:", faculty)
	year, err := strconv.ParseInt(this.Input().Get("year"), 10, 64)
	if err != nil {
		fmt.Println("year error!")
	}
	if pmname == "" || faculty == "" || year == 0 {
		this.ajaxMsg("year,faculty,pmname参数不能为空", MSG_ERR_Param)
	}
	plan := new(models.Plan)
	o := orm.NewOrm()
	num, err := o.QueryTable(plan).Filter("Plname", pmname).Filter("Year", year).Filter("Faculty", faculty).Delete()
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

}
