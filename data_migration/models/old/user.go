package old

type User struct {
	Code string
	Price uint
}

func (User) TableName() string {
	return "user"
}
