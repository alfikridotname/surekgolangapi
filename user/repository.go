package user

import "gorm.io/gorm"

type Repository interface {
	FindByID(ID int) (User, error)
	FindByUsername(username string) (User, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindByID(ID int) (User, error) {
	var user User
	err := r.db.Select(`users.id, 
						users.name, 
						master_pegawai.nip,
						master_pegawai.jabatan_id, 
						master_pegawai.unit_kerja_id`).Joins(`LEFT JOIN master_pegawai ON users.username = master_pegawai.nip`).Where("users.id = ?", ID).Find(&user).Error
	if err != nil {
		return user, err
	}

	return user, nil
}

func (r *repository) FindByUsername(username string) (User, error) {
	var user User
	err := r.db.Where("users.username = ?", username).Find(&user).Error
	if err != nil {
		return user, err
	}

	return user, nil
}
