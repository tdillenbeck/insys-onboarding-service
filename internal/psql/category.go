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

type CategoryService struct {
	DB *wsql.PG
}

func (s *CategoryService) ByID(ctx context.Context, id uuid.UUID) (*app.Category, error) {
	var category app.Category

	query := `SELECT id, display_text, display_order, created_at, updated_at FROM insys_onboarding.onboarding_categories WHERE id = $1`
	row := s.DB.QueryRowContext(ctx, query, id.String())
	err := row.Scan(
		&category.ID,
		&category.DisplayText,
		&category.DisplayOrder,
		&category.CreatedAt,
		&category.UpdatedAt,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, werror.Wrap(err).SetCode(wgrpc.CodeNotFound)
		} else {
			return nil, werror.Wrap(err, "error querying Category")
		}
	}

	return &category, nil
}
