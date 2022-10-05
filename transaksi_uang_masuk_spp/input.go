package transaksi_uang_masuk_spp

type ParamInputSPP struct {
	Kd_group       int     `form:"kd_group" json:"kd_group" binding:"required"`
	Kd_kategori    int     `form:"kd_kategori" json:"kd_kategori" binding:"required"`
	Kd_trans_masuk int     `form:"kd_trans_masuk" json:"kd_trans_masuk"`
	Nis_siswa      string  `form:"nis_siswa" json:"nis_siswa" binding:"required"`
	Nm_kelas       string  `form:"nm_kelas" json:"nm_kelas" binding:"required"`
	Tahun_akademik string  `form:"tahun_akademik" json:"tahun_akademik" binding:"required"`
	Total_biaya    float64 `form:"total_biaya" json:"total_biaya"`
	Total_bayar    float64 `form:"total_bayar" json:"total_bayar"`
	Sisa_biaya     float64 `form:"sisa_biaya" json:"sisa_biaya"`
	Keterangan     string  `form:"keterangan" json:"keterangan"`
	Flag_aktif     int     `form:"flag_aktif" json:"flag_aktif"`
	Created_on     string  `form:"created_on" json:"created_on"`
	Created_by     string  `form:"created_by" json:"created_by"`
	Edited_on      string  `form:"edited_on" json:"edited_on"`
	Edited_by      string  `form:"edited_by" json:"edited_by"`
}

type ParamEditSPPDetail struct {
	Tgl_bayar  string  `form:"tgl_bayar" json:"tgl_bayar" binding:"required"`
	Jml_bayar  float64 `form:"jml_bayar" json:"jml_bayar" binding:"number"`
	Keterangan string  `form:"keterangan" json:"keterangan" binding:"required"`
	Edited_on  string  `form:"edited_on" json:"edited_on"`
	Edited_by  string  `form:"edited_by" json:"edited_by"`
}

type ParamChangeNmKelas struct {
	Tahun_akademik string `form:"tahun_akademik" json:"tahun_akademik" binding:"required"`
	Nm_kelas       string `form:"nm_kelas" json:"nm_kelas" binding:"required"`
}

type GetIdAndNameKelas struct {
	Id_kelas string
	Nm_kelas string
}

type GetNisAndNameSiswa struct {
	Nis      string
	Nm_siswa string
}

type ParamChangeSiswa struct {
	Tahun_akademik string `form:"tahun_akademik" json:"tahun_akademik" binding:"required"`
	Nm_kelas       string `form:"nm_kelas" json:"nm_kelas" binding:"required"`
	Nis_siswa      string `form:"nis_siswa" json:"nis_siswa" binding:"required"`
}

type GetDataUmSpp struct {
	// Kd_trans_masuk        int
	Kd_trans_masuk_detail int
	Seqno                 int
	Periode_bayar         string
	Tgl_bayar             string
	Jml_bayar             float64
	Keterangan            string
}

type GetBiayaAndSisa struct {
	Kd_trans_masuk int
	Total_biaya    float64
	Total_bayar    float64
	Sisa_biaya     float64
	Detail         interface{}
}
