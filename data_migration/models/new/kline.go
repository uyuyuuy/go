package new

type Kline struct {
	ID	int64	`gorm:"column:F01"`
	OpenPrice float64 `gorm:"column:F02"`
	ClosePrice float64 `gorm:"column:F03"`
	HightPrice float64 `gorm:"column:F04"`
	LowPrice float64 `gorm:"column:F05"`
	TotalAmount	float64	`gorm:"column:F06"`
	KlineTime	int64	`gorm:"column:F07"`
	CreateTime	int64	`gorm:"column:F08"`
	TotalNumber	float64	`gorm:"column:F09"`
	JosnData	string	`gorm:"column:F10"`
	TimeType	int64	`gorm:"column:F11"`
	MarketID	int64	`gorm:"column:F12"`
}


