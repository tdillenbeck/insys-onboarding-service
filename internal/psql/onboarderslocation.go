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
	// assign onboarder to location
	// check existing tasks for location
	// if tasks, update links for new onboarder

	tx, err := s.DB.BeginTx(ctx, &sql.TxOptions{Isolation: sql.LevelDefault, ReadOnly: false})
	if err != nil {
		return nil, werror.Wrap(err, "error opening a transaction")
	}

	onboardersLocation, err := assignOnboarderToLocation(ctx, tx, onbl)
	if err != nil {
		// abort transaction
	}
	locationTasks, err := getLocationTasks(ctx, tx, onboardersLocation.LocationID)
	if err != nil {
		// abort transaction
	}
	if len(locationTasks) > 0 {
		err := updateTasksInstancesForOnboarder(ctx, tx, onboardersLocation.LocationID, onboardersLocation.OnboarderID)
		if err != nil {
			// abort transaction
		}
	}

	// commit transaction
	return onboardersLocation, nil

}

func assignOnboarderToLocation(ctx context.Context, tx string, onbl *app.OnboardersLocation) (*app.OnboardersLocatin, error) {
	var onboardersLocation app.OnboardersLocation

	query := `
INSERT INTO insys_onboarding.onboarders_location
	(id, onboarder_id, location_id, created_at, updated_at)
VALUES ($1, $2, $3, now(), now())
ON CONFLICT(location_id) DO UPDATE SET
	(onboarder_id, updated_at) = ($2, now())
RETURNING id, onboarder_id, location_id, created_at, updated_at;
`

	row := tx.QueryRowContext(ctx, query, uuid.NewV4().String(), onbl.OnboarderID.String(), onbl.LocationID.String())
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
