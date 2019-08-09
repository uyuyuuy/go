package new

type Openapi struct {
	ID	int64	`gorm:"column:F01"`
	AccessKey string `gorm:"column:F02"`
	SecretKey string `gorm:"column:F03"`
	Ip string `gorm:"column:F04"`
	UID int64 `gorm:"column:F05"`
}

func (Openapi) TableName() string {
	return "t6054"
}

