package model

import (
	"gorm.io/gorm"
)

type Question struct {
	gorm.Model
	Title       string  `gorm:"column:title;NOT NULL" json:"title"`       // 标题
	Content     string  `gorm:"column:content;NOT NULL" json:"content"`   // 内容
	Tag         string  `gorm:"column:tag;NOT NULL" json:"tag"`           // 标签
	Degree      int     `gorm:"column:degree;NOT NULL" json:"degree"`     // 难度
	OwnerID     int64   `gorm:"column:owner_id;NOT NULL" json:"owner_id"` // 发布人 ID
	PassingRate float32 `gorm:"column:passing_rate" json:"passing_rate"`  // 正确率
}

func (Question) TableName() string {
	return "questions"
}