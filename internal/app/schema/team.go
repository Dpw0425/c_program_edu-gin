package schema

type TeamItem struct {
	ID               uint   `gorm:"column:id" json:"id"`
	Name             string `gorm:"column:name" json:"name"`
	Manager          string `gorm:"column:manager" json:"manager"`
	Member           string `gorm:"column:member" json:"member"`
	CompetitionTimes int    `gorm:"column:competition_times" json:"competition_times"`
}
