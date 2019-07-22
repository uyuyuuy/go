package new

type Product struct {
	ID int64
	Code  string
	Price uint
}

func (Product) TableName() string {
	return "product"
}

