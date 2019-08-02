package old

type CoinPair struct {

	//交易对信息
	CoinFrom	string		`gorm:"column:coin_from"`
	CoinTo	string	`gorm:"column:coin_to"`
	RateSale	float64	`gorm:"column:rate"`
	RateBuy	float64	`gorm:"column:rate_buy"`
	MinTrade	float64	`gorm:"column:min_trade"`
	MaxTrade	float64	`gorm:"column:max_trade"`
	Status	int	`gorm:"column:status"`
	PairPriceFloat	int	`gorm:"column:price_float"`
	PairNumberFloat	int	`gorm:"column:number_float"`
	OrderBy	int	`gorm:"column:order_by"`

}

func (CoinPair) TableName() string {
	return "coin_pair"
}
