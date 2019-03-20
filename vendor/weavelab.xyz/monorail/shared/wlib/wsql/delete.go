package wsql

import (
	"context"
	"database/sql"

	"weavelab.xyz/monorail/shared/wlib/werror"
)

func (p *PG) DeleteRow(ctx context.Context, table string, key string, id string) (int64, error) {

	queries := []string{
		"DELETE FROM " + table + " WHERE " + key + "=$1 AND " + id + "=$2 AND deleted=true",
		"UPDATE " + table + " SET deleted=true WHERE " + key + "=$1 AND " + id + "=$2",
	}

	opts := sql.TxOptions{}
	tx, err := p.BeginTx(ctx, &opts)
	if err != nil {
		return 0, werror.Wrap(err, "error beginning tx")
	}

	var result sql.Result
	for _, q := range queries {
		result, err = tx.ExecContext(ctx, q, key, id)
		if err != nil {
			_ = tx.Rollback()
			return 0, werror.Wrap(wrapError(err), "error exec query")
		}
	}

	err = tx.Commit()
	if err != nil {
		return 0, werror.Wrap(wrapError(err), "error committing tx")
	}

	return result.RowsAffected()
}
