package models

type PPLink struct {
	Pid  int64 `orm:"PK"`
	Pmid int64
}
