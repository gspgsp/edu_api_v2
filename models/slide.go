package models

type Slide struct {
	ID       int64  `db:"id"json:"id"`
	Port     int    `db:"port" json:"port"`
	Title    string `db:"title" json:"title"`
	Url      string `db:"url" json:"url"`
	Carousel string `db:"carousel" json:"carousel"`
	Sort     string `db:"sort" json:"sort"`
}
