package models

type User struct {
	Id               *int64  `db:"id"json:"id"`
	Avatar           *string `db:"avatar"json:"avatar"`
	No               string  `db:"no"json:"no,omitempty"`
	Nickname         *string `db:"nickname"json:"nickname"`
	Title            *int64  `db:"title"json:"title,omitempty"`
	Mobile           string  `db:"mobile"json:"mobile,omitempty"`
	Email            *string `db:"email"json:"email,omitempty"`
	Password         string  `db:"password"json:"password,omitempty"`
	MobileVerified   int     `db:"mobile_verified"json:"mobile_verified,omitempty"`
	MobileVerifiedAt string  `db:"mobile_verified_at"json:"mobile_verified_at,omitempty"`
	EmailVerified    int     `db:"email_verified"json:"email_verified,omitempty"`
	EmailVerifiedAt  string  `db:"email_verified_at"json:"email_verified_at,omitempty"`
	WechatVerified   int     `db:"wechat_verified"json:"wechat_verified,omitempty"`
	WechatVerifiedAt string  `db:"wechat_verified_at"json:"wechat_verified_at,omitempty"`
	Level            string  `db:"level"json:"level,omitempty"`
	StartAt          string  `db:"start_at"json:"start_at,omitempty"`
	EndAt            string  `db:"end_at"json:"end_at,omitempty"`
	Status           int     `db:"status"json:"status,omitempty"`
	IsLecturer       int     `db:"is_lecturer"json:"is_lecturer,omitempty"`
	About            *string `db:"about"json:"about,omitempty"`
	Source           string  `db:"source"json:"source,omitempty"`
	RegisterAt       string  `db:"register_at"json:"register_at,omitempty"`
	RegisterIp       string  `db:"register_ip"json:"register_ip,omitempty"`
	RegisterCity     string  `db:"register_city"json:"register_city,omitempty"`
	LastLoginAt      string  `db:"last_login_at"json:"last_login_at,omitempty"`
	LastLoginIp      string  `db:"last_login_ip"json:"last_login_ip,omitempty"`
	LastLoginCity    string  `db:"last_login_city"json:"last_login_city,omitempty"`
	LoginCount       int     `db:"login_count"json:"login_count,omitempty"`
	CreatedAt        string  `db:"created_at"json:"created_at,omitempty"`
	UpdatedAt        string  `db:"updated_at"json:"updated_at,omitempty"`
}
