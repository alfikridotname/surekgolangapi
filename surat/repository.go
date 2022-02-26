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
	Save(surat MasterSurat, tembusan string, penerimaID string, pemeriksa string, unitKerjaID int, jabatanID int) (bool, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) Save(surat MasterSurat, tembusan string, penerimaID string, pemeriksa string, unitKerjaID int, inputJabatanID int) (bool, error) {
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
		// Slice data pemeriksa
		pemeriksaSlice := strings.Split(pemeriksa, ",")
		// Tambah pembuat konsep sebagai pemeriksa
		createdJabatanID := inputJabatanID
		createdBy := strconv.Itoa(surat.CreatedBy)
		creatorID := strconv.Itoa(createdJabatanID) + "-" + createdBy
		pemeriksaSlice = append([]string{creatorID}, pemeriksaSlice...)
		// Ambil data pemeriksa aktif
		pemeriksaAktif, _ := strconv.Atoi(pemeriksaSlice[1])

		for index, pemeriksa := range pemeriksaSlice {
			// Split data pemeriksa
			pemeriksaSlice := strings.Split(pemeriksa, "-")
			jabatanIDSlice := pemeriksaSlice[0]
			userIDSlice := pemeriksaSlice[1]

			// Struct master pemeriksa
			masterPemeriksa := MasterPemeriksa{}
			masterPemeriksa.MasterSuratID = surat.ID.String()
			jabatanID, _ := strconv.Atoi(jabatanIDSlice)
			masterPemeriksa.JabatanID = jabatanID
			masterPemeriksa.UserID, _ = strconv.Atoi(userIDSlice)

			// Cari data user berdasarkan jabatan id menggunakan pegawai repository
			pegawaiRepo := pegawai.NewRepository(tx)
			user, _ := pegawaiRepo.FindUser(unitKerjaID, jabatanID)

			// Jika jabatan pmeriksa adalah aktif
			if jabatanID == pemeriksaAktif {
				masterPemeriksa.StatusKoreksi = true
			}

			// Jika pembuat konsep
			if index == 0 {
				masterPemeriksa.StatusPembuat = true
			}

			// Insert ke table master_pemeriksa
			result := tx.Table("master_pemeriksa").Create(&masterPemeriksa)
			if result.Error != nil {
				tx.Rollback()
				return false, result.Error
			}

			// Jika Pemerika adalah aktif
			if index == 1 {
				// Insert ke table master_notifikasi menggunakan repository notifikasi
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
	}

	tx.Commit()

	return true, nil
}
