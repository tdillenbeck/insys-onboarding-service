package wsql

import (
	"context"
	"database/sql"
	"strings"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func (p *PG) BeginTx(ctx context.Context, opts *sql.TxOptions) (*sql.Tx, error) {
	stop := p.middleware(ctx, "begintx")
	defer stop()
	return p.rw(ctx).BeginTx(ctx, opts)
}

func (p *PG) BeginTxx(ctx context.Context, opts *sql.TxOptions) (*sqlx.Tx, error) {
	stop := p.middleware(ctx, "begintxx")
	defer stop()
	return p.rw(ctx).BeginTxx(ctx, opts)
}

func (p *PG) ExecContext(ctx context.Context, query string, args ...interface{}) (sql.Result, error) {
	stop := p.middleware(ctx, query, args...)
	defer stop()
	result, err := p.rw(ctx).ExecContext(ctx, query, args...)
	if err != nil {
		return nil, wrapError(err)
	}

	return result, nil
}

func (p *PG) GetContext(ctx context.Context, dest interface{}, query string, args ...interface{}) error {
	stop := p.middleware(ctx, query, args...)
	defer stop()
	err := p.r(ctx).GetContext(ctx, dest, query, args...)
	if err != nil {
		return wrapError(err)
	}

	return nil
}

func (p *PG) NamedExecContext(ctx context.Context, query string, args interface{}) (sql.Result, error) {
	stop := p.middleware(ctx, query, args)
	defer stop()
	result, err := p.rw(ctx).NamedExecContext(ctx, query, args)
	if err != nil {
		return nil, wrapError(err)
	}

	return result, nil
}

func (p *PG) NamedQueryContext(ctx context.Context, query string, args interface{}) (*sqlx.Rows, error) {
	stop := p.middleware(ctx, query, args)
	defer stop()

	db := p.r(ctx)
	if isPrimary(query) {
		db = p.rw(ctx)
	}

	result, err := db.NamedQueryContext(ctx, query, args)
	if err != nil {
		return nil, wrapError(err)
	}

	return result, nil
}

func (p *PG) QueryContext(ctx context.Context, query string, args ...interface{}) (*sql.Rows, error) {
	stop := p.middleware(ctx, query, args...)
	defer stop()

	db := p.r(ctx)
	if isPrimary(query) {
		db = p.rw(ctx)
	}

	result, err := db.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, wrapError(err)
	}

	return result, nil
}

func (p *PG) QueryxContext(ctx context.Context, query string, args ...interface{}) (*sqlx.Rows, error) {
	stop := p.middleware(ctx, query, args...)
	defer stop()

	db := p.r(ctx)
	if isPrimary(query) {
		db = p.rw(ctx)
	}

	return db.QueryxContext(ctx, query, args...)
}

func (p *PG) QueryRowContext(ctx context.Context, query string, args ...interface{}) *sql.Row {
	stop := p.middleware(ctx, query, args...)
	defer stop()

	db := p.r(ctx)
	if isPrimary(query) {
		db = p.rw(ctx)
	}

	row := db.QueryRowContext(ctx, query, args...)

	return row
}

func (p *PG) QueryRowxContext(ctx context.Context, query string, args ...interface{}) *sqlx.Row {
	stop := p.middleware(ctx, query, args...)
	defer stop()

	db := p.r(ctx)
	if isPrimary(query) {
		db = p.rw(ctx)
	}

	row := db.QueryRowxContext(ctx, query, args...)

	return row
}

func (p *PG) SelectContext(ctx context.Context, dest interface{}, query string, args ...interface{}) error {
	stop := p.middleware(ctx, query, args...)
	defer stop()

	db := p.r(ctx)
	if isPrimary(query) {
		db = p.rw(ctx)
	}

	err := db.SelectContext(ctx, dest, query, args...)
	if err != nil {
		return wrapError(err)
	}

	return nil
}

func isPrimary(q string) bool {

	tq := strings.TrimSpace(q)
	tq = strings.ToLower(tq)
	if strings.HasPrefix(tq, "insert into") {
		return true
	}

	if strings.HasPrefix(tq, "update") {
		return true
	}

	if strings.HasPrefix(tq, "delete from") {
		return true
	}

	return false
}

type Query struct {
	Type    string
	Columns []string
	Tables  string
	Where   []string
	GroupBy []string
	Having  []string
	OrderBy string
	Limit   string
}

func (q *Query) String() string {
	query := q.Type + " " + strings.Join(q.Columns, ", ") + "\n\tFROM " + q.Tables
	if len(q.Where) > 0 {
		query += "\n\tWHERE " + strings.Join(q.Where, " AND ")
	}
	if len(q.GroupBy) > 0 {
		query += "\n\tGROUP BY " + strings.Join(q.GroupBy, ", ")
	}
	if len(q.Having) > 0 {
		query += "\n\tHAVING " + strings.Join(q.Having, " AND ")
	}
	if len(q.OrderBy) > 0 {
		query += "\n\tORDER BY" + q.OrderBy
	}
	if len(q.Limit) > 0 {
		query += "\n\t" + q.Limit
	}

	return query
}

func (q *Query) AddWhere(f string) {
	q.Where = append(q.Where, f)
}
