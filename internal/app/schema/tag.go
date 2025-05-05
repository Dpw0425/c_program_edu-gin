package schema

type TagItem struct {
	ID    uint   `gorm:"column:id" json:"id"`
	Name  string `gorm:"column:name;" json:"name"`
	Count int    `gorm:"column:count;" json:"count"`
}
