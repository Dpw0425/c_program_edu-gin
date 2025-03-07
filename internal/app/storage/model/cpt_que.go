package model

import "gorm.io/gorm"

type CptQue struct {
	gorm.Model
	CompetitionID uint `gorm:"column:competition_id;NOT NULL" json:"competition_id"` // 比赛 ID
	QuestionID    uint `gorm:"column:question_id;NOT NULL" json:"question_id"`       // 题目 ID
	CommitTimes   int  `gorm:"column:commit_times;default:0" json:"commit_times"`    // 提交次数
	Accepted      int  `gorm:"column:accepted;default:0" json:"accepted"`            // 通过人数
}

func (CptQue) TableName() string {
	return "cpt_ques"
}
