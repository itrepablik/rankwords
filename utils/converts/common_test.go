package converts

import (
	"io/ioutil"
	"log"
	"strings"
	"testing"
)

func TestSchemaToBytes(t *testing.T) {
	schema := loadFile(t, "../../schemas/contents_data.json")
	contents := SchemaToBytes(string(schema))
	if len(strings.TrimSpace(contents)) == 0 {
		t.Error("SchemaToBytes failed")
	}
	log.Println(contents)
}

func loadFile(t *testing.T, path string) []byte {
	t.Helper()
	contents, err := ioutil.ReadFile(path)
	if err != nil {
		t.Fatal(err)
	}
	return contents
}
