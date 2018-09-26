package psql

import (
	"context"
	"database/sql"

	"weavelab.xyz/insys-onboarding/internal/app"
	"weavelab.xyz/wlib/uuid"
	"weavelab.xyz/wlib/werror"
	"weavelab.xyz/wlib/wgrpc"
	"weavelab.xyz/wlib/wsql"
)

type OnboarderService struct {
	DB *wsql.PG
}

func (s *OnboarderService) CreateOrUpdate(ctx context.Context, onb *app.Onboarder) (*app.Onboarder, error) {
	var id, userID string
	var onboarder app.Onboarder

	query := `
INSERT INTO insys_onboarding.onboarders
	(id, user_id, schedule_customization_link, schedule_porting_link, schedule_network_link, schedule_software_install_link, schedule_phone_install_link, schedule_software_training_link, schedule_phone_training_link, created_at, updated_at)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, now(), now())
ON CONFLICT(user_id) DO UPDATE SET
	(schedule_customization_link, schedule_porting_link, schedule_network_link, schedule_software_install_link, schedule_phone_install_link, schedule_software_training_link, schedule_phone_training_link, updated_at)
	= ($3, $4, $5, $6, $7, $8, $9, now())
RETURNING id, user_id, schedule_customization_link, schedule_porting_link, schedule_network_link, schedule_software_install_link, schedule_phone_install_link, schedule_software_training_link, schedule_phone_training_link, created_at, updated_at;
`

	row := s.DB.QueryRowContext(ctx, query, uuid.NewV4().String(), onb.UserID.String(), onb.ScheduleCustomizationLink, onb.SchedulePortingLink, onb.ScheduleNetworkLink, onb.ScheduleSoftwareInstallLink, onb.SchedulePhoneInstallLink, onb.ScheduleSoftwareTrainingLink, onb.SchedulePhoneTrainingLink)
	err := row.Scan(
		&id,
		&userID,
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

	onboarderUUID, err := uuid.Parse(id)
	if err != nil {
		return nil, werror.Wrap(err, "failed to parse onboarder id into uuid")
	}
	onboarder.ID = onboarderUUID

	userUUID, err := uuid.Parse(userID)
	if err != nil {
		return nil, werror.Wrap(err, "failed to parse onboarder user_id into uuid")
	}
	onboarder.UserID = userUUID

	return &onboarder, nil
}

func (s *OnboarderService) ReadByUserID(ctx context.Context, userID uuid.UUID) (*app.Onboarder, error) {
	var id, user string
	var onboarder app.Onboarder

	query := `
SELECT
	id, user_id, schedule_customization_link, schedule_porting_link, schedule_network_link, schedule_software_install_link, schedule_phone_install_link,schedule_software_training_link, schedule_phone_training_link, created_at, updated_at
FROM insys_onboarding.onboarders
WHERE user_id = $1`

	row := s.DB.QueryRowContext(ctx, query, userID.String())
	err := row.Scan(
		&id,
		&user,
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

	onboarderUUID, err := uuid.Parse(id)
	if err != nil {
		return nil, werror.Wrap(err, "failed to parse onboarder id into uuid")
	}
	onboarder.ID = onboarderUUID

	userUUID, err := uuid.Parse(user)
	if err != nil {
		return nil, werror.Wrap(err, "failed to parse onboarder user_id into uuid")
	}
	onboarder.UserID = userUUID

	return &onboarder, nil
}
