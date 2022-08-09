package user

type UserFormatter struct {
	Username string `json:"username"`
	Token    string `json:"token"`
}

func FormatUser(user Tbl_user, token string) UserFormatter {
	formatter := UserFormatter{
		Username: user.Username,
		Token:    token,
	}
	return formatter
}
