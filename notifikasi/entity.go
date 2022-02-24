package notifikasi

import (
	"time"

	"gorm.io/datatypes"
)

type MasterNotifikasi struct {
	ID           int            `json:"id"`
	Data         datatypes.JSON `json:"data"`
	UserTujuanID int            `json:"user_tujuan_id"`
	CreatedBY    int            `json:"created_by"`
	CreatedAt    time.Time      `json:"created_at"`
	UpdatedAt    time.Time      `json:"updated_at"`
}

type Tabler interface {
	TableName() string
}

func (MasterNotifikasi) TableName() string {
	return "master_notifikasi"
}
