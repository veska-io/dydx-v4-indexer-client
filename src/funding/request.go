package funding

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"

	"github.com/veska-io/dydx-v4-indexer-client/src/config"
)

func APIRequest(
	ticker string, limit uint8, effectiveBeforeOrAtHeight uint64, effectiveBeforeOrAt string,
) (*HistoricalFundingResponse, error) {
	var fundingResponse HistoricalFundingResponse

	url, err := generateUrl(ticker, limit, effectiveBeforeOrAtHeight, effectiveBeforeOrAt)
	if err != nil {
		return nil, fmt.Errorf("failed to generate URL: %w", err)
	}

	resp, err := http.Get(url.String())
	if err != nil {
		return nil, fmt.Errorf("failed to get response from external API: %w", err)
	}

	body, err := io.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	err = json.Unmarshal(body, &fundingResponse)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal response body: %w", err)
	}

	return &fundingResponse, nil
}

func generateUrl(
	ticker string, limit uint8, effectiveBeforeOrAtHeight uint64, effectiveBeforeOrAt string,
) (*url.URL, error) {
	cfg := config.MustNew()
	baseURL := cfg.Url + cfg.FundingPath + "/" + ticker

	params := url.Values{}

	if limit > 0 {
		params.Add("limit", fmt.Sprint(limit))
	}

	if effectiveBeforeOrAtHeight > 0 {
		params.Add("effectiveBeforeOrAtHeight", fmt.Sprint(effectiveBeforeOrAtHeight))
	}

	if effectiveBeforeOrAt != "" {
		params.Add("effectiveBeforeOrAt", effectiveBeforeOrAt)
	}

	u, err := url.Parse(baseURL)
	if err != nil {
		return nil, fmt.Errorf("failed to parse base URL: %w", err)
	}

	u.RawQuery = params.Encode()
	return u, nil
}
