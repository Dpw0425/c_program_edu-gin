package model

import "gorm.io/gorm"

type BankQue struct {
	gorm.Model
	BankID      uint `gorm:"column:bank_id;NOT NULL" json:"bank_id"`            // 题库 ID
	QuestionID  uint `gorm:"column:question_id;NOT NULL" json:"question_id"`    // 题目 ID
	CommitTimes int  `gorm:"column:commit_times;default:0" json:"commit_times"` // 提交次数
	Accepted    int  `gorm:"column:accepted;default:0" json:"accepted"`         // 通过人数
}

func (BankQue) TableName() string {
	return "bank_ques"
}
