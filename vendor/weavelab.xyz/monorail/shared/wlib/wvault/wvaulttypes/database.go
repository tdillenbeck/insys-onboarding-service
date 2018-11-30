package wvaulttypes

type CredentialsResponse struct {
	RequestID     string `json:"request_id"`
	LeaseID       string `json:"lease_id"`
	Renewable     bool   `json:"renewable"`
	LeaseDuration int    `json:"lease_duration"`

	Data CredentialsResponseData `json:"data"`
}

type CredentialsResponseData struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
