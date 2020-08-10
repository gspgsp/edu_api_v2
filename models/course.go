package models

type Course struct {
	ID              int64   `db:"id"json:"id"`
	Title           string  `db:"title"json:"title"`
	Subtitle        *string `db:"subtitle"json:"subtitle"`
	DifficultyLevel int64   `db:"difficulty_level"json:"difficulty_level"`
	CoverPicture    *string `db:"cover_picture"json:"cover_picture"`
}
