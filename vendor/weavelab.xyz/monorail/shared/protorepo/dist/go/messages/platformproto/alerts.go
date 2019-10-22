package platformproto

import "time"

// TODO: this will be removed and replaced with proto message when Jacob can figure out how to add custom json tags or when the desktop client updates
// DesktopPhoneAlert is what the desktop client is expecting
type DesktopPhoneAlert struct {
	UUID           string `json:"uuid"`
	LocationID     string `json:"location_id"`
	CallerIDName   string `json:"caller_id_name"`
	CallerIDNumber string `json:"caller_id_number"`
	PhoneNumber    string `json:"phone_number"`
	Event          string `json:"event"`
	FirstName      string `json:"first_name"`
	LastName       string `json:"last_name"`
	HouseholdID    string `json:"household_id"`
	HouseholdName  string `json:"household_name"`
	Queue          string `json:"queue"`
	Workstation    string `json:"workstation"`
	Employee       string `json:"employee"`
	Version        string `json:"version"`
	PersonID       string `json:"PersonID"`
}

// TODO: this will be removed and replaced with proto message when Jacob can figure out how to add custom json tags or when the desktop client updates
// DesktopSMSAlert is what the desktop client is expecting
type DesktopSMSAlert struct {
	OpenTracing   string   `json:"OpenTracing"`
	PersonID      string   `json:"PersonID"`
	Action        string   `json:"action"`
	AptID         string   `json:"apt_id"`
	Body          string   `json:"body"`
	FirstName     string   `json:"first_name"`
	HouseholdID   string   `json:"household_id"`
	IsOutbound    bool     `json:"is_outbound"`
	LastName      string   `json:"last_name"`
	LocationID    string   `json:"location_id"`
	PatientID     string   `json:"patient_id"`
	PatientNumber string   `json:"patient_number"`
	PhoneMobile   string   `json:"phone_mobile"`
	SmsID         string   `json:"sms_id"`
	Status        string   `json:"status"`
	To            string   `json:"to"`
	UUID          string   `json:"uuid"`
	WeaveNumber   string   `json:"weave_number"`
	Version       string   `json:"version"`
	MediaIDs      []string `json:"media_ids"`
	NumMedia      int64    `json:"num_media"`
}

type DesktopFollowupAlert struct {
	Datetime time.Time `json:"datetime"`
	ID       string    `json:"id"`
	Method   string    `json:"method"`
	Text     string    `json:"text"`
	Patient  Patient   `json:"patient"`

	LocationSlug string `json:"location_slug"`
	LocationID   string `json:"location_id"`
	Workstation  string `json:"workstation"`
	Employee     string `json:"employee"`
	Version      string `json:"version"`
}

type Patient struct {
	FirstName     string `json:"first_name"`
	ID            string `json:"id"`
	IsGuardian    bool   `json:"is_guardian"`
	LastName      string `json:"last_name"`
	Photo         string `json:"photo"`
	PreferredName string `json:"preferred_name"`
	Status        string `json:"status"`
	PhoneNumber   string `json:"phone_number"`
	Email         string `json:"email"`
	HouseholdID   string `json:"household_id"`
}
