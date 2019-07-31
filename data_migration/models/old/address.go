package old

type Address struct {
	ID int64
	UID uint64
	Address string
	Coin string
	Status int8
	Created int64
	Updated int64
	Secret string
	PublicKey string `gorm:"column:publicKey"`
	Label string
}

func (Address) TableName() string {
	return "address"
}


