package table_data

import (
	"time"
)

type TableData struct {
	Total     int64 `json:"total"`
	Page      int   `json:"page"`
	Last_page int   `json:"last_page"`
}

type Tbl_group_kategoris struct {
	Kd_jenis   int
	Kd_group   int `gorm:"primary_key;auto_increment;not_null" json:"kd_group"`
	Nm_group   string
	Nm_header  string
	Flag_aktif int
	Created_on time.Time
	Created_by string
	Edited_on  time.Time
	Edited_by  string
}

type Tbl_jenis_trans struct {
	Kd_jenis    int `gorm:"primary_key;auto_increment;not_null" json:"kd_jenis"`
	Proses_uang string
	Flag_aktif  int
	Created_on  time.Time
	Created_by  string
	Edited_on   time.Time
	Edited_by   string
}

type Tbl_link_kategoris struct {
	Id_link     int `gorm:"primary_key;auto_increment;not_null" json:"id_link"`
	Link_name   string
	Kd_group    int
	Kd_kategori int
}

type Tbl_kategori_uangs struct {
	Kd_jenis    int
	Kd_kategori int `gorm:"primary_key;auto_increment;not_null" json:"kd_kategori"`
	Nm_kategori string
	Nm_detail   string
	Flag_aktif  int
	Created_on  time.Time
	Created_by  string
	Edited_on   time.Time
	Edited_by   string
}

type Tbl_sub_kategori_uangs struct {
	Kd_kategori     int
	Kd_sub_kategori int `gorm:"primary_key;auto_increment;not_null" json:"kd_sub_kategori"`
	Nm_sub_kategori string
	Flag_aktif      int
	Created_on      time.Time
	Created_by      string
	Edited_on       time.Time
	Edited_by       string
}

type Tbl_conf_periode_spps struct {
	Id_conf        int `gorm:"primary_key;auto_increment;not_null" json:"id_conf"`
	Seqno          int
	Kd_bulan       string
	Tahun          int
	Nm_sett        string
	Tahun_akademik string
	Flag_aktif     int
	Created_on     time.Time
	Created_by     string
	Edited_on      time.Time
	Edited_by      string
	Nm_kelas       string
	Biaya_spp      float64
}

type Tbl_sett_periode_spps struct {
	Kd_periode_spp int
	Kd_sett_spp    int `gorm:"primary_key;auto_increment;not_null" json:"kd_sett_spp"`
	Nm_kelas       string
	Biaya_spp      float64
	Keterangan     string
	Flag_aktif     int
	Created_on     time.Time
	Created_by     string
	Edited_on      time.Time
	Edited_by      string
}

type Tbl_trans_uang_masuk_spp_headers struct {
	Kd_group       int
	Kd_kategori    int
	Kd_trans_masuk int `gorm:"primary_key;auto_increment;not_null" json:"kd_trans_masuk"`
	Nis_siswa      string
	Nm_kelas       string
	Tahun_akademik string
	Total_biaya    float64
	Total_bayar    float64
	Sisa_biaya     float64
	Keterangan     string
	Flag_aktif     int
	Created_on     time.Time
	Created_by     string
	Edited_on      time.Time
	Edited_by      string
}

type Tbl_trans_uang_masuk_spp_details struct {
	Kd_trans_masuk        int
	Kd_trans_masuk_detail int `gorm:"primary_key;auto_increment;not_null" json:"kd_trans_masuk_detail"`
	Seqno                 int
	Periode_bayar         string
	Tgl_bayar             string
	Jml_tagihan           float64
	Jml_bayar             float64
	Keterangan            string
	Flag_aktif            int
	Created_on            time.Time
	Created_by            string
	Edited_on             time.Time
	Edited_by             string
	Kd_pembayaran         int
}

type Tbl_trans_uang_masuk_ppdb_headers struct {
	Kd_group            int
	Kd_kategori         int
	Kd_trans_masuk_ppdb int `gorm:"primary_key;auto_increment;not_null" json:"kd_trans_masuk_ppdb"`
	Nik                 string
	Tgldaftar           string
	Tahun_daftar        string
	Tahun_akademik      string
	Total_biaya         float64
	Total_bayar         float64
	Sisa_biaya          float64
	Keterangan          string
	Flag_aktif          int
	Created_on          time.Time
	Created_by          string
	Edited_on           time.Time
	Edited_by           string
}

type Tbl_trans_uang_masuk_ppdb_details struct {
	Kd_trans_masuk_ppdb        int
	Kd_trans_masuk_detail_ppdb int `gorm:"primary_key;auto_increment;not_null" json:"kd_trans_masuk_detail_ppdb"`
	Seqno                      int
	Kategori_biaya_ppdb        string
	Tgl_bayar                  string
	Jml_bayar                  float64
	Keterangan                 string
	Flag_aktif                 int
	Created_on                 time.Time
	Created_by                 string
	Edited_on                  time.Time
	Edited_by                  string
	Kd_pembayaran              int
}

type Tbl_trans_uang_masuk_siswa_headers struct {
	Kd_group             int
	Kd_kategori          int
	Kd_trans_masuk_siswa int `gorm:"primary_key;auto_increment;not_null" json:"kd_trans_masuk_siswa"`
	Nis_siswa            string
	Nm_kelas             string
	Tahun_akademik       string
	Total_biaya          float64
	Total_bayar          float64
	Sisa_biaya           float64
	Keterangan           string
	Flag_aktif           int
	Created_on           time.Time
	Created_by           string
	Edited_on            time.Time
	Edited_by            string
}

type Tbl_trans_uang_masuk_siswa_details struct {
	Kd_trans_masuk_siswa        int
	Kd_trans_masuk_detail_siswa int `gorm:"primary_key;auto_increment;not_null" json:"kd_trans_masuk_detail_siswa"`
	Seqno                       int
	Tgl_bayar                   string
	Jml_bayar                   float64
	Keterangan                  string
	Flag_aktif                  int
	Created_on                  time.Time
	Created_by                  string
	Edited_on                   time.Time
	Edited_by                   string
	Kd_pembayaran               int
}

type Tbl_biaya_masuk_keluars struct {
	Kd_biaya_kategori int `gorm:"primary_key;auto_increment;not_null" json:"kd_biaya_kategori"`
	Kd_kategori       int
	Jml_biaya         float64
}

type Tbl_trans_uang_masuk_lain_headers struct {
	Kd_group            int
	Kd_trans_masuk_lain int `gorm:"primary_key;auto_increment;not_null" json:"kd_trans_masuk_lain"`
	No_document         string
	Tgl_document        string
	Total_biaya         float64
	Total_bayar         float64
	Sisa_biaya          float64
	Keterangan          string
	Flag_aktif          int
	Created_on          time.Time
	Created_by          string
	Edited_on           time.Time
	Edited_by           string
}

type Tbl_trans_uang_masuk_lain_details struct {
	Kd_trans_masuk_lain        int
	Kd_trans_masuk_detail_lain int `gorm:"primary_key;auto_increment;not_null" json:"kd_trans_masuk_detail_lain"`
	Seqno                      int
	Tgl_bayar                  string
	Jml_bayar                  float64
	Keterangan                 string
	Flag_aktif                 int
	Created_on                 time.Time
	Created_by                 string
	Edited_on                  time.Time
	Edited_by                  string
	Kd_pembayaran              int
	Kd_kategori                int
}

type Tbl_trans_uang_keluar_pra_headers struct {
	Kd_group        int
	Kd_kategori     int
	Kd_trans_keluar int `gorm:"primary_key;auto_increment;not_null" json:"kd_trans_keluar"`
	No_document     string
	Tgl_document    string
	Total_biaya     float64
	Total_bayar     float64
	Sisa_biaya      float64
	Keterangan      string
	Flag_aktif      int
	Created_on      time.Time
	Created_by      string
	Edited_on       time.Time
	Edited_by       string
}

type Tbl_trans_uang_keluar_pra_details struct {
	Kd_trans_keluar        int
	Kd_trans_keluar_detail int `gorm:"primary_key;auto_increment;not_null" json:"kd_trans_keluar_detail"`
	Seqno                  int
	Jml_bayar              float64
	Keterangan             string
	Flag_aktif             int
	Created_on             time.Time
	Created_by             string
	Edited_on              time.Time
	Edited_by              string
}

type Tbl_trans_uang_keluar_pra_act_headers struct {
	Kd_group        int
	Kd_kategori     int
	Kd_trans_keluar int `gorm:"primary_key;auto_increment;not_null" json:"kd_trans_keluar"`
	Kd_proses       string
	No_document     string
	Tgl_document    string
	Total_biaya     float64
	Total_bayar     float64
	Sisa_biaya      float64
	Keterangan      string
	Flag_aktif      int
	Created_on      time.Time
	Created_by      string
	Edited_on       time.Time
	Edited_by       string
}

type Tbl_trans_uang_keluar_pra_act_details struct {
	Kd_trans_keluar        int
	Kd_trans_keluar_detail int `gorm:"primary_key;auto_increment;not_null" json:"kd_trans_keluar_detail"`
	Seqno                  int
	Kd_post_uang_masuk     string
	Tgl_bayar              string
	Jml_bayar              float64
	Keterangan             string
	Flag_aktif             int
	Created_on             time.Time
	Created_by             string
	Edited_on              time.Time
	Edited_by              string
	Kd_pembayaran          int
}

type Tbl_trans_uang_keluar_act_headers struct {
	Kd_group        int
	Kd_trans_keluar int `gorm:"primary_key;auto_increment;not_null" json:"kd_trans_keluar"`
	Kd_proses       string
	No_document     string
	Tgl_document    string
	Total_biaya     float64
	Total_bayar     float64
	Sisa_biaya      float64
	Keterangan      string
	Flag_aktif      int
	Created_on      time.Time
	Created_by      string
	Edited_on       time.Time
	Edited_by       string
}

type Tbl_trans_uang_keluar_act_details struct {
	Kd_trans_keluar        int
	Kd_trans_keluar_detail int `gorm:"primary_key;auto_increment;not_null" json:"kd_trans_keluar_detail"`
	Seqno                  int
	Kd_post_uang_masuk     string
	Jml_bayar              float64
	Keterangan             string
	Flag_aktif             int
	Created_on             time.Time
	Created_by             string
	Edited_on              time.Time
	Edited_by              string
	Kd_pembayaran          int
	Kd_kategori            int
}

type Tbl_open_lock_historis struct {
	Open       string
	Request_by string
	Request_on time.Time
}

type Tbl_runcode_documents struct {
	Id             int
	Kode           string
	Tahun          string
	Bulan          string
	Nomor          int
	Generate_nomor string
	Namahalaman    string
}
