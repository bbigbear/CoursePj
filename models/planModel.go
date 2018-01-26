package models

type Plan struct {
	Id      int64
	Plid    int64
	Year    int64  `orm:"index"`
	Faculty string `orm:"index"`
	Plname  string
}
