package candles_test

import (
	"testing"

	"github.com/veska-io/dydx-v4-indexer-client/src/candles"
)

func TestAPIRequest(t *testing.T) {
	ticker := "ETH-USD"
	resolution := "1HOUR"
	limit := uint8(100)

	fromISO := "2024-04-01T00:00:00.000Z"
	toISO := "2024-04-03T00:00:00.000Z"

	resp, err := candles.APIRequest(ticker, resolution, limit, fromISO, toISO)
	if err != nil {
		t.Fatalf("failed to get response: %v", err)
	}

	t.Logf("response: %+v", resp)
}
