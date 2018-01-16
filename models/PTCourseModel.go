package models

type PTCourse struct {
	Cid  int64 `orm:"PK"`
	Pmid int64
}
