package model

import "gorm.io/gorm"

type QuestionBank struct {
	gorm.Model
	Name         string `gorm:"column:name;NOT NULL" json:"name"`        // 题库名称
	Content      string `gorm:"column:content" json:"content"`           // 内容
	OpenGrade    string `gorm:"column:open_grade" json:"open_grade"`     // 开放年级
	Participants int    `gorm:"column:participants" json:"participants"` // 参与人数
	Completed    int    `gorm:"column:completed" json:"completed"`       // 完成人数
}

func (QuestionBank) TableName() string {
	return "question_banks"
}
