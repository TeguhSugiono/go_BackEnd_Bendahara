package master_sett_periode

type ConfPeriodeInput struct {
	Kd_periode_spp int    `form:"kd_periode_spp"`
	Seqno          int    `form:"seqno"`
	Kd_bulan       string `form:"kd_bulan"`
	Tahun          int    `form:"tahun"`
	Nm_sett        string `form:"nm_sett"`
	Tahun_akademik string `form:"tahun_akademik"`
	Flag_aktif     int    `form:"flag_aktif"`
	Created_on     string `form:"created_on"`
	Created_by     string `form:"created_by"`
	Edited_on      string `form:"edited_on"`
	Edited_by      string `form:"edited_by"`
}

type InputTahunAkademik struct {
	Tahun_akademik string  `form:"tahun_akademik" json:"tahun_akademik" binding:"required"`
	Nm_kelas       string  `form:"nm_kelas" json:"nm_kelas" binding:"required"`
	Biaya_spp      float64 `form:"biaya_spp" json:"biaya_spp" binding:"required"`
}

type DeleteTahunAkademik struct {
	Tahun_akademik string `form:"tahun_akademik" json:"tahun_akademik" binding:"required"`
	Nm_kelas       string `form:"nm_kelas" json:"nm_kelas" binding:"required"`
}

type EditTahunAkademik struct {
	Biaya_spp float64 `form:"biaya_spp" json:"biaya_spp" binding:"required"`
}

type CekDataSettPeriode struct {
	Jumlah int
}
