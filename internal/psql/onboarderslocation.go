package psql

import (
	"context"
	"database/sql"
	"time"

	"weavelab.xyz/insys-onboarding-service/internal/app"

	"weavelab.xyz/monorail/shared/go-utilities/null"
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
	(id, onboarder_id, location_id, created_at, updated_at, user_first_logged_in_at)
VALUES ($1, $2, $3, now(), now(), $4)
ON CONFLICT(location_id) DO UPDATE SET
	(onboarder_id, updated_at, user_first_logged_in_at) = ($2, now(), $4)
RETURNING id, onboarder_id, location_id, created_at, updated_at, user_first_logged_in_at;
`

	row := s.DB.QueryRowContext(ctx, query, uuid.NewV4().String(), onbl.OnboarderID.String(), onbl.LocationID.String(), onbl.UserFirstLoggedInAt.Time)
	err := row.Scan(
		&onboardersLocation.ID,
		&onboardersLocation.OnboarderID,
		&onboardersLocation.LocationID,
		&onboardersLocation.CreatedAt,
		&onboardersLocation.UpdatedAt,
		&onboardersLocation.UserFirstLoggedInAt,
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
	id, onboarder_id, location_id, created_at, updated_at, user_first_logged_in_at
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
		&onboardersLocation.UserFirstLoggedInAt,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, werror.Wrap(err).SetCode(wgrpc.CodeNotFound)
		}
		return nil, werror.Wrap(err, "error selecting onboarders location by location id")
	}

	return &onboardersLocation, nil
}

func (s *OnboardersLocationService) RecordFirstLogin(ctx context.Context, locationID uuid.UUID) error {
	onboardersLocation, err := s.ReadByLocationID(ctx, locationID)
	if err != nil {
		return werror.Wrap(err, "error selecting onboarders location by location id")
	}

	if onboardersLocation.UserFirstLoggedInAt.Valid {
		return nil
	}

	onboardersLocation.UserFirstLoggedInAt = null.NewTime(time.Now())

	_, err = s.CreateOrUpdate(ctx, onboardersLocation)
	if err != nil {
		return werror.Wrap(err, "error updating onboarders location by location id")
	}

	return nil
}
