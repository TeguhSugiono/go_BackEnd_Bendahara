package table_data

import "time"

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
	Nm_detail  string
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
	Kd_periode_spp int
	Seqno          int
	Kd_bulan       string
	Tahun          int
	Nm_sett        string
	Flag_aktif     int
	Created_on     time.Time
	Created_by     string
	Edited_on      time.Time
	Edited_by      string
}

//Baca Ketabel akademik di SiA
// type Tbl_tahun_akademik struct {
// 	Id_tahun       int
// 	Tahun_akademik string
// 	Flag_tahun     int
// }
