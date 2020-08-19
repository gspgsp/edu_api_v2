package models

type UserReview struct {
	Anonymous  int     `db:"anonymous"json:"anonymous"`
	Rating     float32 `db:"rating"json:"rating"`
	Review     *string `db:"review"json:"review"`
	ReviewedAt *string `db:"reviewed_at"json:"reviewed_at"`
	Reply      *string `db:"reply"json:"reply"`
	Nickname   *string `db:"nickname"json:"nickname"`
}
