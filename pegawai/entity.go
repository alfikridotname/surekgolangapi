package pegawai

import "time"

type MasterPegawai struct {
	ID                int       `gorm:"primary_key"`
	UnitKerjaID       int       `gorm:"column:unit_kerja_id"`
	SubOpdID          int       `gorm:"column:sub_opd_id"`
	KategoriJabatanID int       `gorm:"column:kat_jabatan_id"`
	JabatanID         int       `gorm:"column:jabatan_id"`
	JabatanNM         string    `gorm:"column:jabatan_nm"`
	NIP               string    `gorm:"column:nip"`
	Nama              string    `gorm:"column:nama"`
	EselonID          int       `gorm:"column:eselon_id"`
	Kategori          string    `gorm:"column:kategori"`
	CreatedAt         time.Time `gorm:"column:created_at"`
	UpdatedAt         time.Time `gorm:"column:updated_at"`
}

type UserID struct {
	ID int `json:"id"`
}

type Tabler interface {
	TableName() string
}

func (MasterPegawai) TableName() string {
	return "master_pegawai"
}
