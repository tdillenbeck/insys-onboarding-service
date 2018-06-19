package psql

import (
	"context"
	"reflect"
	"testing"

	"weavelab.xyz/insys-onboarding/internal/app"
	"weavelab.xyz/protorepo/dist/go/enums/insysenums"
	"weavelab.xyz/wlib/uuid"
	"weavelab.xyz/wlib/wsql"
)

func TestTaskInstanceService_ByLocationID(t *testing.T) {
	t.Skip("Not implemented yet")

	type fields struct {
		DB *wsql.PG
	}
	type args struct {
		ctx        context.Context
		locationID uuid.UUID
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []app.TaskInstance
		wantErr bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tis := &TaskInstanceService{
				DB: tt.fields.DB,
			}
			got, err := tis.ByLocationID(tt.args.ctx, tt.args.locationID)
			if (err != nil) != tt.wantErr {
				t.Errorf("TaskInstanceService.ByLocationID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TaskInstanceService.ByLocationID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTaskInstanceService_CreateFromTasks(t *testing.T) {
	t.Skip("Not implemented yet")

	type fields struct {
		DB *wsql.PG
	}
	type args struct {
		ctx        context.Context
		locationID uuid.UUID
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []app.TaskInstance
		wantErr bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tis := &TaskInstanceService{
				DB: tt.fields.DB,
			}
			got, err := tis.CreateFromTasks(tt.args.ctx, tt.args.locationID)
			if (err != nil) != tt.wantErr {
				t.Errorf("TaskInstanceService.CreateFromTasks() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TaskInstanceService.CreateFromTasks() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTaskInstanceService_Update(t *testing.T) {
	t.Skip("Not implemented yet")

	type fields struct {
		DB *wsql.PG
	}
	type args struct {
		ctx             context.Context
		id              uuid.UUID
		status          insysenums.OnboardingTaskStatus
		statusUpdatedBy string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *app.TaskInstance
		wantErr bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tis := &TaskInstanceService{
				DB: tt.fields.DB,
			}
			got, err := tis.Update(tt.args.ctx, tt.args.id, tt.args.status, tt.args.statusUpdatedBy)
			if (err != nil) != tt.wantErr {
				t.Errorf("TaskInstanceService.Update() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TaskInstanceService.Update() = %v, want %v", got, tt.want)
			}
		})
	}
}
