package wsql

import (
	"context"
	"fmt"
	"strings"

	"weavelab.xyz/monorail/shared/go-utilities/null"
	"weavelab.xyz/monorail/shared/wlib/werror"
)

//Update executes a SQL update for the given record on every record where idCol = id in the given table. It returns the number of records that were affected.
//has no decoder for GRPC since we already have the object in memeory and do not need to decode it.
//Table needs `deleted` column for this update call to work
func (p *PG) Update(ctx context.Context, table string, idCol string, id string, record interface{}) (int64, error) {

	setSQL, values, err := null.ColumnsValues(record)
	if err != nil {
		return 0, werror.Wrap(err, "Could not get columns and values to update record").Add("id", id)
	}

	values[idCol] = id

	query := fmt.Sprintf(`UPDATE %[1]s
				SET %[2]s
				WHERE %[3]s = :%[3]s
					AND deleted = false`,
		table,
		strings.Join(setSQL, ", "),
		idCol)

	res, err := p.NamedExecContext(ctx, query, values)
	if err != nil {
		return 0, wrapError(err)
	}

	return res.RowsAffected()

}
