package psql

import (
	"context"
	"os"
	"testing"
	"time"

	"weavelab.xyz/monorail/shared/wlib/wsql"
)

const (
	psqlConnString = "postgresql://127.0.0.1:5432/insys_onboarding_test?sslmode=disable"
)

func initDBConnection(t *testing.T) *wsql.PG {
	connString, exists := os.LookupEnv("PG_PRIMARY_CONNECT_STRING")
	if !exists {
		connString = psqlConnString
	}

	dbOptions := &ConnectionOptions{
		MaxOpenConnections:    10,
		MaxIdleConnections:    2,
		MaxConnectionLifetime: 5 * time.Minute,
		LogQueries:            false,
	}

	conn, err := ConnectionFromConnString(context.Background(), connString, connString, dbOptions)
	if err != nil {
		t.Fatalf("could not connect to test database. make sure the test database has been created and is running. connection string: %v", psqlConnString)
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

	clearChiliPiperScheduleEventsQuery := "DELETE FROM insys_onboarding.chili_piper_schedule_events;"
	_, _ = db.ExecContext(context.Background(), clearChiliPiperScheduleEventsQuery)

	clearHandOffSnapshotsQuery := "DELETE FROM insys_onboarding.handoff_snapshots;"
	_, _ = db.ExecContext(context.Background(), clearHandOffSnapshotsQuery)
}
