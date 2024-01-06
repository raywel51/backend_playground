package model

type CarTypeDao struct {
	ID      int    `json:"id"`
	CarType int    `json:"car_type"`
	ValueTH string `json:"value_th"`
	ValueEN string `json:"value_en"`
	ValueCN string `json:"value_cn"`
}
