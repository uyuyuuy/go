package new

import "time"

type UserInfo struct {
	UserID	int64	`gorm:"column:F01"`
	Name  string	`gorm:"column:F02"`
	CredentialsType string	`gorm:"column:F03"`
	ID	string	`gorm:"column:F04"`
	CountryID	int	`gorm:"column:F05"`
	ReferrerTag	string	`gorm:"column:F06"`
	MyselfTag	string	`gorm:"column:F07"`
	TradePassword	string	`gorm:"column:F08"`
	RegisterTime	time.Time	`gorm:"column:F09"`
	LastLoginTime	time.Time	`gorm:"column:F10"`
	VipStatus string `gorm:"column:F03"`
	VipTime time.Time `gorm:"column:F14"`

}


func (UserInfo) TableName() string {
	return "t6011"
}

