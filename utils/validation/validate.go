package validation

import (
	"encoding/json"
	"errors"
	"fmt"
	"rankwords/utils/validation/loaders"
	"strings"

	"github.com/xeipuuv/gojsonschema"
)

// APIVersion is a map containing the API versions
var (
	APIVersion = map[string]string{
		"1": "1",
	}
)

// CustomValidationError is a struct that holds the custom validation error
type CustomValidationError struct {
	ErrorMessage  []string `json:"error_message"`
	ErrorResponse string   `json:"error_response"`
}

// PrepCustomValidationError is a function that prepares the custom validation error
func PrepCustomValidationError(errMsg CustomValidationError) string {
	// Check if the error message is empty
	if len(errMsg.ErrorMessage) == 0 {
		return ""
	}

	// Remove curly braces from the error message using strings.Replace
	modErrMsg := strings.Replace(strings.Join(errMsg.ErrorMessage, " "), "{", "", -1)

	// Add square brackets to the error message
	modErrMsg = fmt.Sprintf("[%s]", modErrMsg)

	return modErrMsg
}

// CustomInvalidError is a struct that holds the context and the validation result
type CustomInvalidError struct {
	gojsonschema.ResultErrorFields
}

// AddCustomInvalidError adds a custom error to the validation result
func AddCustomInvalidError(context *gojsonschema.JsonContext, details gojsonschema.ErrorDetails) *CustomInvalidError {
	err := CustomInvalidError{}
	err.SetContext(context)
	err.SetType("custom_invalid_error")
	err.SetDescriptionFormat("{{.error}}")
	err.SetDetails(details)
	return &err
}

// NewValidationError is a function that creates a new validation error
func NewValidationError(result *gojsonschema.Result, errField, errDesc string) {
	jsonContext := gojsonschema.NewJsonContext(errField, nil)
	errDetail := gojsonschema.ErrorDetails{
		"error": errDesc,
	}
	result.AddError(
		AddCustomInvalidError(
			gojsonschema.NewJsonContext("error", jsonContext),
			errDetail,
		),
		errDetail,
	)
}

// ValidateJSONPayload checks if the given JSON payload matches the schema
func ValidateJSONPayload(input, schemaFileName string) ([]string, error) {
	body := gojsonschema.NewStringLoader(input)
	factory := loaders.InternalLoaderFactory{}

	schema := factory.New(schemaFileName)
	result, err := gojsonschema.Validate(schema, body)
	if err != nil {
		return nil, err
	}

	// Get the meta version from the JSON body
	var meta map[string]interface{}
	err = json.Unmarshal([]byte(input), &meta)
	if err != nil {
		return nil, err
	}

	// Later add the custom errors using the results.AddError()
	// Get the specific version from the payload map
	version := meta["meta"].(map[string]interface{})

	// Check if version is nil
	if len(version) == 0 {
		NewValidationError(result, "meta", "meta version is missing")
	}

	// Validate the version
	if _, ok := APIVersion[version["version"].(string)]; !ok {
		NewValidationError(result, "meta", "meta version not found: "+version["version"].(string))
	}

	// Get the validation errors
	if result == nil || result.Valid() {
		return nil, nil
	}

	errors := result.Errors()
	out := make([]string, 0)

	for _, e := range errors {
		out = append(out, e.String())
	}
	return out, nil
}

// SchemaValidation is a struct that holds the schema and the validation result
func SchemaValidation(body, schemaFileName string) ([]string, error) {
	validationErrors, err := ValidateJSONPayload(body, schemaFileName)
	if err != nil {
		return nil, err
	}

	if len(validationErrors) > 0 {
		return validationErrors, errors.New(fmt.Sprint(validationErrors))
	}
	return nil, nil
}
