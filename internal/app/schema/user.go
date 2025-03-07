package schema

type UserRegister struct {
	UserName  string
	Password  string
	StudentID uint
	Avatar    string
	Email     string
	Grade     int
}

type UserLogin struct {
	Email    string
	Password string
}
