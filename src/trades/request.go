package trades

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"

	"github.com/veska-io/dydx-v4-indexer-client/src/config"
)

/*
	ratelimit-remaining: 99
	ratelimit-reset: 1723115940868
*/

func APIRequest(
	ticker string,
	limit uint8,
	createdBeforeOrAtHeight uint64,
	createdBeforeOrAt string,
	page uint8,
) (*TradesResponse, error) {
	var candlesResponse TradesResponse

	url, err := generateUrl(ticker, limit, createdBeforeOrAtHeight, createdBeforeOrAt, page)
	if err != nil {
		return nil, fmt.Errorf("failed to generate URL: %w", err)
	}

	fmt.Println(url.String())
	resp, err := http.Get(url.String())
	if err != nil {
		return nil, fmt.Errorf("failed to get response from external API: %w", err)
	}

	body, err := io.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	err = json.Unmarshal(body, &candlesResponse)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal response body: %w", err)
	}

	if len(candlesResponse.Errors) > 0 {
		return nil, fmt.Errorf("external API returned errors: %s",
			candlesResponse.Errors[0].Msg)
	}

	return &candlesResponse, nil
}

func generateUrl(
	ticker string,
	limit uint8,
	createdBeforeOrAtHeight uint64,
	createdBeforeOrAt string,
	page uint8,
) (*url.URL, error) {
	cfg := config.MustNew()
	baseURL := cfg.Url + cfg.TradesPath + "/" + ticker

	params := url.Values{}

	if limit > 0 {
		params.Add("limit", fmt.Sprint(limit))
	}

	if createdBeforeOrAtHeight > 0 {
		params.Add("createdBeforeOrAtHeight", fmt.Sprint(createdBeforeOrAtHeight))
	}

	if createdBeforeOrAt != "" {
		params.Add("createdBeforeOrAt", createdBeforeOrAt)
	}

	if page > 0 {
		params.Add("page", fmt.Sprint(page))
	}

	u, err := url.Parse(baseURL)
	if err != nil {
		return nil, fmt.Errorf("failed to parse base URL: %w", err)
	}

	u.RawQuery = params.Encode()
	return u, nil
}
