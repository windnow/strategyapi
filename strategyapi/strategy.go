package strategyapi

type Strategy interface {
	Init(config StrategyConfig) error
	Feed(candle Candle) (*TradeSignal, error)
	Meta() StrategyMeta
	Config() StrategyConfig
}
