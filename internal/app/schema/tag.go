package schema

type TagItem struct {
	Name  string `gorm:"column:name;" json:"name"`
	Count int    `gorm:"column:count;" json:"count"`
}
