package psql

import (
	"context"
	"database/sql"

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
		(
			id,
			onboarders_location_id,
			csat_sent_at,
			created_at,
			updated_at,
			handed_off_at,
			point_of_contact_email,
			reason_for_purchase,
			customizations,
			customization_setup,
			fax_port_submitted,
			router_type,
			disclaimer_type_sent,
			router_make_and_model,
			network_decision,
			billing_notes,
			notes
		)
		VALUES 
		(
			$1,
			$2,
			NULL,
			now(),
			now(),
			NULL,
			$3,
			$4,
			$5,
			$6,
			$7,
			$8,
			$9,
			$10,
			$11,
			$12,
			$13
		)
		ON CONFLICT(onboarders_location_id) DO UPDATE SET
		(
			updated_at,
			point_of_contact_email,
			reason_for_purchase,
			customizations,
			customization_setup,
			fax_port_submitted,
			router_type,
			disclaimer_type_sent,
			router_make_and_model,
			network_decision,
			billing_notes,
			notes
		)
		=
		(
			now(),
			$3,
			$4,
			$5,
			$6,
			$7,
			$8,
			$9,
			$10,
			$11,
			$12,
			$13
		)
		RETURNING
			id,
			onboarders_location_id,
			csat_recipient_user_email,
			csat_sent_at,
			created_at,
			updated_at,
			handed_off_at,
			point_of_contact_email,
			reason_for_purchase,
			customizations,
			customization_setup,
			fax_port_submitted,
			router_type,
			disclaimer_type_sent,
			router_make_and_model,
			network_decision,
			billing_notes,
			notes;
		`

	row := hos.DB.QueryRowContext(
		ctx,
		query,
		uuid.NewV4(),
		snapshot.OnboardersLocationID,
		snapshot.PointOfContactEmail,
		snapshot.ReasonForPurchase,
		snapshot.Customizations,
		snapshot.CustomizationSetup,
		snapshot.FaxPortSubmitted,
		snapshot.RouterType,
		snapshot.DisclaimerTypeSent,
		snapshot.RouterMakeAndModel,
		snapshot.NetworkDecision,
		snapshot.BillingNotes,
		snapshot.Notes,
	)

	err := row.Scan(
		&result.ID,
		&result.OnboardersLocationID,
		&result.CsatRecipientUserEmail,
		&result.CSATSentAt,
		&result.CreatedAt,
		&result.UpdatedAt,
		&result.HandedOffAt,
		&result.PointOfContactEmail,
		&result.ReasonForPurchase,
		&result.Customizations,
		&result.CustomizationSetup,
		&result.FaxPortSubmitted,
		&result.RouterType,
		&result.DisclaimerTypeSent,
		&result.RouterMakeAndModel,
		&result.NetworkDecision,
		&result.BillingNotes,
		&result.Notes,
	)

	if err != nil {
		return app.HandoffSnapshot{}, werror.Wrap(err, "failed to insert or update handoff snapshot")
	}

	return result, nil
}

func (hos HandoffSnapshotService) ReadByOnboardersLocationID(ctx context.Context, onboardersLocationID uuid.UUID) (app.HandoffSnapshot, error) {
	var result app.HandoffSnapshot
	query := `
		SELECT 
			id,
			onboarders_location_id,
			csat_recipient_user_email,
			csat_sent_at,
			created_at,
			updated_at,
			handed_off_at,
			point_of_contact_email,
			reason_for_purchase,
			customizations,
			customization_setup,
			fax_port_submitted,
			router_type,
			disclaimer_type_sent,
			router_make_and_model,
			network_decision,
			billing_notes,
			notes
		FROM
			insys_onboarding.handoff_snapshots
			WHERE onboarders_location_id = $1
		`

	row := hos.DB.QueryRowxContext(ctx, query, onboardersLocationID.String())

	err := row.StructScan(&result)
	if err != nil {
		return app.HandoffSnapshot{}, err
	}

	return result, nil
}

func (hos HandoffSnapshotService) SubmitCSAT(ctx context.Context, onboardersLocationID uuid.UUID, csatRecipientUserEmail string) (app.HandoffSnapshot, error) {
	var result app.HandoffSnapshot

	query := `
		UPDATE insys_onboarding.handoff_snapshots 
		SET
			csat_recipient_user_email = $2,
			csat_sent_at = now()
		WHERE onboarders_location_id = $1
		RETURNING
			id,
			onboarders_location_id,
			csat_recipient_user_email,
			csat_sent_at,
			created_at,
			updated_at,
			handed_off_at,
			point_of_contact_email,
			reason_for_purchase,
			customizations,
			customization_setup,
			fax_port_submitted,
			router_type,
			disclaimer_type_sent,
			router_make_and_model,
			network_decision,
			billing_notes,
			notes
		`

	row := hos.DB.QueryRowxContext(ctx, query, onboardersLocationID.String(), csatRecipientUserEmail)

	err := row.StructScan(&result)
	if err != nil {
		if err == sql.ErrNoRows {
			return app.HandoffSnapshot{}, werror.Wrap(err, "no handoff snapshot found with that location id")
		}
		return app.HandoffSnapshot{}, werror.Wrap(err, "error marshalling result into handoff snapshot")
	}

	return result, nil
}

func (hos HandoffSnapshotService) SubmitHandoff(ctx context.Context, onboardersLocationID uuid.UUID) (app.HandoffSnapshot, error) {
	var result app.HandoffSnapshot

	query := `
		UPDATE insys_onboarding.handoff_snapshots 
		SET
			handed_off_at = now()
		WHERE onboarders_location_id = $1 AND handed_off_at IS NULL
		RETURNING
			id,
			onboarders_location_id,
			csat_recipient_user_email,
			csat_sent_at,
			created_at,
			updated_at,
			handed_off_at,
			point_of_contact_email,
			reason_for_purchase,
			customizations,
			customization_setup,
			fax_port_submitted,
			router_type,
			disclaimer_type_sent,
			router_make_and_model,
			network_decision,
			billing_notes,
			notes
		`

	row := hos.DB.QueryRowxContext(ctx, query, onboardersLocationID.String())

	err := row.StructScan(&result)
	if err != nil {
		if err == sql.ErrNoRows {
			return app.HandoffSnapshot{}, werror.Wrap(err, "no handoff snapshot found with that location id")
		}
		return app.HandoffSnapshot{}, werror.Wrap(err, "error marshalling result into handoff snapshot")
	}

	return result, nil
}
