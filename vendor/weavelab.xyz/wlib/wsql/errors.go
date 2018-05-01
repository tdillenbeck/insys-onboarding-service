package wsql

import (
	"github.com/lib/pq"
	"weavelab.xyz/wlib/werror"
)

var ErrorNotFound = werror.Template("sql: no rows in result set")

func wrapError(err error) error {

	perr, ok := err.(*pq.Error)
	if !ok {
		return err
	}

	werr := werror.Wrap(err).Add("pgCode", perr.Code).Add("pgDetail", perr.Detail).Add("pgHint", perr.Hint)
	return werr

}
