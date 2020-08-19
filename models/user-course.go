package models

type UserCourse struct {
	Id              int     `db:"id"json:"id"`
	Type            string  `db:"type"json:"type"`
	CreatedAt       string  `db:"created_at"json:"created_at"`
	FinishedAt      *string `db:"finished_at"json:"finished_at"`
	Reviewed        int     `db:"reviewed"json:"reviewed"`
	Anonymous       int     `db:"anonymous"json:"anonymous"`
	Rating          float32 `db:"rating"json:"rating"`
	PracticalRating int     `db:"practical_rating"json:"practical_rating"`
	PopularRating   int     `db:"popular_rating"json:"popular_rating"`
	LogicRating     int     `db:"logic_rating"json:"logic_rating"`
	Status          int     `db:"status"json:"status"`
	Review          *string `db:"review"json:"review"`
	ReviewedAt      *string `db:"reviewed_at"json:"reviewed_at"`
	Reply           *string `db:"reply"json:"reply"`
	ReplyAt         *string `db:"reply_at"json:"reply_at"`
	Schedule        int     `db:"schedule"json:"schedule"`
	CourseId        int     `db:"course_id"json:"course_id"`
	LessonId        int     `db:"lesson_id"json:"lesson_id"`
	UserId          int     `db:"user_id"json:"user_id"`
}
