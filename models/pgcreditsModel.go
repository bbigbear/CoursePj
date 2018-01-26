package models

type Pgcredits struct {
	Id                          int64
	Pgcid                       int64
	Open_require_credit         float64
	Open_option_credit          float64
	Professional_require_credit float64
	Professional_option_credit  float64
	Professional_limit_credit   float64
	Practice_credit             float64
	Total_credit                float64
	Year                        int64  `orm:"index"`
	Faculty                     string `orm:"index"`
}
