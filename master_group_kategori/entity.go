package master_group_kategori

import "time"

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
