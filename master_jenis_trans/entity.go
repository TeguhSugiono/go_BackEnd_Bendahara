package master_jenis_trans

import "time"

type Tbl_jenis_trans struct {
	Kd_jenis    int `gorm:"primary_key;auto_increment;not_null" json:"kd_jenis"`
	Proses_uang string
	Flag_aktif  int
	Created_on  time.Time
	Created_by  string
	Edited_on   time.Time
	Edited_by   string
}
