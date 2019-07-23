package new

type Autonym struct {
	UserID int64	`gorm:"column:F01"`
	Front  string	`gorm:"column:F02"`
	Reverse string	`gorm:"column:F03"`
	HandFront string	`gorm:"column:F04"`
	Remark string	`gorm:"column:F05"`
}

func (Autonym) TableName() string {
	return "t6011_1"
}

