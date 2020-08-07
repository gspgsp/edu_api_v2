package models

type Boutique struct {
	ID              int64    `db:"id"json:"id"`
	Title           string   `db:"title"json:"title"`
	Subtitle        *string  `db:"subtitle"json:"subtitle"`
	DifficultyLevel int64    `db:"difficulty_level"json:"difficulty_level"`
	LearnCount      int64    `db:"learn_count"json:"learn_count"`
	Price           float64  `db:"price"json:"price"`
	Discount        *float64 `db:"discount"json:"discount"`
	DiscountEndAt   *string  `db:"discount_end_at"json:"discount_end_at"`
	VipLevel        int64    `db:"vip_level"json:"vip_level"`
	VipPrice        float64  `db:"vip_price"json:"vip_price"`
	CoverPicture    *string  `db:"cover_picture"json:"cover_picture"`
}
