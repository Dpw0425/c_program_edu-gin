package model

import "gorm.io/gorm"

type Admin struct {
	gorm.Model
	UserName   string `gorm:"column:user_name;NOT NULL" json:"user_name"`          // 用户名
	TeacherID  uint   `gorm:"column:teacher_id;NOT NULL;unique" json:"teacher_id"` // 教师工号
	Password   string `gorm:"column:password;NOT NULL" json:"-"`                   // 密码
	Permission string `gorm:"column:permission" json:"permission"`                 // 权限
	Status     int    `gorm:"column:status;default:1" json:"status"`               // 帐号状态
}

func (Admin) TableName() string {
	return "admins"
}
