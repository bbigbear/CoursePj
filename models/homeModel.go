package models

type TheoryCourse struct {
	Cid         int64 `orm:"PK"`
	Cunit       string
	Cname       string `orm:"index"`
	Ccg1        string `orm:"index"`
	Ccg2        string `orm:"index"`
	Cname_en    string
	Status      string `orm:"index"`
	Credit      float64
	Tteach      int
	Texperiment int
	Tcomputer   int
	Tother      int
	Ttotal      int
	Syllabus    string `orm:"size(5000)"`
	Year        int64  `orm:"index"`
}
