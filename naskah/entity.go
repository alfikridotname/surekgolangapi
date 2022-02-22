package naskah

import "time"

type MasterNaskah struct {
	ID        int       `gorm:"primary_key" json:"id"`
	Nama      string    `json:"nama"`
	Deskripsi string    `json:"deskripsi"`
	IsActive  bool      `json:"is_active"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type Tabler interface {
	TableName() string
}

func (MasterNaskah) TableName() string {
	return "master_naskah"
}
