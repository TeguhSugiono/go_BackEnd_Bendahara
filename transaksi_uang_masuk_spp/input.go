package transaksi_uang_masuk_spp

type ParamInputSPP struct {
	Kd_group       int     `form:"kd_group" json:"kd_group" binding:"required"`
	Kd_kategori    int     `form:"kd_kategori" json:"kd_kategori" binding:"required"`
	Kd_trans_masuk int     `form:"kd_trans_masuk" json:"kd_trans_masuk"`
	Nis_siswa      string  `form:"nis_siswa" json:"nis_siswa" binding:"required"`
	Nm_kelas       string  `form:"nm_kelas" json:"nm_kelas" binding:"required"`
	Tahun_akademik string  `form:"tahun_akademik" json:"tahun_akademik" binding:"required"`
	Total_biaya    float64 `form:"total_biaya" json:"total_biaya"`
	Sisa_biaya     float64 `form:"sisa_biaya" json:"sisa_biaya"`
	Keterangan     string  `form:"keterangan" json:"keterangan"`
	Flag_aktif     int     `form:"flag_aktif" json:"flag_aktif"`
	Created_on     string  `form:"created_on" json:"created_on"`
	Created_by     string  `form:"created_by" json:"created_by"`
	Edited_on      string  `form:"edited_on" json:"edited_on"`
	Edited_by      string  `form:"edited_by" json:"edited_by"`
}
