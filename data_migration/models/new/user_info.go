package new

import "database/sql"

type UserInfo struct {
	UserID	uint64	`gorm:"column:F01"`
	Name  string	`gorm:"column:F02"`
	CredentialsType string	`gorm:"column:F03"`
	ID	string	`gorm:"column:F04"`
	CountryID	int	`gorm:"column:F05"`
	ReferrerTag	sql.NullString	`gorm:"column:F06"`
	MyselfTag	sql.NullString	`gorm:"column:F07"`
	TradePassword	string	`gorm:"column:F08"`
	RegisterTime	string	`gorm:"column:F09"`
	LastLoginTime	string	`gorm:"column:F10"`
	VipStatus string `gorm:"column:F13"`
	VipTime string `gorm:"column:F14"`

}


func (UserInfo) TableName() string {
	return "t6011"
}

