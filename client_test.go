package clearbit

import (
	"net/http"
	"reflect"
	"strings"
	"testing"

	"github.com/jarcoal/httpmock"
)

func requestRecorder(out **http.Request, next httpmock.Responder) httpmock.Responder {
	return func(req *http.Request) (*http.Response, error) {
		*out = req
		return next(req)
	}
}

func TestClientEnrichPerson(t *testing.T) {
	var (
		email  = "user@example.com"
		apiKey = "clearbit-key"

		request *http.Request
	)

	transport := httpmock.NewMockTransport()
	transport.RegisterResponder(
		"GET",
		EnrichPersonStreamingURL,
		requestRecorder(
			&request,
			httpmock.NewBytesResponder(
				200,
				readFixture(t, "enrichment_person_response"),
			),
		),
	)

	client := NewClient(apiKey, &http.Client{Transport: transport})

	person, err := client.EnrichPerson(email)
	if err != nil {
		t.Fatal("EnrichPerson failed:", err)
	}

	if person.ID == "" {
		t.Fatal("Expected person to be present")
	}

	expectRequestWithAPIKey(t, request, apiKey)

	requestedEmail := request.URL.Query().Get("email")
	if requestedEmail != email {
		t.Errorf("email param = %q, want %q", requestedEmail, email)
	}
}

func TestClientEnrichPersonTransportError(t *testing.T) {
	var (
		client = NewClient("key", &http.Client{
			Transport: httpmock.NewMockTransport(),
		})
	)

	_, err := client.EnrichPerson("email")
	if err == nil {
		t.Fatal("EnrichPerson succeeded, should have failed")
	}
}

func TestClientEnrichPersonClearbitError(t *testing.T) {
	transport := httpmock.NewMockTransport()
	transport.RegisterResponder(
		"GET",
		EnrichPersonStreamingURL,
		httpmock.NewStringResponder(
			404,
			`{"error": {"type": "unknown_record", "message": "Unknown email address"}}`,
		),
	)

	client := NewClient("key", &http.Client{
		Transport: transport,
	})

	_, err := client.EnrichPerson("email")
	if err == nil {
		t.Fatal("Expected 404 to be an error")
	}

	errorMessage := err.Error()

	if !strings.Contains(errorMessage, "404") {
		t.Errorf(
			"Error message = %q, doesn't contain %q",
			errorMessage,
			"404",
		)
	}

	if !strings.Contains(errorMessage, "unknown_record") {
		t.Errorf(
			"Error message = %q, doesn't contain %q",
			errorMessage,
			"unknown_record",
		)
	}

	if !strings.Contains(errorMessage, "Unknown email address") {
		t.Errorf(
			"Error message = %q, doesn't contain %q",
			errorMessage,
			"Unknown email address",
		)
	}
}

func TestClientEnrichCompany(t *testing.T) {
	var (
		domain = "example.com"
		apiKey = "clearbit-key"

		request *http.Request

		transport = httpmock.NewMockTransport()
		client    = NewClient(apiKey, &http.Client{Transport: transport})
	)

	transport.RegisterResponder(
		"GET",
		EnrichCompanyStreamingURL,
		requestRecorder(
			&request,
			httpmock.NewBytesResponder(
				200,
				readFixture(t, "enrichment_company_response"),
			),
		),
	)

	company, err := client.EnrichCompany(domain)
	if err != nil {
		t.Fatal("EnrichCompany failed:", err)
	}

	if company.ID == "" {
		t.Fatal("Expected company to be present")
	}

	expectRequestWithAPIKey(t, request, apiKey)

	requestedDomain := request.URL.Query().Get("domain")
	if requestedDomain != domain {
		t.Errorf("domain param = %q, want %q", requestedDomain, domain)
	}
}

func TestClientCombinedEnrich(t *testing.T) {
	var (
		email  = "user@example.com"
		apiKey = "clearbit-key"

		request *http.Request

		transport = httpmock.NewMockTransport()
		client    = NewClient(apiKey, &http.Client{Transport: transport})
	)

	transport.RegisterResponder(
		"GET",
		EnrichCombinedStreamingURL,
		requestRecorder(
			&request,
			httpmock.NewBytesResponder(
				200,
				readFixture(t, "enrichment_combined_response"),
			),
		),
	)

	combined, err := client.Enrich(email)
	if err != nil {
		t.Fatal("EnrichCombined failed:", err)
	}

	if combined.Company.ID == "" {
		t.Fatal("Expected company to be present in combined response")
	}

	expectRequestWithAPIKey(t, request, apiKey)

	requestedEmail := request.URL.Query().Get("email")
	if requestedEmail != email {
		t.Errorf("email param = %q, want %q", requestedEmail, email)
	}
}

func TestClientProspect(t *testing.T) {
	var (
		domain    = "example.com"
		name      = "Jane"
		role      = "marketing"
		seniority = "executive"
		titles    = []string{"VP", "CEO"}
		apiKey    = "clearbit-key"

		request *http.Request

		transport = httpmock.NewMockTransport()
		client    = NewClient(apiKey, &http.Client{Transport: transport})
	)

	transport.RegisterResponder(
		"GET",
		ProspectURL,
		requestRecorder(
			&request,
			httpmock.NewBytesResponder(
				200,
				readFixture(t, "prospect_response"),
			),
		),
	)

	prospects, err := client.Prospect(ProspectQuery{
		Domain:    domain,
		Name:      name,
		Role:      role,
		Seniority: seniority,
		Titles:    titles,
	})
	if err != nil {
		t.Fatal("Prospect failed:", err)
	}

	if len(prospects) == 0 {
		t.Fatal("Expected prospects to be present")
	}

	expectRequestWithAPIKey(t, request, apiKey)

	requestedDomain := request.URL.Query().Get("domain")
	if requestedDomain != domain {
		t.Errorf("domain param = %q, want %q", requestedDomain, domain)
	}

	requestedName := request.URL.Query().Get("name")
	if requestedName != name {
		t.Errorf("name param = %q, want %q", requestedName, name)
	}

	requestedRole := request.URL.Query().Get("role")
	if requestedRole != role {
		t.Errorf("role param = %q, want %q", requestedRole, role)
	}

	requestedSeniority := request.URL.Query().Get("seniority")
	if requestedSeniority != seniority {
		t.Errorf("seniority param = %q, want %q", requestedSeniority, seniority)
	}

	requestedTitles := request.URL.Query()["titles[]"]
	if !reflect.DeepEqual(requestedTitles, titles) {
		t.Errorf("titles param = %v, want %v", requestedTitles, titles)
	}
}

func expectRequestWithAPIKey(t *testing.T, r *http.Request, apiKey string) {
	if r == nil {
		t.Fatal("Request not sent")
	}

	username, _, _ := r.BasicAuth()
	if username != apiKey {
		t.Errorf("basic auth username = %q, want %q", username, apiKey)
	}
}
