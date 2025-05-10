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

type UserItem struct {
	UserID    int64  `gorm:"column:user_id" json:"user_id"`
	UserName  string `gorm:"column:user_name" json:"user_name"`
	StudentID string `gorm:"column:student_id" json:"student_id"`
	Grade     int    `gorm:"column:grade" json:"grade"`
	Status    int    `gorm:"column:status" json:"status"`
	Email     string `gorm:"column:email" json:"email"`
	Avatar    string `gorm:"column:avatar" json:"avatar"`
}
