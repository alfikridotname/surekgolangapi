package naskah

type NaskahFormatter struct {
	ID   int    `json:"id"`
	Nama string `json:"nama"`
}

func FormatNaskah(masterNaskah MasterNaskah) NaskahFormatter {
	return NaskahFormatter{
		ID:   masterNaskah.ID,
		Nama: masterNaskah.Nama,
	}
}

func FormatMultipleNaskah(MasterNaskah []MasterNaskah) []NaskahFormatter {
	var naskahFormatter []NaskahFormatter
	for _, masterNaskah := range MasterNaskah {
		naskahFormatter = append(naskahFormatter, FormatNaskah(masterNaskah))
	}
	return naskahFormatter
}
