package surat

import (
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

type Repository interface {
	Save(surat MasterSurat) (bool, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) Save(surat MasterSurat) (bool, error) {
	tx := r.db.Begin()
	surat.ID = uuid.NewV4()
	if err := tx.Create(&surat).Error; err != nil {
		tx.Rollback()
		return false, err
	}
	// var tembusanStruct MasterTembusan
	// tembusanSlice := strings.Split(tembusanStruct.Tembusan, ",")
	// for _, tembusan := range tembusanSlice {
	// 	tembusanStruct.MasterSuratID = surat.ID
	// 	intVar, _ := strconv.Atoi(tembusan)
	// 	tembusanStruct.JabatanID = intVar
	// 	if err := tx.Raw("INSERT INTO master_tembusan SET master_surat_id = ?", surat.ID).Error; err != nil {
	// 		tx.Rollback()
	// 		return false, err
	// 	}
	// }

	tx.Commit()

	return true, nil
}
