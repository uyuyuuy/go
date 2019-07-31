package new

type UserAsset struct {
	ID	int64	`gorm:"column:F01"`
	UserID	uint64	`gorm:"column:F02"`
	CoinID	int	`gorm:"column:F03"`
	Over	float64	`gorm:"column:F04"`
	Lock	float64	`gorm:"column:F05"`
}

func (UserAsset) TableName() string {
	return "t6025"
}

