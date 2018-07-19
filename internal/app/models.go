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

	"weavelab.xyz/go-utilities/null"
	"weavelab.xyz/wlib/uuid"
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

// TaskInstance represent a location's status in a task (Waiting on Customer, Waiting on Weave, Completed, etc)
type TaskInstance struct {
	ID         uuid.UUID
	LocationID uuid.UUID
	CategoryID uuid.UUID
	TaskID     uuid.UUID

	ButtonContent     null.String
	ButtonExternalURL null.String
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