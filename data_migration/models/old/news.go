package old

type News struct {
	ID int64
	Title	string	`gorm:"column:uid"`
	Content string
	Receive string
	Created	int64
	Expired int64
	Sort	int64
	Category	int64
	Updated	int64
	Source	string
	Click	int64
	LanguageCode	string

}

func (News) TableName() string {
	return "news"
}


