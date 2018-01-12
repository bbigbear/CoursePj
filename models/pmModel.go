package models

type Pm struct {
	Pmid        int64  `orm:"PK"`
	Pmname      string `orm:"index"`
	Pmname_en   string
	Faculty     string `orm:"index"`
	Status      string `orm:"index"`
	Train_level string
	Isminor     string
	Year        int64 `orm:"index"`
}
