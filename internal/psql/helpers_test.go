package psql

import (
	"context"
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

func clearExistingData(db *wsql.PG) {
	// Clear existing database values
	clearOnboardersQuery := "DELETE FROM insys_onboarding.onboarders;"
	db.ExecContext(context.Background(), clearOnboardersQuery)

	clearOnboardersLocationQuery := "DELETE FROM insys_onboarding.onboarders_location;"
	db.ExecContext(context.Background(), clearOnboardersLocationQuery)

	clearOnboardingTaskInstancesQuery := "DELETE FROM insys_onboarding.onboarding_task_instances;"
	db.ExecContext(context.Background(), clearOnboardingTaskInstancesQuery)

	clearOnboardingTasksQuery := "DELETE FROM insys_onboarding.onboarding_tasks;"
	db.ExecContext(context.Background(), clearOnboardingTasksQuery)

	clearOnboardingCategoriesQuery := "DELETE FROM insys_onboarding.onboarding_categories;"
	db.ExecContext(context.Background(), clearOnboardingCategoriesQuery)
}
