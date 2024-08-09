package client_test

import (
	"testing"
	"time"

	client "github.com/veska-io/dydx-v4-indexer-client/src"
)

func TestCandles(t *testing.T) {
	ticker := "ETH-USD"
	resolution := "1HOUR"
	limit := uint8(100)

	now := time.Now().UTC()
	dateTo := time.Date(now.Year(), now.Month(), now.Day(), now.Hour(), 0, 0, 0, time.UTC)
	dateFrom := dateTo.Add(-5 * 24 * time.Hour)

	c := client.New()

	resp, err := c.GetCandles(ticker, resolution, limit, dateFrom, dateTo)
	if err != nil {
		t.Fatalf("failed to get response: %v", err)
	}

	t.Logf("response: %+v", resp)
}

func TestFunding(t *testing.T) {
	ticker := "ETH-USD"
	limit := uint8(100)

	now := time.Now().UTC()
	dateTo := time.Date(now.Year(), now.Month(), now.Day(), now.Hour(), 0, 0, 0, time.UTC)

	c := client.New()

	resp, err := c.GetHistoricalFunding(ticker, limit, 0, dateTo)
	if err != nil {
		t.Fatalf("failed to get response: %v", err)
	}

	t.Logf("response: %+v", resp)
}

func TestTrades(t *testing.T) {
	ticker := "ETH-USD"
	limit := uint8(100)

	now := time.Now().UTC()
	dateTo := time.Date(now.Year(), now.Month(), now.Day(), now.Hour(), 0, 0, 0, time.UTC)

	c := client.New()

	resp, err := c.GetTrades(ticker, limit, 0, dateTo, 0)
	if err != nil {
		t.Fatalf("failed to get response: %v", err)
	}

	t.Logf("response: %+v", resp)
}
