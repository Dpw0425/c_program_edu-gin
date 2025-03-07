package model

import (
	"gorm.io/gorm"
	"time"
)

type Competition struct {
	gorm.Model
	Name       string    `gorm:"column:name;NOT NULL" json:"name"`             // 比赛名称
	Contestant string    `gorm:"column:contestant" json:"contestant"`          // 参赛选手
	OwnerID    uint      `gorm:"column:owner_id;NOT NULL" json:"owner_id"`     // 创办者 ID
	StartTime  time.Time `gorm:"column:start_time;NOT NULL" json:"start_time"` // 开始时间
	Deadline   time.Time `gorm:"column:deadline;NOT NULL" json:"deadline"`     // 截止时间
	Status     int       `gorm:"column:status;default:0" json:"status"`        // 比赛状态
	Category   int       `gorm:"column:category;NOT NULL" json:"category"`     // 比赛类别
	Permission int       `gorm:"column:permission;NOT NULL" json:"permission"` // 比赛权限
}

func (Competition) TableName() string {
	return "competitions"
}
