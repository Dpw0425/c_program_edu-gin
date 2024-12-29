package schema

type UserRegister struct {
	NickName string
	Password string
	Avatar   string
	Email    string
}

type UserLogin struct {
	Email    string
	Password string
}
