package clearbit_test

import (
	"encoding/json"
	"io/ioutil"
	"path"
)

func unmarshalFixture(name string, v interface{}) error {
	fixtureFilename := path.Join("fixtures", name) + ".json"

	data, err := ioutil.ReadFile(fixtureFilename)
	if err != nil {
		return err
	}

	return json.Unmarshal(data, v)
}
