package models

type Practice struct {
	Id       int64
	Pid      string
	Punit    string
	Pname    string `orm:"index"`
	Pcg1     string `orm:"index"`
	Pname_en string
	Status   string `orm:"index"`
	Credit   float64
	Tclass   int
	Nw       float64
	Syllabus string `orm:"size(5000)"`
	Year     int64  `orm:"index"`
}
