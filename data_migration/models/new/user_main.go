package new

import "database/sql"

type UserMain struct {
	UID uint64	`gorm:"column:F01"`
	AccountName	sql.NullString	`gorm:"column:F02"`
	Password	string	`gorm:"column:F03"`
	Mobile	sql.NullString	`gorm:"column:F04"`
	Email	sql.NullString	`gorm:"column:F05"`
	IsLock	string	`gorm:"column:F06;DEFAULT:F"`
	RegisterFrom	string	`gorm:"column:F07;DEFAULT:Web"`
	LockRemark	string	`gorm:"column:F08"`
	Secretkey	string	`gorm:"column:F09"`
	AvatarUrl string	`gorm:"column:F10"`
	Nick	string	`gorm:"column:F11"`
	AreaCode	string	`gorm:"column:F12"`
}

func (UserMain) TableName() string {
	return "t6010"
}

