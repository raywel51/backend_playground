package entity

type CarTypeDao struct {
	ID      int    `bson:"id"`
	CarType int    `bson:"car_type"`
	ValueTH string `bson:"value_th"`
	ValueEN string `bson:"value_en"`
	ValueCN string `bson:"value_cn"`
}
