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

type OnboarderService struct {
	DB *wsql.PG
}

func (s *OnboarderService) CreateOrUpdate(ctx context.Context, onb *app.Onboarder) (*app.Onboarder, error) {
	var onboarder app.Onboarder

	query := `
		INSERT INTO insys_onboarding.onboarders
		(
			id,
			user_id,
			salesforce_user_id,
			schedule_customization_link,
			schedule_porting_link,
			schedule_network_link,
			schedule_software_install_link,
			schedule_phone_install_link,
			schedule_software_training_link,
			schedule_phone_training_link,
			created_at,
			updated_at
		)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, now(), now())
		ON CONFLICT(user_id) DO UPDATE SET
		(
			salesforce_user_id,
			schedule_customization_link,
			schedule_porting_link,
			schedule_network_link,
			schedule_software_install_link,
			schedule_phone_install_link,
			schedule_software_training_link,
			schedule_phone_training_link,
			updated_at
		) = ($3, $4, $5, $6, $7, $8, $9, $10, now())
		RETURNING id, user_id, salesforce_user_id, schedule_customization_link, schedule_porting_link, schedule_network_link, schedule_software_install_link, schedule_phone_install_link, schedule_software_training_link, schedule_phone_training_link, created_at, updated_at;
		`

	row := s.DB.QueryRowContext(ctx, query, uuid.NewV4().String(), onb.UserID.String(), onb.SalesforceUserID, onb.ScheduleCustomizationLink, onb.SchedulePortingLink, onb.ScheduleNetworkLink, onb.ScheduleSoftwareInstallLink, onb.SchedulePhoneInstallLink, onb.ScheduleSoftwareTrainingLink, onb.SchedulePhoneTrainingLink)
	err := row.Scan(
		&onboarder.ID,
		&onboarder.UserID,
		&onboarder.SalesforceUserID,
		&onboarder.ScheduleCustomizationLink,
		&onboarder.SchedulePortingLink,
		&onboarder.ScheduleNetworkLink,
		&onboarder.ScheduleSoftwareInstallLink,
		&onboarder.SchedulePhoneInstallLink,
		&onboarder.ScheduleSoftwareTrainingLink,
		&onboarder.SchedulePhoneTrainingLink,
		&onboarder.CreatedAt,
		&onboarder.UpdatedAt,
	)
	if err != nil {
		return nil, werror.Wrap(err, "inserting or updating onboarder")
	}

	return &onboarder, nil
}

// Delete will soft delete the onboarder. This allows us to preserve a history of which onboarder
// handled which location
func (s *OnboarderService) Delete(ctx context.Context, id uuid.UUID) error {
	query := `
	  UPDATE insys_onboarding.onboarders
			SET deleted_at = now()
		 WHERE id=$1`
	_, err := s.DB.ExecContext(ctx, query, id)
	if err != nil {
		return werror.Wrap(err, "could not soft delete onboarder").Add("id", id)
	}

	return nil
}

func (s *OnboarderService) List(ctx context.Context) ([]app.Onboarder, error) {
	var result []app.Onboarder
	query := `
		SELECT
			id, user_id, salesforce_user_id, schedule_customization_link, schedule_porting_link, schedule_network_link, schedule_software_install_link, schedule_phone_install_link, schedule_software_training_link, schedule_phone_training_link, created_at, updated_at, deleted_at
		FROM insys_onboarding.onboarders`

	rows, err := s.DB.QueryxContext(ctx, query)
	if err != nil {
		return nil, werror.Wrap(err, "error executing ByLocationID query")
	}
	defer rows.Close()

	for rows.Next() {
		var onboarder app.Onboarder
		err = rows.StructScan(&onboarder)
		if err != nil {
			return nil, werror.Wrap(err, "error scanning onboarder into struct for List onboarders")
		}

		result = append(result, onboarder)
	}

	return result, nil
}

func (s *OnboarderService) ReadByUserID(ctx context.Context, userID uuid.UUID) (*app.Onboarder, error) {
	var onboarder app.Onboarder

	query := `
		SELECT
			id, user_id, salesforce_user_id, schedule_customization_link, schedule_porting_link, schedule_network_link, schedule_software_install_link, schedule_phone_install_link,schedule_software_training_link, schedule_phone_training_link, created_at, updated_at
		FROM insys_onboarding.onboarders
		WHERE user_id = $1`

	row := s.DB.QueryRowContext(ctx, query, userID.String())
	err := row.Scan(
		&onboarder.ID,
		&onboarder.UserID,
		&onboarder.SalesforceUserID,
		&onboarder.ScheduleCustomizationLink,
		&onboarder.SchedulePortingLink,
		&onboarder.ScheduleNetworkLink,
		&onboarder.ScheduleSoftwareInstallLink,
		&onboarder.SchedulePhoneInstallLink,
		&onboarder.ScheduleSoftwareTrainingLink,
		&onboarder.SchedulePhoneTrainingLink,
		&onboarder.CreatedAt,
		&onboarder.UpdatedAt,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, werror.Wrap(err).SetCode(wgrpc.CodeNotFound)
		}
		return nil, werror.Wrap(err, "error selecting onboarder by user id")
	}

	return &onboarder, nil
}
