package notifikasi

import "gorm.io/datatypes"

type NotifikasiInput struct {
	ID           int            `json:"id"`
	Data         datatypes.JSON `json:"data"`
	UserTujuanID int            `json:"user_tujuan_id"`
	CreatedBY    int            `json:"created_by"`
}
