package funding

import (
	types "github.com/veska-io/dydx-v4-indexer-client/src/types"
)

type HistoricalFunding struct {
	Ticker            string `json:"ticker"`
	Rate              string `json:"rate"`
	Price             string `json:"price"`
	EffectiveAt       string `json:"effectiveAt"`
	EffectiveAtHeight string `json:"effectiveAtHeight"`
}

type HistoricalFundingResponse struct {
	HistoricalFunding []HistoricalFunding `json:"historicalFunding"`
	Errors            []types.DYDXError   `json:"errors"`
}
