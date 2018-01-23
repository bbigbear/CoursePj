package models

type Pgcredits struct {
	Pgcid                       int64 `orm:"PK"`
	Open_require_credit         float64
	Open_option_credit          float64
	Professional_require_credit float64
	Professional_option_credit  float64
	Professional_limit_credit   float64
	Practice_credit             float64
	Total_credit                float64
}
