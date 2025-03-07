package model

import "gorm.io/gorm"

type Team struct {
	gorm.Model
	Name             string `gorm:"column:name;NOT NULL" json:"name"`                            // 团队名称
	Manager          int64  `gorm:"colum:manager;NOT NULL" json:"manager"`                       // 队长 ID
	Member           string `gorm:"column:member;NOT NULL" json:"member"`                        // 团队成员
	CompetitionTimes int    `gorm:"column:competition_times;default:0" json:"competition_times"` // 参与比赛次数
}

func (Team) TableName() string {
	return "teams"
}
