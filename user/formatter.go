package user

type UserFormatter struct {
	Username  string `json:"username"`
	Full_name string `json:"full_name"`
	Token     string `json:"token"`
}

func FormatUser(user Tbl_user, token string) UserFormatter {
	formatter := UserFormatter{
		Username:  user.Username,
		Full_name: user.Full_name,
		Token:     token,
	}
	return formatter
}
