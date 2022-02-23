package unitkerja

import "gorm.io/gorm"

type Repository interface {
	FindAll(unitKerjaID int) ([]MasterUnitKerja, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindAll(unitKerjaID int) ([]MasterUnitKerja, error) {
	var masterUnitKerja []MasterUnitKerja
	err := r.db.Select("id, name").Where("id <> ?", unitKerjaID).Find(&masterUnitKerja).Error
	if err != nil {
		return masterUnitKerja, err
	}

	return masterUnitKerja, nil
}
