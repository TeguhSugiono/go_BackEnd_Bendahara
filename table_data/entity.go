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

type Tbl_kategori_uangs struct {
	Kd_group    int
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
	Kd_group        int
	Kd_kategori     int
	Kd_trans_masuk  int `gorm:"primary_key;auto_increment;not_null" json:"kd_trans_masuk"`
	Nis_siswa       string
	Nm_siswa        string
	Tgl_trans_masuk time.Time
	Total_biaya     float64
	Sisa_biaya      float64
	Keterangan      string
	Flag_aktif      int
	Created_on      time.Time
	Created_by      string
	Edited_on       time.Time
	Edited_by       string
}

type Tbl_trans_uang_masuk_spp_details struct {
	Kd_trans_masuk        int
	Kd_trans_masuk_detail int `gorm:"primary_key;auto_increment;not_null" json:"kd_trans_masuk_detail"`
	Seqno                 int
	Periode_bayar         string
	Tgl_bayar             time.Time
	Jml_bayar             float64
	Keterangan            string
	Flag_aktif            int
	Created_on            time.Time
	Created_by            string
	Edited_on             time.Time
	Edited_by             string
}
