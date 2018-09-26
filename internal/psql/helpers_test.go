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

func clearExistingData(t *testing.T, db *wsql.PG) {
	// Clear existing database values

	clearOnboardersLocationQuery := "DELETE FROM insys_onboarding.onboarders_location;"
	_, err := db.ExecContext(context.Background(), clearOnboardersLocationQuery)
	if err != nil {
		t.Fatalf("could not delete onboarders_location. error: %v\n", err)
	}

	clearOnboardersQuery := "DELETE FROM insys_onboarding.onboarders;"
	_, err = db.ExecContext(context.Background(), clearOnboardersQuery)
	if err != nil {
		t.Fatalf("could not delete onboarders. error: %v\n", err)
	}

	clearOnboardingTaskInstancesQuery := "DELETE FROM insys_onboarding.onboarding_task_instances;"
	_, err = db.ExecContext(context.Background(), clearOnboardingTaskInstancesQuery)
	if err != nil {
		t.Fatalf("could not delete onboarding_task_instances. error: %v\n", err)
	}

	clearOnboardingTasksQuery := "DELETE FROM insys_onboarding.onboarding_tasks;"
	_, err = db.ExecContext(context.Background(), clearOnboardingTasksQuery)
	if err != nil {
		t.Fatalf("could not delete onboarding_tasks. error: %v\n", err)
	}

	clearOnboardingCategoriesQuery := "DELETE FROM insys_onboarding.onboarding_categories;"
	_, err = db.ExecContext(context.Background(), clearOnboardingCategoriesQuery)
	if err != nil {
		t.Fatalf("could not delete onboarding_categories. error: %v\n", err)
	}
}
