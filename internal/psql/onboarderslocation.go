package psql

import (
	"context"
	"database/sql"

	"weavelab.xyz/insys-onboarding-service/internal/app"

	"weavelab.xyz/monorail/shared/wlib/uuid"
	"weavelab.xyz/monorail/shared/wlib/werror"
	"weavelab.xyz/monorail/shared/wlib/wgrpc"
	"weavelab.xyz/monorail/shared/wlib/wsql"
)

type OnboardersLocationService struct {
	DB *wsql.PG
}

func (s *OnboardersLocationService) CreateOrUpdate(ctx context.Context, onbl *app.OnboardersLocation) (*app.OnboardersLocation, error) {
	var onboardersLocation app.OnboardersLocation

	query := `
INSERT INTO insys_onboarding.onboarders_location
	(id, onboarder_id, location_id, created_at, updated_at)
VALUES ($1, $2, $3, now(), now())
ON CONFLICT(location_id) DO UPDATE SET
	(onboarder_id, updated_at) = ($2, now())
RETURNING id, onboarder_id, location_id, created_at, updated_at;
`

	row := s.DB.QueryRowContext(ctx, query, uuid.NewV4().String(), onbl.OnboarderID.String(), onbl.LocationID.String())
	err := row.Scan(
		&onboardersLocation.ID,
		&onboardersLocation.OnboarderID,
		&onboardersLocation.LocationID,
		&onboardersLocation.CreatedAt,
		&onboardersLocation.UpdatedAt,
	)
	if err != nil {
		return nil, werror.Wrap(err, "inserting or updating onboarders location")
	}

	return &onboardersLocation, nil
}

func (s *OnboardersLocationService) ReadByLocationID(ctx context.Context, locationID uuid.UUID) (*app.OnboardersLocation, error) {
	var onboardersLocation app.OnboardersLocation

	query := `
SELECT
	id, onboarder_id, location_id, created_at, updated_at
FROM insys_onboarding.onboarders_location
WHERE location_id = $1
`

	row := s.DB.QueryRowContext(ctx, query, locationID.String())
	err := row.Scan(
		&onboardersLocation.ID,
		&onboardersLocation.OnboarderID,
		&onboardersLocation.LocationID,
		&onboardersLocation.CreatedAt,
		&onboardersLocation.UpdatedAt,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, werror.Wrap(err).SetCode(wgrpc.CodeNotFound)
		}
		return nil, werror.Wrap(err, "error selecting onboarders location by location id")
	}

	return &onboardersLocation, nil
}

// SyncTaskInstanceLinksFromOnboarderLinks will go to the assigned onbaorder in the onboarderslocation table and lookup the onboarder link
// for the list of task id's.
func (s *OnboardersLocationService) SyncTaskInstanceLinksFromOnboarderLinks(ctx context.Context, locationID uuid.UUID) error {
	onboarderLinkToTaskID := []struct {
		onboarderLinkColumnName string
		taskID                  string
	}{
		{
			"schedule_customization_link",
			"2d2df285-9211-48fc-a057-74f7dee2d9a4",
		},
		{
			"schedule_porting_link",
			"9aec502b-f8b8-4f10-9748-1fe4050eacde",
		},
		{
			"schedule_network_link",
			"7b15e061-8002-4edc-9bf4-f38c6eec6364",
		},
		{
			"schedule_software_install_link",
			"16a6dc91-ec6b-4b09-b591-a5b0dfa92932",
		},
		{
			"schedule_phone_install_link",
			"fd4f656c-c9f1-47b8-96ad-3080b999a843",
		},
		{
			"schedule_software_training_link",
			"c20b65d8-e281-4e62-98f0-4aebf83e0bee",
		},
		{
			"schedule_phone_training_link",
			"47743fae-c775-45d5-8a51-dc7e3371dfa4",
		},
	}

	query := `
UPDATE insys_onboarding.onboarding_task_instances
SET button_external_url = (
	SELECT
			COALESCE(NULLIF((SELECT $3 FROM insys_onboarding.onboarders AS a INNER JOIN insys_onboarding.onboarders_location AS b ON a.id = b.onboarder_id WHERE b.location_id=$1), ''), button_external_url)
	FROM insys_onboarding.onboarding_tasks
	WHERE id=$2
)
WHERE location_id=$1 AND onboarding_task_id=$2;
	`

	for _, mapping := range onboarderLinkToTaskID {
		_, err := s.DB.ExecContext(ctx, query, locationID, mapping.taskID, mapping.onboarderLinkColumnName)
		if err != nil {
			return werror.Wrap(err, "could not sync task instance links for onboarder").Add("locationID", locationID).Add("taskID", mapping.taskID)
		}
	}

	return nil
}
