package report_uangmasuk

type ParamReportSPP struct {
	Tahun_akademik string `form:"tahun_akademik" json:"tahun_akademik" binding:"required"`
	Nis_siswa      string `form:"nis_siswa" json:"nis_siswa"`
	Periode_bayar1 string `form:"periode_bayar1" json:"periode_bayar1"`
	Periode_bayar2 string `form:"periode_bayar2" json:"periode_bayar2"`
	Tgl_bayar1     string `form:"tgl_bayar1" json:"tgl_bayar1"`
	Tgl_bayar2     string `form:"tgl_bayar2" json:"tgl_bayar2"`
}

type GetDataHeaderSPP struct {
	Nm_group       string
	Nm_kategori    string
	Tahun_akademik string
	Nis_siswa      string
	Nm_siswa       string
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
}
