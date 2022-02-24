package surat

import (
	"strconv"
	"strings"

	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

type Repository interface {
	Save(surat MasterSurat, tembusan string, penerimaID string) (bool, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) Save(surat MasterSurat, tembusan string, penerimaID string) (bool, error) {
	tx := r.db.Begin()
	surat.ID = uuid.NewV4()
	if err := tx.Create(&surat).Error; err != nil {
		tx.Rollback()
		return false, err
	}

	if penerimaID != "" {
		penerimaIDSlice := strings.Split(penerimaID, ",")

		for _, penerima := range penerimaIDSlice {
			masterPenerima := MasterPenerima{}
			masterPenerima.MasterSuratID = surat.ID.String()
			masterPenerima.UnitKerjaID, _ = strconv.Atoi(penerima)

			result := tx.Table("master_penerima").Create(&masterPenerima)
			if result.Error != nil {
				tx.Rollback()
				return false, result.Error
			}
		}
	}

	if tembusan != "" {
		tembusanSlice := strings.Split(tembusan, ",")

		for _, jabatanID := range tembusanSlice {
			masterTembusan := MasterTembusan{}
			masterTembusan.MasterSuratID = surat.ID.String()
			masterTembusan.JabatanID, _ = strconv.Atoi(jabatanID)

			result := tx.Table("master_tembusan").Create(&masterTembusan)
			if result.Error != nil {
				tx.Rollback()
				return false, result.Error
			}
		}
	}

	tx.Commit()

	return true, nil
}
