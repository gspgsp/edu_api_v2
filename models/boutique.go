package models

type Boutique struct {
	Course
	LearnCount    int64    `db:"learn_count"json:"learn_count"`
	Price         float64  `db:"price"json:"price"`
	Discount      *float64 `db:"discount"json:"discount"`
	DiscountEndAt *string  `db:"discount_end_at"json:"discount_end_at"`
	VipLevel      int64    `db:"vip_level"json:"vip_level"`
	VipPrice      float64  `db:"vip_price"json:"vip_price"`
}
