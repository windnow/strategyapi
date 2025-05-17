package strategyapi

const (
	DirectionBuy  = "BUY"
	DirectionSell = "SELL"
	DirectionNone = "NONE"
)

type Candle struct {
	Symbol    string
	Interval  string
	OpenTime  int64
	CloseTime int64
	Open      float64
	High      float64
	Low       float64
	Close     float64
	Volume    float64
}

type TradeSignal struct {
	Symbol     string
	Direction  string // "BUY", "SELL", "NONE"
	Price      float64
	StopLoss   float64
	TakeProfit float64
	TimeStamp  int64
	StrategyID string
}

type StrategyConfig struct {
	StrategyName   string
	ParametersJSON string
	TimeFrame      string
}

type StrategyMeta struct {
	Name                 string
	Version              string
	RequiredHistory      int
	ParametersSchemaJSON string
}
