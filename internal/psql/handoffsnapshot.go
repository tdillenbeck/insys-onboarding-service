package psql

import (
	"context"

	"weavelab.xyz/insys-onboarding-service/internal/app"
	"weavelab.xyz/monorail/shared/wlib/uuid"
	"weavelab.xyz/monorail/shared/wlib/werror"
	"weavelab.xyz/monorail/shared/wlib/wsql"
)

// to ensure we implement the interface correctly.
var _ app.HandOffSnapshotService = &HandOffSnapshotService{}

type HandOffSnapshotService struct {
	DB *wsql.PG
}

func (hos HandOffSnapshotService) CreateOrUpdate(ctx context.Context, snapshot *app.HandOffSnapshot) (*app.HandOffSnapshot, error) {

	var result app.HandOffSnapshot

	query := `
		INSERT INTO insys_onboarding.handoff_snapshots
			(id, onboarders_location_id, csat_recipient_user_id, csat_sent_at, created_at, updated_at)
		VALUES ($1, $2, $3, $4, now(), now())
		ON CONFLICT(onboarders_location_id) DO UPDATE SET
			(csat_recipient_user_id, csat_sent_at, updated_at) = ($3, $4, now())
		RETURNING id, created_at, updated_at;
		`

	row := hos.DB.QueryRowContext(
		ctx,
		query,
		uuid.NewV4(),
		snapshot.OnboardersLocationID,
		snapshot.CustomerSatisfactionSurveyRecipientUserID,
		snapshot.CustomerSatisfactionSurveySentAt,
	)

	err := row.Scan(
		&result.ID,
		&result.CreatedAt,
		&result.UpdatedAt,
	)

	result.CustomerSatisfactionSurveyRecipientUserID = snapshot.CustomerSatisfactionSurveyRecipientUserID
	result.CustomerSatisfactionSurveySentAt = snapshot.CustomerSatisfactionSurveySentAt
	result.OnboardersLocationID = snapshot.OnboardersLocationID

	if err != nil {
		return nil, werror.Wrap(err, "failed to insert or update hand-off snapshot")
	}

	return &result, nil
}
