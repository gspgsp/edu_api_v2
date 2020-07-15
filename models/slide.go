package models

import "database/sql"

type Slide struct {
	ID             int64          `db:"id"json:"id"`
	Port           int            `db:"port" json:"port"`
	Title          string         `db:"title" json:"title"`
	Url            string         `db:"url" json:"url"`
	Carousel       string         `db:"carousel" json:"carousel"`
	Sort           sql.NullString `db:"sort" json:"sort,omitempty"`
	Status         int            `db:"status" json:"status"`
	Description    sql.NullString `db:"description" json:"description,omitempty"`
	CreatedAt      string         `db:"created_at" json:"created_at"`
	UpdatedAt      string         `db:"updated_at" json:"updated_at"`
	CreatedAdminId int            `db:"created_admin_id" json:"created_admin_id"`
	UpdatedAdminId int            `db:"updated_admin_id" json:"-"`
}
