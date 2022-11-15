package transaksi_uang_masuk_siswa

type ParamChangeKategori struct {
	Kd_group string `form:"kd_group" json:"kd_group" binding:"required"`
}

type ParamGetSiswaAdd struct {
	Nis         string `form:"nis" json:"nis" binding:"required"`
	Kd_kategori int    `form:"kd_kategori" json:"kd_kategori" binding:"required"`
}

type ListData struct {
	Id_tahun       int
	Tahun_akademik string
	Id_kelas       string
	Nm_kelas       string
	Total_biaya    float64
	Total_bayar    float64
	Sisa_biaya     float64
}

type ParamInputSiswa struct {
	Kd_group             int     `form:"kd_group" json:"kd_group" binding:"required"`
	Kd_kategori          int     `form:"kd_kategori" json:"kd_kategori" binding:"required"`
	Kd_trans_masuk_siswa int     `form:"kd_trans_masuk_siswa" json:"kd_trans_masuk_siswa"`
	Nis_siswa            string  `form:"nis_siswa" json:"nis_siswa" binding:"required"`
	Nm_kelas             string  `form:"nm_kelas" json:"nm_kelas" binding:"required"`
	Tahun_akademik       string  `form:"tahun_akademik" json:"tahun_akademik" binding:"required"`
	Total_biaya          float64 `form:"total_biaya" json:"total_biaya" binding:"number"`
	Total_bayar          float64 `form:"total_bayar" json:"total_bayar"`
	Sisa_biaya           float64 `form:"sisa_biaya" json:"sisa_biaya"`
	Keterangan           string  `form:"keterangan" json:"keterangan"`
	Flag_aktif           int     `form:"flag_aktif" json:"flag_aktif"`
	Created_on           string  `form:"created_on" json:"created_on"`
	Created_by           string  `form:"created_by" json:"created_by"`
	Edited_on            string  `form:"edited_on" json:"edited_on"`
	Edited_by            string  `form:"edited_by" json:"edited_by"`
}

type ParamInputSiswaEdit struct {
	Kd_group       int     `form:"kd_group" json:"kd_group" binding:"required"`
	Kd_kategori    int     `form:"kd_kategori" json:"kd_kategori" binding:"required"`
	Nis_siswa      string  `form:"nis_siswa" json:"nis_siswa" binding:"required"`
	Nm_kelas       string  `form:"nm_kelas" json:"nm_kelas" binding:"required"`
	Tahun_akademik string  `form:"tahun_akademik" json:"tahun_akademik" binding:"required"`
	Total_biaya    float64 `form:"total_biaya" json:"total_biaya" binding:"required,number"`
	Total_bayar    float64 `form:"total_bayar" json:"total_bayar"`
	Sisa_biaya     float64 `form:"sisa_biaya" json:"sisa_biaya"`
	Keterangan     string  `form:"keterangan" json:"keterangan"`
	Flag_aktif     int     `form:"flag_aktif" json:"flag_aktif"`
	Created_on     string  `form:"created_on" json:"created_on"`
	Created_by     string  `form:"created_by" json:"created_by"`
	Edited_on      string  `form:"edited_on" json:"edited_on"`
	Edited_by      string  `form:"edited_by" json:"edited_by"`
}

type GetDataUmSiswa struct {
	Kd_trans_masuk_detail_siswa int
	Seqno                       int
	Tgl_bayar                   string
	Jml_bayar                   float64
	Keterangan                  string
}

type GetBiayaAndSisa struct {
	Kd_trans_masuk_siswa int
	Total_biaya          float64
	Total_bayar          float64
	Sisa_biaya           float64
	Detail               interface{}
}

type ParamEditUmSiswaDetail struct {
	Tgl_bayar  string  `form:"tgl_bayar" json:"tgl_bayar" binding:"required"`
	Jml_bayar  float64 `form:"jml_bayar" json:"jml_bayar" binding:"number"`
	Keterangan string  `form:"keterangan" json:"keterangan" binding:"required"`
	Edited_on  string  `form:"edited_on" json:"edited_on"`
	Edited_by  string  `form:"edited_by" json:"edited_by"`
}

type ParamAddDetail struct {
	Kd_trans_masuk_siswa int `form:"kd_trans_masuk_siswa" json:"kd_trans_masuk_siswa"`
}
