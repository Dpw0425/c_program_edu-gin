package model

import "gorm.io/gorm"

type UserQue struct {
	gorm.Model
	UserID        int64  `gorm:"column:user_id;NOT NULL" json:"user_id"`                    // 用户 ID
	QuestionID    uint   `gorm:"column:question_id;NOT NULL" json:"question_id"`            // 题目 ID
	CompetitionID uint   `gorm:"column:competition_id;NOT NULL" json:"competition_id"`      // 比赛 ID
	Status        string `gorm:"column:status;NOT NULL;default:'incomplete'" json:"status"` // 完成情况
}

func (UserQue) TableName() string {
	return "user_ques"
}
