package markets

type Market struct {
	ClobPairId                string `json:"clobPairId"`
	Ticker                    string `json:"ticker"`
	Status                    string `json:"status"`
	LastPrice                 string `json:"lastPrice"`
	OraclePrice               string `json:"oraclePrice"`
	PriceChange24H            string `json:"priceChange24H"`
	Volume24H                 string `json:"volume24H"`
	Trades24H                 int64  `json:"trades24H"`
	NextFundingRate           string `json:"nextFundingRate"`
	InitialMarginFraction     string `json:"initialMarginFraction"`
	MaintenanceMarginFraction string `json:"maintenanceMarginFraction"`
	BasePositionNotional      string `json:"basePositionNotional"`
	OpenInterest              string `json:"openInterest"`
	AtomicResolution          int64  `json:"atomicResolution"`
	QuantumConversionExponent int64  `json:"quantumConversionExponent"`
	TickSize                  string `json:"tickSize"`
	StepSize                  string `json:"stepSize"`
	StepBaseQuantums          int64  `json:"stepBaseQuantums"`
	SubticksPerTick           int64  `json:"subticksPerTick"`
}

type PerpetualMarketsResponse struct {
	Markets map[string]Market `json:"markets"`
}
