package trades

import (
	types "github.com/veska-io/dydx-v4-indexer-client/src/types"
)

type Trade struct {
	Id              string `json:"id"`
	Side            string `json:"side"`
	Size            string `json:"size"`
	Price           string `json:"price"`
	Type            string `json:"type"`
	CreatedAt       string `json:"createdAt"`
	CreatedAtHeight string `json:"createdAtHeight"`
}

type TradesResponse struct {
	Trades []Trade           `json:"trades"`
	Errors []types.DYDXError `json:"errors"`
}
