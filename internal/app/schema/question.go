package schema

type QuestionItem struct {
	ID          uint    `gorm:"column:id" json:"id"`
	Title       string  `gorm:"column:title" json:"title"`
	Tag         string  `gorm:"column:tag" json:"tag"`
	Degree      int     `gorm:"column:degree" json:"degree"`
	PassingRate float32 `gorm:"column:passing_rate" json:"passing_rate"`
	OwnerID     int64   `gorm:"column:owner_id" json:"owner_id"`
}
