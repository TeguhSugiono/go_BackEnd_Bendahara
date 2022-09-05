package master_tahun_akademik

func FormatShowData(table []ListData) []ReturnData {
	arraydata := []ReturnData{}
	for _, resultdata := range table {
		arraytemporary := FormatTampungData(resultdata)
		arraydata = append(arraydata, arraytemporary)
	}

	return arraydata
}

type ReturnData struct {
	Id_tahun       string `json:"id_tahun"`
	Tahun_akademik string `json:"tahun_akademik"`
}

type ListData struct {
	Id_tahun       string
	Tahun_akademik string
}

func FormatTampungData(table ListData) ReturnData {
	arraydata := ReturnData{}
	arraydata.Id_tahun = table.Id_tahun
	arraydata.Tahun_akademik = table.Tahun_akademik
	return arraydata
}
