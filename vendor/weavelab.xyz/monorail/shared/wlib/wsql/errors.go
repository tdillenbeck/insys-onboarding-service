package wsql

import (
	"reflect"

	"github.com/lib/pq"
	"weavelab.xyz/monorail/shared/wlib/werror"
)

var ErrorNotFound = werror.CodedTemplate("sql: no rows in result set", werror.CodeNotFound)

var (
	// these are both deprecated, should use werror.Wrap
	WrapError = werror.Wrap
	wrapError = werror.Wrap
)

func init() {
	werror.RegisterWrapper(reflect.TypeOf(&pq.Error{}), wrapPQError)
}

func wrapPQError(werr *werror.Error, err error) *werror.Error {

	perr, _ := err.(*pq.Error)

	werr = werr.Add("pgCode", perr.Code).Add("pgDetail", perr.Detail).Add("pgHint", perr.Hint)

	var codePrefix string
	if len(perr.Code) >= 2 {
		codePrefix = string(perr.Code[0:2])
	}

	// https://www.postgresql.org/docs/current/static/errcodes-appendix.html
	switch codePrefix {

	case "02": // No Data
		// replace error with not found error template.
		werr = ErrorNotFound.Here().SetCode(werror.CodeNotFound)

	case "23":
		// Integrity Constraint Violation
		werr.SetCode(werror.CodeInvalidArgument)

	default:
		// default to internal 500 class error
		werr.SetCode(werror.CodeInternal)
	}

	return werr

}
