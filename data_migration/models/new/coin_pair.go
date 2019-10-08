package new

type CoinPair struct {
	ID int64 `gorm:"column:F01"`
	MarketID int `gorm:"column:F02"`
	BuyCoinID int `gorm:"column:F03"`
	SaleCoinID int `gorm:"column:F04"`
	PricePrecision	int	`gorm:"column:F05"`	//交易市场-交易金额显示精确度（小数位数）

	BuyRate	float64	`gorm:"column:F06"`
	SaleRate	float64	`gorm:"column:F07"`

	BuyMin	float64	`gorm:"column:F08"`
	SaleMin	float64	`gorm:"column:F09"`

	BuyMax	float64	`gorm:"column:F10"`
	SaleMax	float64	`gorm:"column:F11"`

	Order	int	`gorm:"column:F12"`
	CreateTime	string	`gorm:"column:F13"`

	IsDisplay	string	`gorm:"column:F14"`
	IsOpen	string	`gorm:"column:F15"`
	IsDelete	string	`gorm:"column:F16;DEFAULT:F"`
	ApiTradeMin	float64	`gorm:"column:F17"`	//api交易最小数量
	ApiTradeMax	float64	`gorm:"column:F18"`	//api交易最大数量
	NumberPrecision	int	`gorm:"column:F19"`	//交易市场-交易数量显示精确度（小数位数）',

}

func (CoinPair) TableName() string {
	return "t6015"
}

