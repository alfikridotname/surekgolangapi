package pegawai

import "gorm.io/gorm"

type Repository interface {
	FindUser(unitKerjaID int, jabatanID int) (UserID, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindUser(unitKerjaID, jabatanID int) (UserID, error) {
	var userID UserID
	// err := r.db.Select("B.id").Joins("LEFT JOIN users B ON master_pegawai.nip = B.username").Where("master_pegawai.unit_kerja_id = ? AND master_pegawai.jabatan_id = ?", unitKerjaID, jabatanID).Find(&pegawai).Error
	err := r.db.Raw(`SELECT
						B.id
					FROM
						master_pegawai A
						LEFT JOIN users B ON A.nip = B.username
					WHERE
						A.unit_kerja_id = ? 
						AND A.jabatan_id = ?`, unitKerjaID, jabatanID).Scan(&userID).Error
	if err != nil {
		return userID, err
	}

	return userID, nil
}
