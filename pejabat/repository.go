package pejabat

import "gorm.io/gorm"

type Repository interface {
	FindByUnitKerjaID(unitKerjaID int, kategori string) ([]MasterPejabat, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindByUnitKerjaID(unitKerjaID int, kategori string) ([]MasterPejabat, error) {
	var masterPejabat []MasterPejabat
	eselonID := []int{12, 21, 22, 31, 32}
	if kategori == "atasnama" {
		eselonID = []int{12, 21, 22}
	}
	err := r.db.Raw(`SELECT
						A.jabatan_id,
						A.jabatan_nm,
						A.nama,
						A.kategori,
						A.nip,
						A.eselon_id
					FROM
						master_pegawai A
						LEFT JOIN asisten_unit_kerja B ON A.unit_kerja_id = B.unit_kerja_id
					WHERE 
						(eselon_id IN (?)
						AND A.unit_kerja_id = ?) 
						OR (A.unit_kerja_id IN (0)) 
						OR (A.unit_kerja_id = 1 AND A.jabatan_nm ILIKE '%SEKRE%')
						OR (A.jabatan_id IN (B.jabatan_id) AND A.jabatan_nm ILIKE '%ASISTEN%')
					ORDER BY
						eselon_id ASC`, eselonID, unitKerjaID).Scan(&masterPejabat).Error
	if err != nil {
		return masterPejabat, err
	}

	return masterPejabat, nil
}
