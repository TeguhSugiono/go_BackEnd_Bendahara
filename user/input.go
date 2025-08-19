package user

type DataTokenInput struct {
	Id_user   int
	Username  string
	Password  string
	Full_name string
}

type SignUpInput struct {
	Username  string `form:"username" binding:"required"`
	Password  string `form:"password" binding:"required"`
	Full_name string `form:"full_name" binding:"required"`
}

type LoginInput struct {
	Username string `form:"username" binding:"required"`
	Password string `form:"password" binding:"required"`
}

type ParamGantiPassword struct {
	Username     string `form:"username" binding:"required"`
	Password_Old string `form:"password_Old,password_old" binding:"required"`
	Password_New string `form:"password_New,password_new" binding:"required"`
}
