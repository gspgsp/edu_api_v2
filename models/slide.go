package models

type Slide struct {
	ID             int64  `db:"id"json:"id"`
	Port           int    `db:"port" json:"port"`
	Title          string `db:"title" json:"title"`
	Url            string `db:"url" json:"url"`
	Carousel       string `db:"carousel" json:"carousel"`
	Sort           string `db:"sort" json:"sort"`
	Status         int    `db:"status" json:"status"`
	Description    string `db:"description" json:"description"`
	CreatedAt      string `db:"created_at" json:"created_at"`
	UpdatedAt      string `db:"updated_at" json:"updated_at"`
	CreatedAdminId int    `db:"created_admin_id" json:"-"`
	UpdatedAdminId int    `db:"updated_admin_id" json:"-"`
}
