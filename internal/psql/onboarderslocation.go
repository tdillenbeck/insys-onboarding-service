package psql

import (
	"context"
	"database/sql"
	"fmt"
	"strings"

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
RETURNING id, onboarder_id, location_id, created_at, updated_at, user_first_logged_in_at;
`

	row := s.DB.QueryRowContext(ctx, query, uuid.NewV4().String(), onbl.OnboarderID.String(), onbl.LocationID.String())
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

	query := `
	UPDATE insys_onboarding.onboarders_location
	SET 
		user_first_logged_in_at = now(),
		updated_at = now()
	WHERE location_id = $1;`

	result, err := s.DB.ExecContext(ctx, query, locationID.String())
	if err != nil {
		return werror.Wrap(err, "error setting first user_first_logged_in_at")
	}

	rowsEffected, err := result.RowsAffected()
	if err != nil {
		return werror.Wrap(err, "error reading number of rows affected")
	}

	if rowsEffected == 0 {
		return werror.New("could not find location by locationID").SetCode(wgrpc.CodeNotFound).Add("locationID", locationID.String())
	}

	return nil
}

func (s *OnboardersLocationService) HasLocationsWithoutLoginRecorded(ctx context.Context, locationIDs []uuid.UUID) (bool, error) {
	query := `
		select * from insys_onboarding.onboarders_location where location_id = any($1);
	`

	var locationUUIDStrs []string

	for _, record := range locationIDs {
		locationUUIDStrs = append(locationUUIDStrs, record.String())
	}

	UUIDList := strings.Join(locationUUIDStrs, ", ")
	filter := fmt.Sprintf("{%s}", UUIDList)

	result, err := s.DB.QueryContext(ctx, query, filter)
	if err != nil {
		return false, werror.Wrap(err, "error setting first user_first_logged_in_at")
	}

	for result.Next() {
		var onboarderLocation *app.OnboardersLocation

		result.Scan(onboarderLocation)

		if !onboarderLocation.UserFirstLoggedInAt.Valid {
			return true, nil
		}
	}

	return false, nil
}
