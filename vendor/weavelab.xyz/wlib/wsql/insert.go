package wsql

import (
	"context"
	"strconv"
	"strings"

	"weavelab.xyz/wlib/werror"
)

type ID interface {
	ID() string
}

type Valid interface {
	Valid() bool
}

type Defaults interface {
	SetDefaultsAsNeeded()
	Clean()
}

// InsertSingle inserts a single record into the database
func (p *PG) Insert(ctx context.Context, table string, record interface{}) error {

	columns, err := ValidColumns(record)
	if err != nil {
		return werror.Wrap(err, "unable to detect record columns")
	}

	queryBase, numFields, err := insertQuery(table, columns)
	if err != nil {
		return werror.Wrap(err, "unable to generate query")
	}

	valueList := insertQueryValues(numFields, 1)
	query := queryBase + strings.Join(valueList, ",")

	if r, ok := record.(Defaults); ok {
		r.SetDefaultsAsNeeded()
		r.Clean()
	}

	values, err := ColumnValues(record)
	if err != nil {
		return werror.Wrap(err, "unable to get column values")
	}

	_, err = p.ExecContext(ctx, query, values...)
	if err != nil {
		return werror.Wrap(wrapError(err), "unable to insert")
	}

	return nil
}

func insertQuery(table string, columns []string) (string, int, error) {

	base := "INSERT INTO " + table + " (" + strings.Join(columns, ",") + ") VALUES\n"

	return base, len(columns), nil
}

func insertQueryValues(columns int, rows int) []string {

	var valueList = make([]string, rows)
	for i := 0; i < rows; i++ {

		var numbers = make([]string, columns)
		for j := 0; j < columns; j++ {
			numbers[j] = "$" + strconv.Itoa(i*columns+j+1)
		}

		valueList[i] = "(" + strings.Join(numbers, ",") + ")"
	}

	return valueList
}
