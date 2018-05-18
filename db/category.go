package db

import (
	"context"
	"database/sql"
	"time"

	"weavelab.xyz/wlib/uuid"
	"weavelab.xyz/wlib/werror"
	"weavelab.xyz/wlib/wgrpc"
)

type OnboardingCategory struct {
	ID           uuid.UUID
	DisplayText  string
	DisplayOrder int
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

func Category(ctx context.Context, id uuid.UUID) (OnboardingCategory, error) {
	var catID string
	var cat OnboardingCategory

	query := `SELECT oc.id, oc.display_text, oc.display_order, oc.created_at, oc.updated_at
	FROM insys_onboarding.onboarding_categories AS oc
	WHERE id = $1`

	err := Conn.QueryRowContext(ctx, query, id.String()).Scan(&catID, &cat.DisplayText, &cat.DisplayOrder, &cat.CreatedAt, &cat.UpdatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return OnboardingCategory{}, werror.Wrap(err).SetCode(wgrpc.CodeNotFound)
		} else {
			return cat, werror.Wrap(err, "error querying Category")
		}
	}

	categoryUUID, err := uuid.Parse(catID)
	if err != nil {
		return cat, werror.Wrap(err, "error parsing uuid from database into wlib/uuid")
	}

	cat.ID = categoryUUID

	return cat, nil
}
