package schema

type QuestionItem struct {
	ID          uint    `gorm:"column:id" json:"id"`
	Title       string  `gorm:"column:title" json:"title"`
	Tag         string  `gorm:"column:tag" json:"tag"`
	Degree      int     `gorm:"column:degree" json:"degree"`
	PassingRate float32 `gorm:"column:passing_rate" json:"passing_rate"`
	OwnerID     int64   `gorm:"column:owner_id" json:"owner_id"`
	Content     string  `gorm:"column:content" json:"content"`
	Answer      string  `gorm:"column:answer" json:"answer"`
	Status      int     `gorm:"column:status" json:"status"`
}

type TestDataItem struct {
	ID         uint   `gorm:"column:id" json:"id"`
	Input      string `gorm:"column:input" json:"input"`
	Output     string `gorm:"column:output" json:"output"`
	QuestionID uint   `gorm:"column:question_id" json:"question_id"`
}
