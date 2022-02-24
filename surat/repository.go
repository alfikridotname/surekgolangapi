package surat

import (
	"strconv"
	"strings"
	"surekapi/notifikasi"
	"surekapi/pegawai"

	uuid "github.com/satori/go.uuid"
	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type Repository interface {
	Save(surat MasterSurat, tembusan string, penerimaID string, pemeriksa string, unitKerjaID int) (bool, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) Save(surat MasterSurat, tembusan string, penerimaID string, pemeriksa string, unitKerjaID int) (bool, error) {
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

	if pemeriksa != "" {
		pemeriksaSlice := strings.Split(pemeriksa, ",")
		pemeriksaAktif, _ := strconv.Atoi(pemeriksaSlice[0])

		for _, jabatanID := range pemeriksaSlice {
			masterPemeriksa := MasterPemeriksa{}
			masterPemeriksa.MasterSuratID = surat.ID.String()
			jabatanID, _ := strconv.Atoi(jabatanID)
			masterPemeriksa.JabatanID = jabatanID

			pegawaiRepo := pegawai.NewRepository(tx)
			user, _ := pegawaiRepo.FindUser(unitKerjaID, jabatanID)

			if jabatanID == pemeriksaAktif {
				masterPemeriksa.StatusKoreksi = true
			}

			result := tx.Table("master_pemeriksa").Create(&masterPemeriksa)
			if result.Error != nil {
				tx.Rollback()
				return false, result.Error
			}

			notifikasiRepo := notifikasi.NewRepository(tx)
			notifikasiInput := notifikasi.MasterNotifikasi{}
			notifikasiInput.Data = datatypes.JSON([]byte(`{"kategori": "konsep","surat_id": "` + surat.ID.String() + `"}`))
			notifikasiInput.UserTujuanID = user.ID
			notifikasiInput.CreatedBY = surat.CreatedBy
			ok, err := notifikasiRepo.Save(notifikasiInput)
			if !ok {
				tx.Rollback()
				return false, err
			}
		}
	}

	tx.Commit()

	return true, nil
}
