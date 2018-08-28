package wsql

import (
	"github.com/pressly/goose"
	"weavelab.xyz/wlib/werror"
)

func EnsureLatest(conn *PG, migrationPath string) error {
	err := goose.Up(conn.db.xdb.DB, migrationPath)
	if err != nil {
		return werror.Wrap(err, "unable to migrate")
	}

	return nil
}

func FullDown(conn *PG, migrationPath string) error {
	err := goose.Reset(conn.db.xdb.DB, migrationPath)
	if err != nil {
		return werror.Wrap(err, "unable to reset database")
	}

	return nil
}
