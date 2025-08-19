package master_kelas_akademik

func FormatShowData(table []ListData) []ReturnData {
	arraydata := []ReturnData{}
	for _, resultdata := range table {
		arraytemporary := FormatTampungData(resultdata)
		arraydata = append(arraydata, arraytemporary)
	}

	return arraydata
}

type ReturnData struct {
	Id_kelas string `json:"id_kelas"`
	Nm_kelas string `json:"nm_kelas"`
}

type ListData struct {
	Id_kelas string
	Nm_kelas string
}

func FormatTampungData(table ListData) ReturnData {
	arraydata := ReturnData{}
	arraydata.Id_kelas = table.Id_kelas
	arraydata.Nm_kelas = table.Nm_kelas
	return arraydata
}
