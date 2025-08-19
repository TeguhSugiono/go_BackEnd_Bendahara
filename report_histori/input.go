package report_histori

type ListData struct {
	Nis      string
	Nik      string
	Nm_siswa string
}

type ParamSearch struct {
	Nis      string `form:"nis" json:"nis" binding:"required"`
	Nik      string `form:"nik" json:"nik"`
	Nm_siswa string `form:"nm_siswa" json:"nm_siswa" binding:"required"`
}

type DetailHistori struct {
	Nis         string
	Nik         string
	Nm_siswa    string
	DataSPP     interface{}
	DataUmSiswa interface{}
	DataPPDB    interface{}
}

type HeaderPPDB struct {
	Nm_group       string
	Nm_kategori    string
	Tgldaftar      string
	Tahun_daftar   string
	Tahun_akademik string
	Total_biaya    float64
	Total_bayar    float64
	Sisa_biaya     float64
	Keterangan     string
	Detail         interface{}
}

type DetailPPDB struct {
	Tgl_bayar                  string
	Jml_bayar                  float64
	Keterangan_detail          string
	Kd_pembayaran              int
	Tipe_pembayaran            string
	Kd_trans_masuk_ppdb        int
	Kd_trans_masuk_detail_ppdb int
}

type GroupTahunKelas struct {
	Tahun_akademik string
	Nm_kelas       string
	HeaderData     interface{}
}

type HeaderSPP struct {
	Nm_group    string
	Nm_kategori string
	Total_biaya float64
	Total_bayar float64
	Sisa_biaya  float64
	Keterangan  string
	DetailData  interface{}
}

type DetailSPP struct {
	Periode_bayar         string
	Tgl_bayar             string
	Jml_tagihan           float64
	Jml_bayar             float64
	Keterangan_detail     string
	Kd_pembayaran         int
	Tipe_pembayaran       string
	Kd_trans_masuk        int
	Kd_trans_masuk_detail int
}

type DetailUmSiswa struct {
	Tgl_bayar                   string
	Jml_bayar                   float64
	Keterangan_detail           string
	Kd_pembayaran               int
	Tipe_pembayaran             string
	Kd_trans_masuk_siswa        int
	Kd_trans_masuk_detail_siswa int
}
