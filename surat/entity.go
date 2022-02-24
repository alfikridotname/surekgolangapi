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
	Tgl                string    `json:"tgl"`
	Perihal            string    `json:"perihal"`
	Isi                string    `json:"isi"`
	CreatedAt          time.Time `json:"created_at"`
	UpdatedAt          time.Time `json:"updated_at"`
	CreatedBy          int       `json:"created_by"`
	UpdatedBy          int       `json:"updated_by"`
}

type MasterTembusan struct {
	ID            int    `gorm:"primary_key" json:"id"`
	MasterSuratID string `json:"master_surat_id"`
	JabatanID     int    `json:"jabatan_id"`
}

type MasterPenerima struct {
	ID            int    `json:"id"`
	MasterSuratID string `json:"master_surat_id"`
	UnitKerjaID   int    `json:"unit_kerja_id"`
}

type Tabler interface {
	TableName() string
}

func (MasterSurat) TableName() string {
	return "master_surat"
}

func (MasterTembusan) TableName() string {
	return "master_tembusan"
}

func (MasterPenerima) TableName() string {
	return "master_penerima"
}
