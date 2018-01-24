package controllers

type CopyPlanController struct {
	BaseController
}

func (this *CopyPlanController) GetProfessionPlan() {

	this.TplName = "copyProfessionPlan.tpl"
}

func (this *CopyPlanController) GetYearPlan() {

	this.TplName = "copyYearPlan.tpl"
}
