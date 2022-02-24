package notifikasi

import "gorm.io/gorm"

type Repository interface {
	Save(notifikasi *MasterNotifikasi) (bool, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) Save(notifikasi MasterNotifikasi) (bool, error) {
	err := r.db.Create(&notifikasi).Error
	if err != nil {
		return false, err
	}

	return true, nil
}
