package mojikiban

import (
	"context"
	"errors"
	"fmt"
	"regexp"
)

type GetResponse struct {
	Version    string       `json:"version"`
	Date       string       `json:"date"`
	MJCharInfo []MJCharInfo `json:"MJ文字情報"`
}

type GetOptions struct {
	MJMojizukeimei string
}

func (o *GetOptions) Valid() error {
	if !regexp.MustCompile(`MJ[0-9]{6}`).Match([]byte(o.MJMojizukeimei)) {
		return errors.New("MJMojiZukeiMei must be in MJ[0-9]{6} format")
	}
	return nil
}

func (c *Client) Get(ctx context.Context, o GetOptions) (*GetResponse, error) {
	if err := o.Valid(); err != nil {
		return nil, err
	}

	spath := fmt.Sprintf("/mji/%s.json", o.MJMojizukeimei)
	req, err := c.newRequest(ctx, "GET", spath, nil)
	if err != nil {
		return nil, err
	}

	res, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	var g GetResponse
	if err := decodeBody(res, &g); err != nil {
		return nil, err
	}

	return &g, err
}
