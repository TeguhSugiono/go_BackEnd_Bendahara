package report_uangmasuk

type ParamReportSPP struct {
	Tahun_akademik string `form:"tahun_akademik" json:"tahun_akademik" binding:"required"`
	Nis_siswa      string `form:"nis_siswa" json:"nis_siswa"`
	Periode_bayar1 string `form:"periode_bayar1" json:"periode_bayar1"`
	Periode_bayar2 string `form:"periode_bayar2" json:"periode_bayar2"`
	Tgl_bayar1     string `form:"tgl_bayar1" json:"tgl_bayar1"`
	Tgl_bayar2     string `form:"tgl_bayar2" json:"tgl_bayar2"`
	Kd_pembayaran  string `form:"kd_pembayaran" json:"kd_pembayaran"`
}

type GetDataHeaderSPP struct {
	Nm_group       string
	Nm_kategori    string
	Tahun_akademik string
	Nis_siswa      string
	Nm_siswa       string
	Nm_kelas       string
	Total_biaya    float64
	Total_bayar    float64
	Sisa_biaya     float64
	Keterangan     string
	Detail         interface{}
}

type GetDataDetailSPP struct {
	Periode_bayar     string
	Tgl_bayar         string
	Jml_tagihan       float64
	Jml_bayar         float64
	Keterangan_detail string
	Tipe_pembayaran   string
}

type ParamReportPPDB struct {
	Tahun_akademik string `form:"tahun_akademik" json:"tahun_akademik" binding:"required"`
	Nik            string `form:"nik" json:"nik"`
	Tgldaftar1     string `form:"tgldaftar1" json:"tgldaftar1"`
	Tgldaftar2     string `form:"tgldaftar2" json:"tgldaftar2"`
	Tgl_bayar1     string `form:"tgl_bayar1" json:"tgl_bayar1"`
	Tgl_bayar2     string `form:"tgl_bayar2" json:"tgl_bayar2"`
	Kd_pembayaran  string `form:"kd_pembayaran" json:"kd_pembayaran"`
}

type GetDataHeaderPPDB struct {
	Nm_group       string
	Nm_kategori    string
	Nik            string
	Nm_siswa       string
	Tgldaftar      string
	Tahun_daftar   string
	Tahun_akademik string
	Total_biaya    float64
	Total_bayar    float64
	Sisa_biaya     float64
	Keterangan     string
	Detail         interface{}
}

type GetDataDetailPPDB struct {
	Tgl_bayar         string
	Jml_bayar         float64
	Keterangan_detail string
	Tipe_pembayaran   string
}

type ParamReportUmSiswa struct {
	Tahun_akademik string `form:"tahun_akademik" json:"tahun_akademik" binding:"required"`
	Nis_siswa      string `form:"nis_siswa" json:"nis_siswa"`
	Nm_kelas       string `form:"nm_kelas" json:"nm_kelas"`
	Tgl_bayar1     string `form:"tgl_bayar1" json:"tgl_bayar1"`
	Tgl_bayar2     string `form:"tgl_bayar2" json:"tgl_bayar2"`
	Kd_pembayaran  string `form:"kd_pembayaran" json:"kd_pembayaran"`
}

type GetDataHeaderUmSiswa struct {
	Nm_group       string
	Nm_kategori    string
	Tahun_akademik string
	Nis_siswa      string
	Nm_siswa       string
	Nm_kelas       string
	Total_biaya    float64
	Total_bayar    float64
	Sisa_biaya     float64
	Keterangan     string
	Detail         interface{}
}

type GetDataDetailUmSiswa struct {
	Tgl_bayar         string
	Jml_bayar         float64
	Keterangan_detail string
	Tipe_pembayaran   string
}

type ParamReportUmLain struct {
	Tgl_document1 string `form:"tgl_document1" json:"tgl_document1" binding:"required"`
	Tgl_document2 string `form:"tgl_document2" json:"tgl_document2" binding:"required"`
	No_document   string `form:"no_document" json:"no_document"`
	Tgl_bayar1    string `form:"tgl_bayar1" json:"tgl_bayar1"`
	Tgl_bayar2    string `form:"tgl_bayar2" json:"tgl_bayar2"`
	Kd_pembayaran string `form:"kd_pembayaran" json:"kd_pembayaran"`
}

type GetDataHeaderUmLain struct {
	Nm_group     string
	Nm_kategori  string
	No_document  string
	Tgl_document string
	Total_biaya  float64
	Total_bayar  float64
	Sisa_biaya   float64
	Keterangan   string
	Detail       interface{}
}

type GetDataDetailUmLain struct {
	Tgl_bayar         string
	Jml_bayar         float64
	Keterangan_detail string
	Tipe_pembayaran   string
}
