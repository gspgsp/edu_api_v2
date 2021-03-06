package models

type Package struct {
	ID           int64  `db:"id"json:"id"`
	Type         string `db:"type"json:"type"`
	Title        string `db:"title"json:"title"`
	Subtitle     string `db:"subtitle"json:"subtitle"`
	CoverPicture string `db:"cover_picture"json:"cover_picture"`
	LearnNum     int    `db:"learn_num"json:"learn_num"`
	Length       int    `db:"length"json:"length"`
}
