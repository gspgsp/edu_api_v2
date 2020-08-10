package models

type Free struct {
	Course
	LearnCount int64 `db:"learn_count"json:"learn_count"`
}
