package clearbit

import (
	"encoding/json"
	"testing"
)

func TestEnrichmentCombinedResponse(t *testing.T) {
	var combined CombinedResponse

	unmarshalFixture(t, "enrichment_combined_response", &combined)

	if combined.Person.ID == "" {
		t.Fatal("Expected person to be present")
	}

	if combined.Company.ID == "" {
		t.Fatal("Expected company to be present")
	}
}

func TestEnrichmentPersonResponse(t *testing.T) {
	var person Person

	unmarshalFixture(t, "enrichment_person_response", &person)

	if person.ID == "" {
		t.Fatal("Expected person to be present")
	}
}

func TestEnrichmentCompanyResponse(t *testing.T) {
	var company Company

	unmarshalFixture(t, "enrichment_company_response", &company)

	if company.ID == "" {
		t.Fatal("Expected company to be present")
	}
}

func TestErrorResponse(t *testing.T) {
	var error ErrorResponse

	unmarshalFixture(t, "error_response", &error)

	if error.Type != "params_invalid" {
		t.Fatal("Expected error to be unmarshaled")
	}
}

func TestTwitterID(t *testing.T) {
	var id TwitterID

	data := []byte(`123`)
	if err := json.Unmarshal(data, &id); err != nil {
		t.Fatalf("Failed to unmarshal %q into TwitterID: %s", data, err)
	}

	if id != "123" {
		t.Fatalf("Unmarshaled id = %v, want %v", id, "123")
	}

	data = []byte(`"321"`)
	if err := json.Unmarshal(data, &id); err != nil {
		t.Fatalf("Failed to unmarshal %q into TwitterID: %s", data, err)
	}

	if id != "321" {
		t.Fatalf("Unmarshaled id = %v, want %v", id, "321")
	}
}
