package models

type Plan struct {
	Id      int64
	Plid    string
	Year    int64  `orm:"index"`
	Faculty string `orm:"index"`
	Plname  string
}
