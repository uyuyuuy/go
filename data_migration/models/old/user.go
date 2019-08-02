package old

type UserScan struct {
	UID uint64	`gorm:"column:uid"`
	Email string	`gorm:"column:email"`
	Mobile string	`gorm:"column:mo"`
	Password	string	`gorm:"column:pwd"`
	TradePassword	string	`gorm:"column:pwdtrade"`
	Prand	string	`gorm:"column:prand"`
	Created	int64	`gorm:"column:created"`
	CreateIP string	`gorm:"column:createip"`
	Source	string	`gorm:"column:source"`
	RegisterType	int		`gorm:"column:registertype"`
	FromUID	int64		`gorm:"column:from_uid"`
	Area	string	`gorm:"column:area"`
	GoogleKey	string	`gorm:"column:google_key"`
	AutonymID	int	`gorm:"column:autonym_id"`
	Realname	string	`gorm:"column:realname"`
	CardType	int	`gorm:"column:cardtype"`
	IDCard	string	`gorm:"column:idcard"`
	Status	int	`gorm:"column:status"`
	FrontFace	string	`gorm:"column:frontFace"`
	BackFace	string	`gorm:"column:backFace"`
	Handkeep	string	`gorm:"column:handkeep"`
	AutonymCreated	int64	`gorm:"column:autonym_created"`
	AutonymUpdated	int64	`gorm:"column:autonym_updated"`
	AdminID	int64	`gorm:"column:admin"`
	Content	string	`gorm:"column:content"`
	Country	string	`gorm:"column:country"`

}
