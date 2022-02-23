package surat

import (
	"time"

	uuid "github.com/satori/go.uuid"
)

type MasterSurat struct {
	ID                 uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4()"`
	KategoriPenerimaID int       `json:"kategori_penerima_id"`
	NaskahID           int       `json:"naskah_id"`
	SignerID           int       `json:"signer_id"`
	AtasNamaID         int       `json:"atas_nama_id"`
	KeamananID         int       `json:"keamanan_id"`
	KecepatanID        int       `json:"kecepatan_id"`
	Tgl                time.Time `json:"tgl"`
	Perihal            string    `json:"perihal"`
	Isi                string    `json:"isi"`
	CreatedAt          time.Time `json:"created_at"`
	UpdatedAt          time.Time `json:"updated_at"`
	CreatedBy          int       `json:"created_by"`
	UpdatedBy          int       `json:"updated_by"`
}

type Tabler interface {
	TableName() string
}

// TableName overrides the table name used by User to `profiles`
func (MasterSurat) TableName() string {
	return "master_surat"
}
