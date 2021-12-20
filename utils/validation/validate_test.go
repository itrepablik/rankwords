package validation

import (
	"io/ioutil"
	"testing"
)

func TestValidateCreateUser(t *testing.T) {
	data := loadFile(t, "../../testData/contents.json")
	out, err := ValidateJSONPayload(string(data), "contents.json")
	if err != nil {
		t.Fatal(err)
	}
	if out != nil {
		t.Errorf("there was validation errors: %v", out)
	}
}

func loadFile(t *testing.T, path string) []byte {
	t.Helper()
	contents, err := ioutil.ReadFile(path)
	if err != nil {
		t.Fatal(err)
	}
	return contents
}
