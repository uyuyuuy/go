package a

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model
	Code string
	Price uint
}

func (User) TableName() string {
	return "user"
}
