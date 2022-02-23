package user

import "time"

type User struct {
	ID              int       `gorm:"primary_key" json:"id"`
	Name            string    `json:"name"`
	Username        string    `json:"username"`
	Nip             string    `json:"nip"`
	Email           string    `json:"email"`
	EmailVerifiedAt time.Time `json:"email_verified_at"`
	Password        string    `json:"password"`
	RememberToken   string    `json:"remember_token"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
	UnitKerjaID     int       `json:"unit_kerja_id"`
}
