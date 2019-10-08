// package app defines all of our domain types. Things like our representation of a category, task instance, and
// anything else specific to our domain.
//
// This does not include anything specific to the underlying technology. For instance, if we wanted to define
// a CategoryService interface that described functions we could use to interact with a categories database that
// is fine, but we wouldn't add any database specific implementations here.

// See //medium.com/@benbjohnson/standard-package-layout-7cdbc8391fc1 for more infor on this.
package app

import (
	"time"

	"weavelab.xyz/monorail/shared/go-utilities/null"
	"weavelab.xyz/monorail/shared/wlib/uuid"
)

// Category represents a group of tasks that are related.
//
// For example the 'Software' category groups all the tasks related to setting up Weave software.
type Category struct {
	ID           uuid.UUID
	DisplayText  string
	DisplayOrder int
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

// ChiliPiperScheduleEvent tracks scheduled appointments that happen in chili piper
type ChiliPiperScheduleEvent struct {
	ID         uuid.UUID `db:"id" json:"id"`
	LocationID uuid.UUID `db:"location_id" json:"location_id"`

	EventID string `db:"event_id" json:"event_id"`

	AssigneeID null.String `db:"assignee_id" json:"assignee_id"`
	ContactID  null.String `db:"contact_id" json:"contact_id"`
	EventType  null.String `db:"event_type" json:"event_type"`
	RouteID    null.String `db:"route_id" json:"route_id"`

	StartAt    null.Time `db:"start_at" json:"start_at"`
	EndAt      null.Time `db:"end_at" json:"end_at"`
	CanceledAt null.Time `db:"canceled_at" json:"canceled_at"`

	CreatedAt time.Time `db:"created_at" json:"created_at"`
	UpdatedAt time.Time `db:"updated_at" json:"updated_at"`
}

// Onboarder represents a user's onboarding specific information
type Onboarder struct {
	ID     uuid.UUID `db:"id"`
	UserID uuid.UUID `db:"user_id"`

	SalesforceUserID             null.String `db:"salesforce_user_id"`
	ScheduleCustomizationLink    null.String `db:"schedule_customization_link"`
	ScheduleNetworkLink          null.String `db:"schedule_network_link"`
	SchedulePhoneInstallLink     null.String `db:"schedule_phone_install_link"`
	SchedulePhoneTrainingLink    null.String `db:"schedule_phone_training_link"`
	SchedulePortingLink          null.String `db:"schedule_porting_link"`
	ScheduleSoftwareInstallLink  null.String `db:"schedule_software_install_link"`
	ScheduleSoftwareTrainingLink null.String `db:"schedule_software_training_link"`

	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
	DeletedAt null.Time `db:"deleted_at"`
}

// Assign an onboarder to a location
type OnboardersLocation struct {
	ID                      uuid.UUID
	OnboarderID             uuid.UUID
	LocationID              uuid.UUID
	Region                  string
	SalesforceOpportunityID string
	UserFirstLoggedInAt     null.Time
	CreatedAt               time.Time
	UpdatedAt               time.Time
}

// TaskInstance represent a location's status in a task (Waiting on Customer, Waiting on Weave, Completed, etc)
type TaskInstance struct {
	ID         uuid.UUID
	LocationID uuid.UUID
	CategoryID uuid.UUID
	TaskID     uuid.UUID

	ButtonContent     null.String
	ButtonExternalURL null.String
	ButtonInternalURL null.String
	CompletedAt       null.Time
	CompletedBy       null.String
	VerifiedAt        null.Time
	VerifiedBy        null.String
	Content           string
	DisplayOrder      int
	Status            int
	StatusUpdatedAt   time.Time
	StatusUpdatedBy   null.String
	Title             string
	Explanation       null.String

	CreatedAt time.Time
	UpdatedAt time.Time
}
