package clearbit

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

// These are the valid services and resources of the Clearbit API.
const (
	ProspectURL               = "https://prospector.clearbit.com/v1/people/search"
	EnrichCompanyStreamingURL = "https://company-stream.clearbit.com/v2/companies/find"
	EnrichPersonStreamingURL  = "https://person-stream.clearbit.com/v2/people/find"
)

// Client provides access to the Clearbit API.
type Client struct {
	apiKey string

	httpClient *http.Client
}

// NewClient initializes a Clearbit API client with the provided apiKey.
//
// If httpClient is nil, http.DefaultClient will be used.
func NewClient(apiKey string, httpClient *http.Client) *Client {
	if httpClient == nil {
		httpClient = http.DefaultClient
	}

	return &Client{
		apiKey:     apiKey,
		httpClient: httpClient,
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

// EnrichPerson finds a person by their email address
// and returns detailed information about them.
func (c *Client) EnrichPerson(email string) (*Person, error) {
	params := url.Values{
		"email": []string{email},
	}

	var person *Person
	err := c.get(EnrichPersonStreamingURL, params, &person)

	return person, err
}

func (c *Client) get(endpoint string, params url.Values, v interface{}) error {
	req, err := http.NewRequest("GET", endpoint, nil)
	if err != nil {
		return err
	}

	req.SetBasicAuth(c.apiKey, "")
	req.URL.RawQuery = params.Encode()

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	if resp.StatusCode != 200 {
		return fmt.Errorf("clearbit HTTP error %d: %s", resp.StatusCode, data)
	}

	return json.Unmarshal(data, v)
}
