package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	UserID    int64  `gorm:"column:user_id;unique" json:"user_id"`           // ID
	UserName  string `gorm:"column:user_name;NOT NULL" json:"nick_name"`     // 昵称
	StudentID string `gorm:"column:student_id;unique" json:"student_id"`     // 学号
	Password  string `gorm:"column:password;NUT NULL" json:"-"`              // 密码
	Email     string `gorm:"column:email;NOT NULL" json:"email"`             // 邮箱
	Avatar    string `gorm:"column:avatar" json:"avatar"`                    // 头像
	Grade     int    `gorm:"column:grade;NOT NULL" json:"grade"`             // 所属年级
	Status    int    `gorm:"column:status;default:1;NOT NULL" json:"status"` // 1: 可用, 2: 封禁
}

func (User) TableName() string {
	return "users"
}
