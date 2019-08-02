package old

type Coin struct {
	//币种信息
	Name string	`gorm:"column:name"`
	AssetName string	`gorm:"column:asset_name"`
	MinOut float64	`gorm:"column:minout"`
	MaxOut	float64	`gorm:"column:maxout"`
	OutLimit	float64	`gorm:"column:out_limit"`
	OutRate	float64	`gorm:"column:rate_out"`
	OutStatus	int	`gorm:"column:out_status"`
	InStatus	int	`gorm:"column:in_status"`
	Logo string	`gorm:"column:logo"`
	NumberFloat	int	`gorm:"column:number_float"`
	CoinTransfer	int		`gorm:"column:coin_transfer"`	//资金划转状态
	Describe	string	`gorm:"column:describe"`


}

func (Coin) TableName() string {
	return "coin"
}
