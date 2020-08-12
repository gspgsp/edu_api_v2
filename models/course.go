package models

type Course struct {
	ID              int64   `db:"id"json:"id"`
	Type            string  `db:"type"json:"type"`
	Title           string  `db:"title"json:"title"`
	Subtitle        *string `db:"subtitle"json:"subtitle"`
	DifficultyLevel int64   `db:"difficulty_level"json:"difficulty_level"`
	CoverPicture    *string `db:"cover_picture"json:"cover_picture"`
}
