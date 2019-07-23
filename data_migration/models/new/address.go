package new

import "time"

type Address struct {
	ID int64 `gorm:"column:F01"`
	UserID int64 `gorm:"column:F02"`
	CoinID int `gorm:"column:F03"`
	Address string `gorm:"column:F04"`
	CreateTime time.Time `gorm:"column:F01"`
	Memo int64 `gorm:"column:F08"`

}

func (Address) TableName() string {
	return "t6012_1"
}

