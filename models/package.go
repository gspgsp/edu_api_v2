package models

type Package struct {
	ID       int64  `db:"id"json:"id"`
	Title    string `db:"title"json:"title"`
	Subtitle string `db:"subtitle"json:"subtitle"`
	LearnNum int    `db:"learn_num"json:"learn_num"`
	Length   int    `db:"length"json:"length"`
}
