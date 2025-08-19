package master_sett_periode

func FormatShowData(table []ListData) []ReturnData {
	arraydata := []ReturnData{}
	for _, resultdata := range table {
		arraytemporary := FormatTampungData(resultdata)
		arraydata = append(arraydata, arraytemporary)
	}

	return arraydata
}

type ReturnData struct {
	//Id_conf        int     `json:"id_conf"`
	//Seqno          int     `json:"seqno"`
	//Kd_bulan       string  `json:"kd_bulan"`
	//Tahun          int     `json:"tahun"`
	Nm_sett        string  `json:"nm_sett"`
	Tahun_akademik string  `json:"tahun_akademik"`
	Nm_kelas       string  `json:"nm_kelas"`
	Biaya_spp      float64 `json:"biaya_spp"`
}

type ListData struct {
	//Id_conf        int
	//Seqno          int
	//Kd_bulan       string
	//Tahun          int
	Nm_sett        string
	Tahun_akademik string
	Nm_kelas       string
	Biaya_spp      float64
}

func FormatTampungData(table ListData) ReturnData {
	arraydata := ReturnData{}
	//arraydata.Id_conf = table.Id_conf
	//arraydata.Seqno = table.Seqno
	//arraydata.Kd_bulan = table.Kd_bulan
	//arraydata.Tahun = table.Tahun
	arraydata.Nm_sett = table.Nm_sett
	arraydata.Tahun_akademik = table.Tahun_akademik
	arraydata.Nm_kelas = table.Nm_kelas
	arraydata.Biaya_spp = table.Biaya_spp
	return arraydata
}
