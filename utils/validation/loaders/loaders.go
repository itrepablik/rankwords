package loaders

import (
	"bytes"
	"encoding/json"
	"fmt"
	"rankwords/utils/validation/schemas"
	"runtime"
	"strings"

	"github.com/xeipuuv/gojsonreference"
	"github.com/xeipuuv/gojsonschema"
)

const (
	SchemaSourceDev  = "file:///C:/itrepablik/rankwords/schemas/" // for Windows local dev
	SchemaSourceProd = "file://itrepablik/"                       // for e.g in GCP Cloud Run Linux prod
)

// InternalLoaderFactory is a JSONLoaderFactory that creates InternalReferenceLoaders
type InternalLoaderFactory struct {
}

// New is a JSONLoaderFactory that creates InternalReferenceLoaders
func (f InternalLoaderFactory) New(source string) gojsonschema.JSONLoader {
	return &internalReferenceLoader{
		source: source,
	}
}

// internalReferenceLoader is a JSONLoader that loads JSON from a file
type internalReferenceLoader struct {
	source string
}

// JsonSource returns the source of the JSONLoader
func (i internalReferenceLoader) JsonSource() interface{} {
	// Check if OS is Windows for local dev
	schemaSrc := SchemaSourceProd
	if runtime.GOOS == "windows" {
		schemaSrc = SchemaSourceDev
	}

	if strings.HasPrefix(i.source, schemaSrc) {
		return i.source
	}
	return schemaSrc + i.source
}

// LoadJSON is a JSONLoader that loads JSON from a file
func (i internalReferenceLoader) LoadJSON() (interface{}, error) {
	// Check if the source is a valid JSON file
	reference, err := gojsonreference.NewJsonReference(i.JsonSource().(string))
	if err != nil {
		return nil, err
	}

	// Check if OS is Windows for local dev
	schemaSrc := SchemaSourceProd
	if runtime.GOOS == "windows" {
		schemaSrc = SchemaSourceDev
	}

	// Check if the file exists
	name := strings.ToLower(strings.Replace(reference.String(), schemaSrc, "", 1))

	schemaBytes, exist := schemas.SchemaMap[name]
	if !exist {
		return nil, fmt.Errorf("unable to find schema for %s", name)
	}
	return i.initFromBytes(schemaBytes)
}

// initFromBytes loads JSON from a byte array
func (i internalReferenceLoader) initFromBytes(data []byte) (interface{}, error) {
	var document interface{}
	r := bytes.NewReader(data)
	decoder := json.NewDecoder(r)
	decoder.UseNumber()

	err := decoder.Decode(&document)
	if err != nil {
		return nil, err
	}
	return document, nil
}

// JsonReference is a JSONLoader that loads JSON from a reference
func (i internalReferenceLoader) JsonReference() (gojsonreference.JsonReference, error) {
	return gojsonreference.NewJsonReference(i.JsonSource().(string))
}

// LoaderFactory is a JSONLoaderFactory that creates InternalReferenceLoaders
func (i internalReferenceLoader) LoaderFactory() gojsonschema.JSONLoaderFactory {
	return &InternalLoaderFactory{}
}
