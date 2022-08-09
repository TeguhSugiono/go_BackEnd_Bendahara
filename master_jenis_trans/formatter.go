package master_jenis_trans

type JenisTransFormatter struct {
	Kd_Jenis    int    `json:"kd_jenis"`
	Proses_Uang string `json:"proses_uang"`
}

func FormatJenisTranss(table Tbl_jenis_trans) JenisTransFormatter {

	arraydata := JenisTransFormatter{}
	arraydata.Kd_Jenis = table.Kd_jenis
	arraydata.Proses_Uang = table.Proses_uang

	return arraydata
}

func FormatJenisTrans(table []Tbl_jenis_trans) []JenisTransFormatter {
	arraydata := []JenisTransFormatter{}
	for _, resultdata := range table {
		arraytemporary := FormatJenisTranss(resultdata)
		arraydata = append(arraydata, arraytemporary)
	}

	return arraydata
}
