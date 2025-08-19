package user

import "time"

type Tbl_user struct {
	Id_user    int `gorm:"primary_key;auto_increment;not_null" json:"id_user"`
	Username   string
	Password   string
	Full_name  string
	Created_on time.Time
}
