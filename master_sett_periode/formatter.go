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
	Kd_periode_spp int    `json:"kd_periode_spp"`
	Seqno          int    `json:"seqno"`
	Kd_bulan       string `json:"kd_bulan"`
	Tahun          int    `json:"tahun"`
	Nm_sett        string `json:"nm_sett"`
}

type ListData struct {
	Kd_periode_spp int
	Seqno          int
	Kd_bulan       string
	Tahun          int
	Nm_sett        string
}

func FormatTampungData(table ListData) ReturnData {
	arraydata := ReturnData{}
	arraydata.Kd_periode_spp = table.Kd_periode_spp
	arraydata.Seqno = table.Seqno
	arraydata.Kd_bulan = table.Kd_bulan
	arraydata.Tahun = table.Tahun
	arraydata.Nm_sett = table.Nm_sett
	return arraydata
}
