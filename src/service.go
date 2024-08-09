package client

import (
	"fmt"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/veska-io/dydx-v4-indexer-client/src/candles"
	"github.com/veska-io/dydx-v4-indexer-client/src/funding"
	"github.com/veska-io/dydx-v4-indexer-client/src/markets"
	"github.com/veska-io/dydx-v4-indexer-client/src/trades"
)

type Client struct{}

func New() *Client {
	return &Client{}
}

func (c *Client) GetCandles(
	ticker, resolution string, limit uint8, dateFrom, dateTo time.Time,
) (*candles.CandlesResponse, error) {
	p := struct {
		Ticker     string    `validate:"required"`
		Resolution string    `validate:"required,oneof=1MIN 5MINS 15MINS 30MINS 1HOUR 4HOURS 1DAY"`
		Limit      uint8     `validate:"min=1,max=100"`
		FromISO    time.Time `validate:"ltcsfield=ToISO"`
		ToISO      time.Time `validate:"gtcsfield=FromISO"`
	}{
		Ticker:     ticker,
		Resolution: resolution,
		Limit:      limit,
		FromISO:    dateFrom,
		ToISO:      dateTo,
	}

	v := validator.New(validator.WithRequiredStructEnabled())
	err := v.Struct(p)
	if err != nil {
		return nil, fmt.Errorf("validation failed: %w", err)
	}

	resp, err := candles.APIRequest(ticker, resolution, limit,
		dateFrom.Format("2006-01-02T15:04:05"), dateTo.Format("2006-01-02T15:04:05"))

	if err != nil {
		return nil, fmt.Errorf("failed to get candles: %w", err)
	}

	return resp, nil
}

func (c *Client) GetMarkets(limit uint8) (*markets.PerpetualMarketsResponse, error) {
	p := struct {
		Limit uint8 `validate:"min=1,max=100"`
	}{
		Limit: limit,
	}

	v := validator.New(validator.WithRequiredStructEnabled())
	err := v.Struct(p)
	if err != nil {
		return nil, fmt.Errorf("validation failed: %w", err)
	}

	resp, err := markets.APIRequest(limit)
	if err != nil {
		return nil, fmt.Errorf("failed to get markets: %w", err)
	}

	return resp, nil
}

func (c *Client) GetHistoricalFunding(
	ticker string, limit uint8, effectiveBeforeOrAtHeight uint64, effectiveBeforeOrAt time.Time,
) (*funding.HistoricalFundingResponse, error) {
	p := struct {
		Ticker string `validate:"required"`
		Limit  uint8  `validate:"min=1,max=100"`
		Height uint64
		Before time.Time
	}{
		Ticker: ticker,
		Limit:  limit,
		Height: effectiveBeforeOrAtHeight,
		Before: effectiveBeforeOrAt,
	}

	v := validator.New(validator.WithRequiredStructEnabled())
	err := v.Struct(p)
	if err != nil {
		return nil, fmt.Errorf("validation failed: %w", err)
	}

	resp, err := funding.APIRequest(ticker, limit, effectiveBeforeOrAtHeight,
		effectiveBeforeOrAt.Format("2006-01-02T15:04:05"))

	if err != nil {
		return nil, fmt.Errorf("failed to get funding: %w", err)
	}

	return resp, nil
}

func (c *Client) GetTrades(
	ticker string,
	limit uint8,
	createdBeforeOrAtHeight uint64,
	createdBeforeOrAt time.Time,
	page uint8,
) (*trades.TradesResponse, error) {
	p := struct {
		Ticker string `validate:"required"`
		Limit  uint8  `validate:"min=1,max=100"`
		Height uint64
		Before time.Time
		Page   uint8
	}{
		Ticker: ticker,
		Limit:  limit,
		Height: createdBeforeOrAtHeight,
		Before: createdBeforeOrAt,
		Page:   page,
	}

	v := validator.New(validator.WithRequiredStructEnabled())
	err := v.Struct(p)
	if err != nil {
		return nil, fmt.Errorf("validation failed: %w", err)
	}

	resp, err := trades.APIRequest(ticker, limit, createdBeforeOrAtHeight,
		createdBeforeOrAt.Format("2006-01-02T15:04:05"), page)

	if err != nil {
		return nil, fmt.Errorf("failed to get funding: %w", err)
	}

	return resp, nil
}
