package master_kategori_uang

import "time"

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
