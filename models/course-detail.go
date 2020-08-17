package models

type CourseDetail struct {
	Boutique
	Uuid            string  `db:"uuid"json:"uuid"`
	Length          int64   `db:"length"json:"length"`
	Rating          float64 `db:"rating"json:"rating"`
	PracticalRating float64 `db:"practical_rating"json:"practical_rating"`
	PopularRating   float64 `db:"popular_rating"json:"popular_rating"`
	LogicRating     float64 `db:"logic_rating"json:"logic_rating"`
	Goals           string  `db:"goals"json:"goals"`
	Audiences       string  `db:"audiences"json:"audiences"`
	Summary         string  `db:"summary"json:"summary"`
}
