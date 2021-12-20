package models

import "encoding/json"

// Contents is the struct that holds the contents of the JSON payload
type Contents struct {
	Text string `json:"text"`
}

// MetaInformation is the meta information of user
type MetaInformation struct {
	Version string `json:"version"`
}

// ContentsPayload is the payload of the contents
type ContentsPayload struct {
	Meta     *MetaInformation `json:"meta"`
	Contents json.RawMessage  `json:"contents"`
}

// TopWords is the struct that holds the top words of the JSON payload
type TopWords struct {
	Word  string `json:"word"`
	Count int    `json:"count"`
}
