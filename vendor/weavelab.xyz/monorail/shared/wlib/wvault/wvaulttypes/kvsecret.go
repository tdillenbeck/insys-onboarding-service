package wvaulttypes

import (
	"encoding/json"
	"time"
)

type KVSecretResponse struct {
	RequestID     string `json:"request_id"`
	LeaseID       string `json:"lease_id"`
	Renewable     bool   `json:"renewable"`
	LeaseDuration int    `json:"lease_duration"`

	Data KVSecretResponseData `json:"data"`
}

type KVSecretResponseData struct {
	Data     map[string]json.RawMessage `json:"data"`
	MetaData MetaDataResponseData       `json:"metadata"`
}

type MetaDataResponseData struct {
	CreatedTime  time.Time `json:"created_time"`
	DeletionTime string    `json:"deletion_time"`
	Destroyed    bool      `json:"destroyed"`
	Version      int       `json:"version"`
}
