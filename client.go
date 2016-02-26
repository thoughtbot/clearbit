package clearbit

import (
	"net/http"
	"net/url"
)

// These are the valid services and resources of the Clearbit API.
const (
	ProspectorPersonSearchURL = "https://prospector.clearbit.com/v1/people/search"
	StreamingCompanySearchURL = "https://company-stream.clearbit.com/v2/companies/find"
	StreamingPersonSearchURL  = "https://person-stream.clearbit.com/v2/people/find"
)

// Client provides access to the Clearbit API.
type Client struct {
	apiKey string

	httpClient *http.Client
}

// NewClient initializes a Clearbit API client with the provided apiKey.
func NewClient(apiKey string) *Client {
	return &Client{
		apiKey:     apiKey,
		httpClient: http.DefaultClient,
	}
}

// Get requests endpoint with the provided params
// and the client's API key.
//
// The caller must close the returned response's Body.
func (c *Client) Get(endpoint string, params url.Values) (*http.Response, error) {
	req, err := http.NewRequest("GET", endpoint, nil)
	if err != nil {
		return nil, err
	}

	req.SetBasicAuth(c.apiKey, "")
	req.URL.RawQuery = params.Encode()

	return c.httpClient.Do(req)
}
