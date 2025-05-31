package schema

import "time"

type CompetitionItem struct {
	ID         uint      `gorm:"column:id" json:"id"`
	Name       string    `gorm:"column:name" json:"name"`
	Contestant string    `gorm:"column:contestant" json:"contestant"`
	OwnerID    uint      `gorm:"column:owner_id" json:"owner_id"`
	StartTime  time.Time `gorm:"column:start_time" json:"start_time"`
	Deadline   time.Time `gorm:"column:deadline" json:"deadline"`
	Status     int       `gorm:"column:status" json:"status"`
	Category   int       `gorm:"column:category" json:"category"`
	Permission int       `gorm:"column:permission" json:"permission"`
}

type WebCompetitionItem struct {
	ID         uint      `gorm:"column:id" json:"id"`
	Name       string    `gorm:"column:name" json:"name"`
	Contestant string    `gorm:"column:contestant" json:"contestant"`
	OwnerID    uint      `gorm:"column:owner_id" json:"owner_id"`
	StartTime  time.Time `gorm:"column:start_time" json:"start_time"`
	Deadline   time.Time `gorm:"column:deadline" json:"deadline"`
	Status     int       `gorm:"column:status" json:"status"`
	Category   int       `gorm:"column:category" json:"category"`
	Permission int       `gorm:"column:permission" json:"permission"`
	Quantity   int       `gorm:"column:quantity" json:"quantity"`
}
