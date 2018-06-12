package psql

import (
	"context"
	"reflect"
	"testing"

	app "weavelab.xyz/insys-onboarding"
	"weavelab.xyz/wlib/uuid"
	"weavelab.xyz/wlib/wsql"
)

func TestCategoryService_ByID(t *testing.T) {
	t.Skip("Not implemented yet")

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
	// TODO: Add test cases.
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
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CategoryService.ByID() = %v, want %v", got, tt.want)
			}
		})
	}
}
