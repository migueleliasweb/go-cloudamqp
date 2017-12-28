package cloudamqp

import (
	"net/http"
	"net/url"

	"path"

	"github.com/dghubble/sling"
)

const (
	DefaultBaseURL = "https://customer.cloudamqp.com/api/"
)

type Client struct {
	sling     *sling.Sling
	Instances *InstanceService
}

// NewClient returns a new Sentry API client.
// If a nil httpClient is given, the http.DefaultClient will be used.
// If a nil baseURL is given, the DefaultBaseURL will be used.
func NewClient(httpClient *http.Client, baseURL *url.URL, token string) *Client {
	if httpClient == nil {
		httpClient = http.DefaultClient
	}

	if baseURL == nil {
		baseURL, _ = url.Parse(DefaultBaseURL)
	}
	baseURL.Path = path.Join(baseURL.Path) + "/"

	base := sling.New().Base(baseURL.String()).Client(httpClient)

	if token != "" {
		base.Add("Authorization", "Bearer "+token)
	}

	c := &Client{
		sling:     base,
		Instances: newInstanceService(base.New()),
	}
	return c
}
