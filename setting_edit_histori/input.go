package setting_edit_histori

type StatusOpen struct {
	Open string `form:"open" json:"open" binding:"required"`
}
