package unitkerja

type MasterUnitKerja struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type Tabler interface {
	TableName() string
}

func (MasterUnitKerja) TableName() string {
	return "master_unit_kerja"
}
