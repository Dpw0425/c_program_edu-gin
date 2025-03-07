package model

import "gorm.io/gorm"

type Entry struct {
	gorm.Model
	UserID        int64 `gorm:"column:user_id;NOT NULL" json:"user_id"`               // 用户 ID
	CompetitionID uint  `gorm:"column:competition_id;NOT NULL" json:"competition_id"` // 比赛 ID
	Status        int   `gorm:"column:status;default:0" json:"status"`                // 审核状态
}

func (Entry) TableName() string {
	return "entries"
}
