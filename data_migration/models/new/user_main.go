package new

type UserMain struct {
	ID int64	`gorm:"column:F01"`
	AccountName	string	`gorm:"column:F02"`
	Password	string	`gorm:"column:F03"`
	Mobile	string	`gorm:"column:F04"`
	Email	string	`gorm:"column:F05"`
	IsLock	string	`gorm:"column:F06"`
	RegisterFrom	string	`gorm:"column:F07"`
	LockRemark	string	`gorm:"column:F08"`
	Secretkey	string	`gorm:"column:F09"`
	AvatarUrl string	`gorm:"column:F10"`
	Nick	string	`gorm:"column:F11"`
	AreaCode	string	`gorm:"column:F12"`
}

func (UserMain) TableName() string {
	return "t6010"
}

