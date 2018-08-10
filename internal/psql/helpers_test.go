package psql

import (
	"os"
	"testing"

	"weavelab.xyz/wlib/wsql"
)

const (
	psqlConnString = "postgresql://127.0.0.1:5432/insys_onboarding_test?sslmode=disable"
)

func skipCI(t *testing.T) {
	if os.Getenv("CI") != "" {
		t.Skip("Skipping testing in CI environment")
	}
}

func initDBConnection(t *testing.T, dbConnString string) *wsql.PG {
	settings := wsql.Settings{}
	settings.PrimaryConnectString.SetConnectString(dbConnString)

	conn, err := wsql.New(&settings)
	if err != nil {
		t.Errorf("could not connect to test database. make sure the test database has been created and is running. connection string: %v", psqlConnString)
		return nil
	}
	return conn
}
