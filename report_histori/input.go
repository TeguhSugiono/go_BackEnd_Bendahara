package report_histori

type ListData struct {
	Nis      string
	Nik      string
	Nm_siswa string
}

type ParamSearch struct {
	Nis      string `form:"nis" json:"nis" binding:"required"`
	Nik      string `form:"nik" json:"nik" binding:"required"`
	Nm_siswa string `form:"nm_siswa" json:"nm_siswa" binding:"required"`
}

type DetailHistori struct {
	Nis          string
	Nik          string
	Nm_siswa     string
	DetailPPDB   interface{}
	DetailSPP    interface{}
	DetailumLain interface{}
}
