package user

type DataTokenInput struct {
	Id_user  int
	Username string
	Password string
}

type SignUpInput struct {
	Username string `form:"username" binding:"required"`
	Password string `form:"password" binding:"required"`
}

type LoginInput struct {
	Username string `form:"username" binding:"required"`
	Password string `form:"password" binding:"required"`
}
