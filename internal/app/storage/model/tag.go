package model

import "gorm.io/gorm"

type Tag struct {
	gorm.Model
	Name string `gorm:"column:name;NOT NULL" json:"name"` // 名称
}

func (Tag) TableName() string {
	return "tags"
}
