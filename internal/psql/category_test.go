package psql

import (
	"context"
	"os"
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
	"weavelab.xyz/insys-onboarding/internal/app"
	"weavelab.xyz/wlib/uuid"
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

func TestCategoryService_ByID(t *testing.T) {
	skipCI(t)
	db := initDBConnection(t, psqlConnString)

	expectedTime := time.Date(1987, 10, 2, 0, 0, 0, 0, time.UTC)
	categoryUUID, err := uuid.Parse("26ba2237-c452-42dd-95ca-a5e59dd2853b")
	if err != nil {
		t.Fatalf("could not parse categoryUUID")
	}

	// Setup Database values for test
	clearQuery := "DELETE FROM insys_onboarding.onboarding_categories;"
	db.QueryRowContext(context.Background(), clearQuery)
	insertQuery := "INSERT INTO insys_onboarding.onboarding_categories VALUES ('26ba2237-c452-42dd-95ca-a5e59dd2853b', 'Software', 1, date ('1987-10-02 00:00:00'), date ('1987-10-02 00:00:00'));"
	db.QueryRowContext(context.Background(), insertQuery)

	type fields struct {
		DB *wsql.PG
	}
	type args struct {
		ctx context.Context
		id  uuid.UUID
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *app.Category
		wantErr bool
	}{
		{
			"retrieve existing category",
			fields{DB: db},
			args{
				ctx: context.Background(),
				id:  categoryUUID,
			},
			&app.Category{
				ID:           categoryUUID,
				DisplayText:  "Software",
				DisplayOrder: 1,
				CreatedAt:    expectedTime,
				UpdatedAt:    expectedTime,
			},
			false,
		},
		{
			"attemp toretrieve a category that doesn't exist",
			fields{DB: db},
			args{
				ctx: context.Background(),
				id:  uuid.NewV4(),
			},
			nil,
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &CategoryService{
				DB: tt.fields.DB,
			}
			got, err := s.ByID(tt.args.ctx, tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("CategoryService.ByID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !cmp.Equal(got, tt.want) {
				t.Errorf("CategoryService.ByID() = %v, want %v", got, tt.want)
			}
		})
	}
}
