package pejabat

type PejabatFormatter struct {
	JabatanID int    `json:"jabatan_id"`
	JabatanNM string `json:"jabatan_nm"`
	Nama      string `json:"nama"`
}

func FormatPejabat(signer MasterPejabat) PejabatFormatter {
	return PejabatFormatter{
		JabatanID: signer.JabatanID,
		JabatanNM: signer.JabatanNM,
		Nama:      signer.Nama,
	}
}

func FormatMultiplePejabat(masterPejabat []MasterPejabat) []PejabatFormatter {
	var pejabatFormatter []PejabatFormatter
	for _, pejabat := range masterPejabat {
		pejabatFormatter = append(pejabatFormatter, FormatPejabat(pejabat))
	}
	return pejabatFormatter
}
