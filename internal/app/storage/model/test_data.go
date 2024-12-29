package model

import "gorm.io/gorm"

type TestData struct {
	gorm.Model
	QuestionID uint   `gorm:"column:question_id;NOT NULL" json:"question_id"` // 题目 ID
	Input      string `gorm:"column:input" json:"input"`                      // 输入数据
	Output     string `gorm:"column:output;NOT NULL" json:"output"`           // 输出数据
}

func (TestData) TableName() string {
	return "test_data"
}
