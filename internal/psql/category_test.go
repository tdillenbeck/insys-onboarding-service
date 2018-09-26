package psql

import (
	"context"
	"io/ioutil"
	"path/filepath"
	"testing"
	"time"

	"weavelab.xyz/insys-onboarding/internal/app"
	"weavelab.xyz/wlib/uuid"
	"weavelab.xyz/wlib/wsql"
)

func TestCategoryService_ByID(t *testing.T) {
	skipCI(t)
	db := initDBConnection(t, psqlConnString)

	expectedTime := time.Date(1987, 10, 2, 0, 0, 0, 0, time.UTC)
	categoryUUID, err := uuid.Parse("26ba2237-c452-42dd-95ca-a5e59dd2853b")
	if err != nil {
		t.Fatalf("could not parse categoryUUID")
	}

	// Setup Database values for test
	absPath, err := filepath.Abs("../../dbconfig/seed.sql")
	if err != nil {
		t.Fatalf("could not find file path for seed file")
	}
	seedFile, err := ioutil.ReadFile(absPath)
	if err != nil {
		t.Fatalf("could not open seed.sql file")
	}
	_, err = db.ExecContext(context.Background(), string(seedFile))
	if err != nil {
		t.Fatalf("could not insert seed data into the db")
	}

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
			if !compareCategory(got, tt.want) {
				t.Errorf("CategoryService.ByID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func compareCategory(a, b *app.Category) bool {
	// handle nil case
	if a == b {
		return true
	}
	if a == nil || b == nil {
		return false
	}

	return a.ID == b.ID &&
		a.DisplayText == b.DisplayText &&
		a.DisplayOrder == b.DisplayOrder
}
