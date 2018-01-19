package models

type Pplink struct {
	Pid  int64 `orm:"PK"`
	Pmid int64
}
