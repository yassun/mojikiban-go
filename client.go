package mojikiban

import (
	"context"
	"encoding/json"
	"io"
	"net"
	"net/http"
	"net/url"
	"path"
	"runtime"
	"time"
)

func newDefaultHTTPClient() *http.Client {
	return &http.Client{
		Transport: &http.Transport{
			Proxy: http.ProxyFromEnvironment,
			DialContext: (&net.Dialer{
				Timeout:   30 * time.Second,
				KeepAlive: 30 * time.Second,
			}).DialContext,
			MaxIdleConns:          10,
			IdleConnTimeout:       60 * time.Second,
			TLSHandshakeTimeout:   10 * time.Second,
			ExpectContinueTimeout: 1 * time.Second,
			MaxIdleConnsPerHost:   runtime.GOMAXPROCS(0) + 1,
		},
	}
}

var defaultHTTPClient = newDefaultHTTPClient()

type HTTPClient interface {
	Do(*http.Request) (*http.Response, error)
}

type Client struct {
	URL        *url.URL
	HTTPClient *http.Client
}

func NewClient() *Client {
	u, _ := url.Parse(ApiEndpoint)
	client := Client{
		URL:        u,
		HTTPClient: defaultHTTPClient,
	}
	return &client
}

func (c *Client) newRequest(ctx context.Context, method, spath string, body io.Reader) (*http.Request, error) {
	u := *c.URL
	u.Path = path.Join(c.URL.Path, spath)

	req, err := http.NewRequest(method, u.String(), body)
	if err != nil {
		return nil, err
	}

	req = req.WithContext(ctx)

	req.Header.Set("User-Agent", "mojikiban-go/"+Version)
	req.Header.Set("Content-Type", "application/json")

	return req, nil
}

func decodeBody(resp *http.Response, out interface{}) error {
	defer resp.Body.Close()
	decoder := json.NewDecoder(resp.Body)
	return decoder.Decode(out)
}
