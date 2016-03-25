package clearbit

import (
	"encoding/json"
	"io/ioutil"
	"path"
	"testing"
)

func readFixture(t *testing.T, name string) []byte {
	fixtureFilename := path.Join("fixtures", name) + ".json"

	data, err := ioutil.ReadFile(fixtureFilename)
	if err != nil {
		t.Fatalf("Unable to read fixture %q: %s", name, err)
	}
	return data
}

func unmarshalFixture(t *testing.T, name string, v interface{}) {
	err := json.Unmarshal(readFixture(t, name), v)
	if err != nil {
		t.Fatal("Failed to unmarshal:", err)
	}
}
