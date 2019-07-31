package new

type Coin struct {
	ID int64 `gorm:"column:F01"`
	Name uint64 `gorm:"column:F02"`
	AssetName string `gorm:"column:F03"`
	Logo string `gorm:"column:F04"`
	InStatus string	`gorm:"column:F05"`
	OutStatus string	`gorm:"column:F06"`
	MoveStatus	string	`gorm:"column:F07"`
	CreateTime string `gorm:"column:F08"`
	MinOut	float64	`gorm:"column:F09"`
	MaxOut	float64	`gorm:"column:F10"`
	MinOutFee	float64	`gorm:"column:F11"`	//最小提币手续费数量
	MaxOutFee	float64	`gorm:"column:F12"`	//最大提币手续费数量
	OutRate	float64	`gorm:"column:F13"`	//提币费率
	ContractUrl	string	`gorm:"column:F14"`	//合约地址
	Series	string	`gorm:"column:F15"`		//系列
	Description	string	`gorm:"column:F16"`
	OutLimit	float64	`gorm:"column:F17"`
	IsDelete	string	`gorm:"column:F18"`
	MaxMove	float64	`gorm:"column:F19"`	//最大划转数量
	OutFloatNumberLimit	int		`gorm:"column:F20"`	//提币小数位控制
	IsFreeFee	string	`gorm:"column:F21"`	//是否免手续费：S-是；F-否
	ChormUrl	string	`gorm:"column:F22"`	//区块链浏览器URL
	IsHaveMemo	string	`gorm:"column:F23"`	//充币memo：S-是；F-否
	MinIn	float64	`gorm:"column:F18"`	//充币最小数量

}

func (Coin) TableName() string {
	return "t6013"
}

