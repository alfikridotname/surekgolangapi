package naskah

import "gorm.io/gorm"

type Repository interface {
	FindAll() ([]MasterNaskah, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindAll() ([]MasterNaskah, error) {
	var masterNaskah []MasterNaskah
	err := r.db.Where("is_active = ?", true).Order("id asc").Find(&masterNaskah).Error
	if err != nil {
		return masterNaskah, err
	}

	return masterNaskah, nil
}
