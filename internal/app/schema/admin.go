package schema

type AdminLogin struct {
	UserName string
	Password string
}

type AdminItem struct {
	ID         uint   `gorm:"column:id" json:"id"`
	UserName   string `gorm:"column:user_name" json:"user_name"`
	TeacherID  string `gorm:"column:teacher_id" json:"teacher_id"`
	Permission string `gorm:"column:permission" json:"permission"`
	Status     int    `gorm:"column:status" json:"status"`
}
