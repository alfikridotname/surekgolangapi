package unitkerja

type UnitKerjaFormatter struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func FormatUnitKerja(unitKerja MasterUnitKerja) UnitKerjaFormatter {
	var unitKerjaFormatter UnitKerjaFormatter
	unitKerjaFormatter.ID = unitKerja.ID
	unitKerjaFormatter.Name = unitKerja.Name
	return unitKerjaFormatter
}

func FormatMultipleUnitKerja(masterUnitKerja []MasterUnitKerja) []UnitKerjaFormatter {
	var unitKerjaFormatter []UnitKerjaFormatter
	for _, unitKerja := range masterUnitKerja {
		unitKerjaFormatter = append(unitKerjaFormatter, FormatUnitKerja(unitKerja))
	}
	return unitKerjaFormatter
}
