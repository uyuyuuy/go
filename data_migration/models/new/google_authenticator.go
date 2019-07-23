package new

type GoogleAuthenticator struct {
	UserID int64	`gorm:"column:F01"`
	SecretKey  string	`gorm:"column:F02"`
	Url string	`gorm:"column:F03"`
}

func (GoogleAuthenticator) TableName() string {
	return "t6057"
}

