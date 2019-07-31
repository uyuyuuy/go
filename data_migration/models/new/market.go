package new

type Market struct {
	ID int `gorm:"column:F01"`
	CoinID int `gorm:"column:F02"`
	Name string `gorm:"column:F03"`
	AssetName string `gorm:"column:F04"`
	Order int	`gorm:"column:F05"`	//排序值（倒序）
	IsDisplay string `gorm:"column:F06"`	//是否开启分区显示：S-是；F-否
	IsDelete	string	`gorm:"column:F07"`

}

func (Market) TableName() string {
	return "t6014"
}

