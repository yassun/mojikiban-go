package mojikiban

import (
	"context"
	"errors"
	"regexp"
)

type SearchResponse struct {
	Status  string       `json:"status"`
	Find    bool         `json:"find"`
	Results []MJCharInfo `json:"results"`
	Count   int          `json:"count"`
}

type SearchOptions struct {
	UCS string
}

func (o *SearchOptions) Valid() error {
	if !regexp.MustCompile(`[0-9A-Fa-f]{4,5}`).Match([]byte(o.UCS)) {
		return errors.New("UCS must be in [0-9A-Fa-f]{4,5} format")
	}
	return nil
}

func (c *Client) Search(ctx context.Context, o SearchOptions) (*SearchResponse, error) {
	if err := o.Valid(); err != nil {
		return nil, err
	}

	req, err := c.newRequest(ctx, "GET", "/mji/q", nil)
	if err != nil {
		return nil, err
	}

	p := req.URL.Query()
	p.Add("UCS", o.UCS)
	req.URL.RawQuery = p.Encode()

	res, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	var s SearchResponse
	if err := decodeBody(res, &s); err != nil {
		return nil, err
	}

	return &s, err
}
