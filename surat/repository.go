package surat

import "gorm.io/gorm"

type Repository interface {
	Save(masterSurat MasterSurat) (MasterSurat, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) Save(masterSurat MasterSurat) (MasterSurat, error) {
	err := r.db.Create(&masterSurat).Error
	if err != nil {
		return masterSurat, err
	}

	return masterSurat, nil
}
