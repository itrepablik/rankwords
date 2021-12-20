package converts

import (
	"fmt"
	"strings"
	"time"
)

// SchemaToBytes replaces the schema's white space with comma and returns the bytes as a string
func SchemaToBytes(JSONSchema string) string {
	b := []byte(JSONSchema)
	byteStr := strings.ReplaceAll(fmt.Sprint(b), " ", ",")
	return byteStr
}

// FormatCurrentDT returns the current date and time in the format of YYYY-MM-DD HH:MM:SS
func FormatCurrentDT() time.Time {
	tranDateTime, _ := time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
	return tranDateTime
}
