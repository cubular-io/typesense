package typesense

import (
	"context"
	"github.com/cubular-io/typesense/tapi"
	"net/http"
	"time"
)

type Client struct {
	api *tapi.ClientWithResponses
	Cfg *Config
}

type Config struct {
	Url        string
	ApiKey     string
	timeout    time.Duration
	httpClient *http.Client
}

type Option func(*Client)

const defaultTimeout = time.Second * 10

// New Create a new Client and Pings for a Health Check the Server
func New(serverUrl string, apiKey string, options ...Option) (*Client, error) {
	c := &Client{
		api: nil,
		Cfg: &Config{
			Url:     serverUrl,
			ApiKey:  apiKey,
			timeout: defaultTimeout,
		},
	}

	for _, option := range options {
		option(c)
	}

	if c.Cfg.httpClient == nil {
		c.Cfg.httpClient = &http.Client{Timeout: c.Cfg.timeout}
	} else {
		c.Cfg.httpClient.Timeout = c.Cfg.timeout
	}
	var err error
	c.api, err = tapi.NewClientWithResponses(c.Cfg.Url, WithAPIKey(c.Cfg.ApiKey), tapi.WithHTTPClient(c.Cfg.httpClient))
	if err != nil {
		return nil, err
	}

	_, err = c.api.HealthWithResponse(context.Background())
	return c, err
}

func WithTimeout(timeout time.Duration) Option {
	return func(client *Client) {
		client.Cfg.timeout = timeout
	}
}

func (c *Client) Collection() *Collection {
	return newCollection(c.api)
}
