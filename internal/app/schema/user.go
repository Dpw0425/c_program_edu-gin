package schema

type UserRegister struct {
	UserName  string
	Password  string
	StudentID string
	Avatar    string
	Email     string
	Grade     int
}

type UserLogin struct {
	Email    string
	Password string
}
