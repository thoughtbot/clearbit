package clearbit_test

import (
	"encoding/json"
	"testing"

	"github.com/thoughtbot/clearbit"
)

func TestEnrichmentPersonResponse(t *testing.T) {
	var person clearbit.Person

	if err := unmarshalFixture("enrichment_person_response", &person); err != nil {
		t.Fatal("Failed to unmarshal:", err)
	}

	if person.ID == "" {
		t.Fatal("Expected person to be present")
	}
}

func TestEnrichmentCompanyResponse(t *testing.T) {
	var company clearbit.Company

	if err := unmarshalFixture("enrichment_company_response", &company); err != nil {
		t.Fatal("Failed to unmarshal:", err)
	}

	if company.ID == "" {
		t.Fatal("Expected company to be present")
	}
}

func TestErrorResponse(t *testing.T) {
	var error clearbit.ErrorResponse

	if err := unmarshalFixture("error_response", &error); err != nil {
		t.Fatal("Failed to unmarshal:", err)
	}

	if error.Type != "params_invalid" {
		t.Fatal("Expected error to be unmarshaled")
	}
}

func TestTwitterID(t *testing.T) {
	var id clearbit.TwitterID

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
