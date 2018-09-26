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
	return p.rw().BeginTx(ctx, opts)
}

func (p *PG) BeginTxx(ctx context.Context, opts *sql.TxOptions) (*sqlx.Tx, error) {
	stop := p.middleware(ctx, "begintxx")
	defer stop()
	return p.rw().BeginTxx(ctx, opts)
}

func (p *PG) ExecContext(ctx context.Context, query string, args ...interface{}) (sql.Result, error) {
	stop := p.middleware(ctx, query, args...)
	defer stop()
	result, err := p.rw().ExecContext(ctx, query, args...)
	if err != nil {
		return nil, wrapError(err)
	}

	return result, nil
}

func (p *PG) GetContext(ctx context.Context, dest interface{}, query string, args ...interface{}) error {
	p.logQueryParameters(query, args)
	stop := findStackAndStartTimer()
	defer stop()
	err := p.r().GetContext(ctx, dest, query, args...)
	if err != nil {
		return wrapError(err)
	}

	return nil
}

func (p *PG) NamedExecContext(ctx context.Context, query string, args interface{}) (sql.Result, error) {
	stop := p.middleware(ctx, query, args)
	defer stop()
	result, err := p.rw().NamedExecContext(ctx, query, args)
	if err != nil {
		return nil, wrapError(err)
	}

	return result, nil
}

func (p *PG) NamedQueryContext(ctx context.Context, query string, args interface{}) (*sqlx.Rows, error) {
	stop := p.middleware(ctx, query, args)
	defer stop()

	db := p.r()
	if isPrimary(query) {
		db = p.rw()
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

	db := p.r()
	if isPrimary(query) {
		db = p.rw()
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

	db := p.r()
	if isPrimary(query) {
		db = p.rw()
	}

	return db.QueryxContext(ctx, query, args...)
}

func (p *PG) QueryRowContext(ctx context.Context, query string, args ...interface{}) *sql.Row {
	stop := p.middleware(ctx, query, args...)
	defer stop()

	db := p.r()
	if isPrimary(query) {
		db = p.rw()
	}

	row := db.QueryRowContext(ctx, query, args...)

	return row
}

func (p *PG) QueryRowxContext(ctx context.Context, query string, args ...interface{}) *sqlx.Row {
	stop := p.middleware(ctx, query, args...)
	defer stop()

	db := p.r()
	if isPrimary(query) {
		db = p.rw()
	}

	row := db.QueryRowxContext(ctx, query, args...)

	return row
}

func (p *PG) SelectContext(ctx context.Context, dest interface{}, query string, args ...interface{}) error {
	stop := p.middleware(ctx, query, args...)
	defer stop()

	db := p.r()
	if isPrimary(query) {
		db = p.rw()
	}

	err := db.SelectContext(ctx, dest, query, args...)
	if err != nil {
		return wrapError(err)
	}

	return nil
}

func isPrimary(q string) bool {

	if strings.HasPrefix(q, "INSERT INTO ") {
		return true
	}

	if strings.HasPrefix(q, "UPDATE ") {
		return true
	}

	if strings.HasPrefix(q, "DELETE FROM ") {
		return true
	}

	return false
}
