package transaksi_uang_masuk_lainlain

type ParamChangeKategori struct {
	Kd_group string `form:"kd_group" json:"kd_group" binding:"required"`
}

// type ParamGetSiswaAdd struct {
// 	Nis         string `form:"nis" json:"nis" binding:"required"`
// 	Kd_kategori int    `form:"kd_kategori" json:"kd_kategori" binding:"required"`
// }

// type ListDataView struct {
// 	Kd_group                    int
// 	Nm_group                    string
// 	Kd_kategori                 int
// 	Nm_kategori                 string
// 	Kd_trans_masuk_lain         int
// 	Tahun_akademik              string
// 	Nis_siswa                   string
// 	Nm_siswa                    string
// 	Nm_kelas                    string
// 	Total_biaya                 float64
// 	Total_bayar                 float64
// 	Sisa_biaya                  float64
// 	Keterangan                  string
// 	Kd_trans_masuk_detail_siswa int
// 	Seqno                       int
// 	Tgl_bayar                   string
// 	Jml_bayar                   float64
// 	Keterangandetail            string
// }

// type ListAddSiswa struct {
// 	Id_tahun       int
// 	Tahun_akademik string
// 	Id_kelas       string
// 	Nm_kelas       string
// 	Total_biaya    float64
// 	Total_bayar    float64
// 	Sisa_biaya     float64
// }

type ParamInputTransaksi struct {
	Kd_group            int     `form:"kd_group" json:"kd_group" binding:"required"`
	Kd_kategori         int     `form:"kd_kategori" json:"kd_kategori" binding:"required"`
	Kd_trans_masuk_lain int     `form:"kd_trans_masuk_lain" json:"kd_trans_masuk_lain"`
	Tgl_transaksi       string  `form:"tgl_transaksi" json:"tgl_transaksi" binding:"required"`
	Total_biaya         float64 `form:"total_biaya" json:"total_biaya" binding:"number"`
	Total_bayar         float64 `form:"total_bayar" json:"total_bayar"`
	Sisa_biaya          float64 `form:"sisa_biaya" json:"sisa_biaya"`
	Keterangan          string  `form:"keterangan" json:"keterangan"`
	Flag_aktif          int     `form:"flag_aktif" json:"flag_aktif"`
	Created_on          string  `form:"created_on" json:"created_on"`
	Created_by          string  `form:"created_by" json:"created_by"`
	Edited_on           string  `form:"edited_on" json:"edited_on"`
	Edited_by           string  `form:"edited_by" json:"edited_by"`
}

type ParamInputTransaksiEdit struct {
	Kd_group      int     `form:"kd_group" json:"kd_group" binding:"required"`
	Kd_kategori   int     `form:"kd_kategori" json:"kd_kategori" binding:"required"`
	Tgl_transaksi string  `form:"tgl_transaksi" json:"tgl_transaksi" binding:"required"`
	Total_biaya   float64 `form:"total_biaya" json:"total_biaya" binding:"required,number"`
	Total_bayar   float64 `form:"total_bayar" json:"total_bayar"`
	Sisa_biaya    float64 `form:"sisa_biaya" json:"sisa_biaya"`
	Keterangan    string  `form:"keterangan" json:"keterangan"`
	Flag_aktif    int     `form:"flag_aktif" json:"flag_aktif"`
	Created_on    string  `form:"created_on" json:"created_on"`
	Created_by    string  `form:"created_by" json:"created_by"`
	Edited_on     string  `form:"edited_on" json:"edited_on"`
	Edited_by     string  `form:"edited_by" json:"edited_by"`
}

type GetDataUmSiswa struct {
	Kd_trans_masuk_detail_siswa int
	Seqno                       int
	Tgl_bayar                   string
	Jml_bayar                   float64
	Keterangan                  string
}

type GetBiayaAndSisa struct {
	Kd_trans_masuk_lain int
	Kd_group            int
	Nm_group            string
	Kd_kategori         int
	Nm_kategori         string
	Tgl_transaksi       string
	Total_biaya         float64
	Total_bayar         float64
	Sisa_biaya          float64
	Detail              interface{}
}

type ParamEditUmSiswaDetail struct {
	Tgl_bayar  string  `form:"tgl_bayar" json:"tgl_bayar" binding:"required"`
	Jml_bayar  float64 `form:"jml_bayar" json:"jml_bayar" binding:"number"`
	Keterangan string  `form:"keterangan" json:"keterangan" binding:"required"`
	Edited_on  string  `form:"edited_on" json:"edited_on"`
	Edited_by  string  `form:"edited_by" json:"edited_by"`
}

type ParamAddDetail struct {
	Kd_trans_masuk_lain int `form:"kd_trans_masuk_lain" json:"kd_trans_masuk_lain"`
}

type ParamChangeSiswa struct {
	Tgl_transaksi1 string `form:"tgl_transaksi1" json:"tgl_transaksi1" binding:"required"`
	Tgl_transaksi2 string `form:"tgl_transaksi2" json:"tgl_transaksi2" binding:"required"`
}
