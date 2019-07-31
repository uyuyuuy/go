package old

type Coin struct {
	//币种信息
	Name uint64	`gorm:"column:name"`
	AssetName string	`gorm:"column:asset_name"`
	MinOut float64	`gorm:"column:maxin"`
	MaxOut	float64	`gorm:"column:maxout"`
	OutLimit	float64	`gorm:"column:out_limit"`
	OutRate	float64	`gorm:"column:rate_out"`
	OutStatus	int	`gorm:"column:out_status"`
	InStatus	int	`gorm:"column:in_status"`
	Logo string	`gorm:"column:logo"`
	NumberFloat	int	`gorm:"column:number_float"`
	CoinTransfer	int		`gorm:"column:coin_transfer"`	//资金划转状态
	Describe	string	`gorm:"column:describe"`

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
