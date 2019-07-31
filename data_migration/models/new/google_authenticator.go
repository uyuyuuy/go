package new

type GoogleAuthenticator struct {
	UserID uint64	`gorm:"column:F01;primary_key"`
	SecretKey  string	`gorm:"column:F02"`
	Url string	`gorm:"column:F03"`
}

func (GoogleAuthenticator) TableName() string {
	return "t6057"
}

