package candles

import (
	types "github.com/veska-io/dydx-v4-indexer-client/src/types"
)

type Candle struct {
	StartedAt            string `json:"startedAt"`
	Ticker               string `json:"ticker"`
	Resolution           string `json:"resolution"`
	Low                  string `json:"low"`
	High                 string `json:"high"`
	Open                 string `json:"open"`
	Close                string `json:"close"`
	BaseTokenVolume      string `json:"baseTokenVolume"`
	UsdVolume            string `json:"usdVolume"`
	Trades               int    `json:"trades"`
	StartingOpenInterest string `json:"startingOpenInterest"`
}

type CandlesResponse struct {
	Candles []Candle          `json:"candles"`
	Errors  []types.DYDXError `json:"errors"`
}
