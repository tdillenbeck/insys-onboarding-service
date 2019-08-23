package wsql

import (
	"context"
	"fmt"
	"runtime"

	"github.com/jmoiron/sqlx"
	"weavelab.xyz/monorail/shared/wlib/werror"
)

func BuildListSelect(returnTotalRows bool, columns []string, tables string, allowed Parameters, options QueryOptions) (Query, map[string]interface{}, error) {
	return buildList("SELECT", returnTotalRows, columns, tables, allowed, options)
}

func BuildListDelete(tables string, allowed Parameters, options QueryOptions) (Query, map[string]interface{}, error) {

	// check for user supplied parameters that are not supported
	if len(options.Orders) > 0 {
		return Query{}, nil, fmt.Errorf("unable to add order by to delete query")
	}

	if options.Limits.limit > 0 {
		return Query{}, nil, fmt.Errorf("unable to add limit to delete query")
	}

	if options.Limits.skip > 0 {
		return Query{}, nil, fmt.Errorf("unable to add skip to delete query")
	}

	allowed.DefaultOrder = nil
	options.Orders = nil
	allowed.ListMax = 0

	return buildList("DELETE", false, nil, tables, allowed, options)
}

func buildList(queryType string, returnTotalRows bool, columns []string, tables string, allowed Parameters, options QueryOptions) (Query, map[string]interface{}, error) {
	filterSQL, filterParameters, err := options.FilterSlice(allowed.Filters)
	if err != nil {
		return Query{}, nil, err
	}

	query := Query{
		Type:   queryType,
		Tables: tables,
		Where:  filterSQL,
	}

	if returnTotalRows {
		query.Columns = []string{`count(*) AS count`}
		return query, filterParameters, nil
	}

	orderSQL, err := options.OrderSQL(allowed.DefaultOrder, allowed.Orders)
	if err != nil {
		return Query{}, nil, err
	}

	limitSQL, err := options.LimitSQL(allowed.ListMax)
	if err != nil {
		return Query{}, nil, err
	}

	query.Columns = columns
	query.Limit = limitSQL
	query.OrderBy = orderSQL
	query.GroupBy = options.GroupBy
	query.Having = options.Having

	return query, filterParameters, nil
}

func (pg *PG) ListQueryHelperContext(ctx context.Context, total bool, query string, parameters map[string]interface{}, results interface{}) (interface{}, error) {

	ctx = setCallerName(ctx)

	q, p, err := sqlx.Named(query, parameters)
	if err != nil {
		return nil, fmt.Errorf("unable to generate query: %s", err)
	}

	q = sqlx.Rebind(sqlx.DOLLAR, q)

	if total {
		var totalRows Total

		row := pg.QueryRowContext(ctx, q, p...)

		err = row.Scan(&totalRows.Total)
		if err != nil {
			return nil, fmt.Errorf("unable to get row count: %s", err)
		}

		return totalRows, nil
	}

	err = pg.SelectContext(ctx, results, q, p...)
	if err != nil {
		return nil, fmt.Errorf("unable to run db.select: %s", err)
	}

	// check number of results
	l := sliceLength(results)
	if l > 500 {
		wsqlLogger.WErrorC(ctx, werror.New("too many results").Add("count", l).Add("query", query))
	}

	return results, nil
}

func setCallerName(ctx context.Context) context.Context {

	var callerName string

	stack := make([]uintptr, 1)
	_ = runtime.Callers(3, stack)

	for _, v := range stack {
		f := runtime.FuncForPC(v - 1)
		if f != nil {
			callerName = f.Name()
			break
		}
	}

	ctx = SetCallerName(ctx, callerName)

	return ctx
}
