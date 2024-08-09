package markets

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"

	"github.com/veska-io/dydx-v4-indexer-client/src/config"
)

func APIRequest(limit uint8) (*PerpetualMarketsResponse, error) {
	var marketsData PerpetualMarketsResponse

	url, err := generateUrl(limit)
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
		return nil, err
	}

	err = json.Unmarshal(body, &marketsData)
	if json.Unmarshal(body, &marketsData) != nil {
		return nil, err
	}

	return &marketsData, nil
}

func generateUrl(limit uint8) (*url.URL, error) {
	cfg := config.MustNew()
	baseURL := cfg.Url + cfg.MarketsPath

	params := url.Values{}

	if limit > 0 {
		params.Add("limit", fmt.Sprint(limit))
	}

	u, err := url.Parse(baseURL)
	if err != nil {
		return nil, fmt.Errorf("failed to parse base URL: %w", err)
	}

	u.RawQuery = params.Encode()
	return u, nil
}
