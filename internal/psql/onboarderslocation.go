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
	var id, onboarderID, location string
	var onboardersLocation app.OnboardersLocation

	// Use a transaction to force the query to be performed against the primary database
	tx, err := s.DB.BeginTx(ctx, &sql.TxOptions{Isolation: sql.LevelDefault, ReadOnly: false})
	if err != nil {
		return nil, werror.Wrap(err, "error opening a database transaction")
	}
	defer tx.Commit()

	query := `
INSERT INTO insys_onboarding.onboarders_location
	(id, onboarder_id, location_id, created_at, updated_at)
VALUES ($1, $2, $3, now(), now())
ON CONFLICT(location_id) DO UPDATE SET
	(onboarder_id, updated_at) = ($2, now())
RETURNING id, onboarder_id, location_id, created_at, updated_at;
`

	row := tx.QueryRowContext(ctx, query, uuid.NewV4().String(), onbl.OnboarderID.String(), onbl.LocationID.String())
	err = row.Scan(
		&id,
		&onboarderID,
		&location,
		&onboardersLocation.CreatedAt,
		&onboardersLocation.UpdatedAt,
	)
	if err != nil {
		return nil, werror.Wrap(err, "inserting or updating onboarders location")
	}

	onboardersLocationUUID, err := uuid.Parse(id)
	if err != nil {
		return nil, werror.Wrap(err, "failed to parse onboarders location id into uuid")
	}
	onboardersLocation.ID = onboardersLocationUUID

	onboarderUUID, err := uuid.Parse(onboarderID)
	if err != nil {
		return nil, werror.Wrap(err, "failed to parse onboarders location onboarder_id into uuid")
	}
	onboardersLocation.OnboarderID = onboarderUUID

	locationUUID, err := uuid.Parse(location)
	if err != nil {
		return nil, werror.Wrap(err, "failed to parse onboarders location location_id into uuid")
	}
	onboardersLocation.LocationID = locationUUID

	return &onboardersLocation, nil
}

func (s *OnboardersLocationService) ReadByLocationID(ctx context.Context, locationID uuid.UUID) (*app.OnboardersLocation, error) {
	var id, onboarderID, location string
	var onbl app.OnboardersLocation

	query := `
SELECT
	id, onboarder_id, location_id, created_at, updated_at
FROM insys_onboarding.onboarders_location
WHERE location_id = $1
`

	row := s.DB.QueryRowContext(ctx, query, locationID.String())
	err := row.Scan(
		&id,
		&onboarderID,
		&location,
		&onbl.CreatedAt,
		&onbl.UpdatedAt,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, werror.Wrap(err).SetCode(wgrpc.CodeNotFound)
		} else {
			return nil, werror.Wrap(err, "error selecting onboarders location by location id")
		}
	}

	onboardersLocationUUID, err := uuid.Parse(id)
	if err != nil {
		return nil, werror.Wrap(err, "failed to parse onboarders location id into uuid")
	}
	onbl.ID = onboardersLocationUUID

	onboarderUUID, err := uuid.Parse(onboarderID)
	if err != nil {
		return nil, werror.Wrap(err, "failed to parse onboarders location onboarder_id into uuid")
	}
	onbl.OnboarderID = onboarderUUID

	locationUUID, err := uuid.Parse(location)
	if err != nil {
		return nil, werror.Wrap(err, "failed to parse onboarders location location_id into uuid")
	}
	onbl.LocationID = locationUUID

	return &onbl, nil
}
