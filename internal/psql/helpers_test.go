package psql

import (
	"context"
	"os"
	"testing"

	"weavelab.xyz/insys-onboarding-service/internal/config"
	"weavelab.xyz/monorail/shared/wlib/wapp"
	"weavelab.xyz/monorail/shared/wlib/werror"
	"weavelab.xyz/monorail/shared/wlib/wsql"
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
	err := config.Init()
	if err != nil {
		wapp.Exit(werror.Wrap(err, "error initializing config values"))
	}

	dbOptions := &ConnectionOptions{
		MaxOpenConnections:    config.MaxOpenConnections,
		MaxIdleConnections:    config.MaxIdleConnections,
		MaxConnectionLifetime: config.MaxConnectionLifetime,
		LogQueries:            config.LogQueries,
	}

	conn, err := ConnectionFromConnString(context.Background(), config.PrimaryConnString, config.PrimaryConnString, dbOptions)
	if err != nil {
		t.Errorf("could not connect to test database. make sure the test database has been created and is running. connection string: %v", psqlConnString)
		return nil
	}

	return conn
}

func clearExistingData(db *wsql.PG) {
	// Clear existing database values
	clearOnboardersQuery := "DELETE FROM insys_onboarding.onboarders;"
	_, _ = db.ExecContext(context.Background(), clearOnboardersQuery)

	clearOnboardersLocationQuery := "DELETE FROM insys_onboarding.onboarders_location;"
	_, _ = db.ExecContext(context.Background(), clearOnboardersLocationQuery)

	clearOnboardingTaskInstancesQuery := "DELETE FROM insys_onboarding.onboarding_task_instances;"
	_, _ = db.ExecContext(context.Background(), clearOnboardingTaskInstancesQuery)

	clearOnboardingTasksQuery := "DELETE FROM insys_onboarding.onboarding_tasks;"
	_, _ = db.ExecContext(context.Background(), clearOnboardingTasksQuery)

	clearOnboardingCategoriesQuery := "DELETE FROM insys_onboarding.onboarding_categories;"
	_, _ = db.ExecContext(context.Background(), clearOnboardingCategoriesQuery)
}
