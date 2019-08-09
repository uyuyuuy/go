package old

type Openapi struct {
	ID int64
	UID int64	`gorm:"column:uid"`
	AccessKey string
	SecretKey string
	Ip string
	Status int64
}

func (Openapi) TableName() string {
	return "openapi"
}


