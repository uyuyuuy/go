package new

type News struct {
	ID	int64	`gorm:"column:F01"`
	Pv int64 `gorm:"column:F02"`
	Sort int64 `gorm:"column:F03"`
	Title string `gorm:"column:F04"`
	Content string `gorm:"column:F05"`
	From	string	`gorm:"column:F06"`
	CreateTime	string	`gorm:"column:F07"`
	Type	string	`gorm:"column:F08;type:enum('Notice','Activity','Project','Information','No');default:No"`
	IsPublic	string	`gorm:"column:F09;default:S"`
	IsDelete	string	`gorm:"column:F10;default:F"`
	Lang	string	`gorm:"column:F11"`
}

func (News) TableName() string {
	return "t5012"
}

