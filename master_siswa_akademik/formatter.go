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

type ListDataSiswa struct {
	Id_siswa    string
	Nm_siswa    string
	Tahun_lulus string
	No_peserta  string
	Nis         string
}

type SearchSiswaLulus struct {
	// Id_siswa    string `form:"id_siswa" json:"id_siswa"`
	// Tahun_lulus string `form:"tahun_lulus" json:"tahun_lulus"`
	No_peserta string `form:"no_peserta" json:"no_peserta"`
}

type ParamSiswaAll struct {
	Nis string `form:"nis" json:"nis"`
}

type ListSiswaAll struct {
	Id_siswa string
	Nis      string
	Nm_siswa string
}
