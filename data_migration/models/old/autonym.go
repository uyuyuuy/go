package old

type Autonym struct {
	Code string
	Price uint
}

func (Autonym) TableName() string {
	return "autonym"
}
