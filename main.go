package main

import (
	"CoursePj/models"
	_ "CoursePj/routers"
	"fmt"

	_ "github.com/GO-SQL-Driver/MySQL"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

func init() {
	DBConnection()
	RegisterModel()
}

func main() {
	orm.Debug = true
	orm.RunSyncdb("default", false, true)
	beego.Run()
}

func DBConnection() {
	fmt.Println("初始化数据库")
	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", "root:qwe!23@/course_arrangement?charset=utf8")
}
func RegisterModel() {
	fmt.Println("注册数据库模型")
	orm.RegisterModel(new(models.TheoryCourse))
	orm.RegisterModel(new(models.Practice))
	orm.RegisterModel(new(models.Pm))
	orm.RegisterModel(new(models.Ptcourse))
}
