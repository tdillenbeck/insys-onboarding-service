package psql

import (
	"context"

	"weavelab.xyz/insys-onboarding-service/internal/app"
	"weavelab.xyz/monorail/shared/wlib/uuid"
	"weavelab.xyz/monorail/shared/wlib/werror"
	"weavelab.xyz/monorail/shared/wlib/wsql"
)

// to ensure we implement the interface correctly.
var _ app.HandoffSnapshotService = &HandoffSnapshotService{}

type HandoffSnapshotService struct {
	DB *wsql.PG
}

func (hos HandoffSnapshotService) CreateOrUpdate(ctx context.Context, snapshot app.HandoffSnapshot) (app.HandoffSnapshot, error) {

	var result app.HandoffSnapshot

	query := `
		INSERT INTO insys_onboarding.handoff_snapshots
			(id, onboarders_location_id, csat_recipient_user_id, csat_sent_at, created_at, updated_at)
		VALUES ($1, $2, $3, $4, now(), now())
		ON CONFLICT(onboarders_location_id) DO UPDATE SET
			(csat_recipient_user_id, csat_sent_at, updated_at) = ($3, $4, now())
		RETURNING id, onboarders_location_id, csat_recipient_user_id, csat_sent_at, created_at, updated_at;
		`

	row := hos.DB.QueryRowContext(
		ctx,
		query,
		uuid.NewV4(),
		snapshot.OnboardersLocationID,
		snapshot.CSATRecipientUserID,
		snapshot.CSATSentAt,
	)

	err := row.Scan(
		&result.ID,
		&result.OnboardersLocationID,
		&result.CSATRecipientUserID,
		&result.CSATSentAt,
		&result.CreatedAt,
		&result.UpdatedAt,
	)

	if err != nil {
		return app.HandoffSnapshot{}, werror.Wrap(err, "failed to insert or update handoff snapshot")
	}

	return result, nil
}
