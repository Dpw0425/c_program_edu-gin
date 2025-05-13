package schema

type BankItem struct {
	ID           uint   `gorm:"column:id" json:"id"`
	Name         string `gorm:"column:name" json:"name"`
	Content      string `gorm:"column:content" json:"content"`
	OpenGrade    string `gorm:"column:open_grade" json:"open_grade"`
	Participants int    `gorm:"column:participants" json:"participants"`
	Completed    int    `gorm:"column:completed" json:"completed"`
	Count        int    `gorm:"column:count" json:"count"`
}
