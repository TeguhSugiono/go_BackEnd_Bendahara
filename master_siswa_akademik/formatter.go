package master_siswa_akademik

func FormatShowData(table []ListData) []ReturnData {
	arraydata := []ReturnData{}
	for _, resultdata := range table {
		arraytemporary := FormatTampungData(resultdata)
		arraydata = append(arraydata, arraytemporary)
	}

	return arraydata
}

type ReturnData struct {
	Nis      string `json:"nis"`
	Nm_siswa string `json:"nm_siswa"`
}

type ListData struct {
	Nis      string
	Nm_siswa string
}

func FormatTampungData(table ListData) ReturnData {
	arraydata := ReturnData{}
	arraydata.Nis = table.Nis
	arraydata.Nm_siswa = table.Nm_siswa
	return arraydata
}
