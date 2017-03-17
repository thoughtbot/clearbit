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
	ProspectURL                = "https://prospector.clearbit.com/v1/people/search"
	EnrichCombinedStreamingURL = "https://person.clearbit.com/v2/combined/find"
	EnrichCompanyStreamingURL  = "https://company-stream.clearbit.com/v2/companies/find"
	EnrichPersonStreamingURL   = "https://person-stream.clearbit.com/v2/people/find"
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

// Enrich finds a person by their email address
// and returns detailed information about both them
// as well as their company
func (c *Client) Enrich(email string) (*CombinedResponse, error) {
	var combined *CombinedResponse

	err := c.get(
		EnrichCombinedStreamingURL,
		url.Values{"email": []string{email}},
		&combined,
	)

	return combined, err
}

// EnrichPerson finds a person by their email address
// and returns detailed information about them.
func (c *Client) EnrichPerson(email string) (*Person, error) {
	var person *Person

	err := c.get(
		EnrichPersonStreamingURL,
		url.Values{"email": []string{email}},
		&person,
	)

	return person, err
}

// EnrichCompany finds a company by its domain
// and returns detailed information about it.
func (c *Client) EnrichCompany(domain string) (*Company, error) {
	var company *Company

	err := c.get(
		EnrichCompanyStreamingURL,
		url.Values{"domain": []string{domain}},
		&company,
	)

	return company, err
}

// ProspectQuery defines the available options for querying
// the Prospector API.
type ProspectQuery struct {
	// Company's domain name to look up (required).
	Domain string

	// Filters results by first or last name (case-insensitive).
	Name string

	// Filters results by job role (case-sensitive).
	Role string

	// Filters results by job seniority (case-sensitive).
	Seniority string

	// Filters results by one or more titles.
	Titles []string
}

// Prospect finds a company by its domain name and returns basic
// information about the people working there.
func (c *Client) Prospect(q ProspectQuery) ([]*Prospect, error) {
	var prospects []*Prospect

	err := c.get(
		ProspectURL,
		url.Values{
			"domain":    []string{q.Domain},
			"email":     []string{"true"},
			"name":      []string{q.Name},
			"role":      []string{q.Role},
			"seniority": []string{q.Seniority},
			"titles[]":  q.Titles,
		},
		&prospects,
	)

	return prospects, err
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
