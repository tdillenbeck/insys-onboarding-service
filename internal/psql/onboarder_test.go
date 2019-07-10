package psql

import (
	"context"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"weavelab.xyz/insys-onboarding-service/internal/app"
	"weavelab.xyz/monorail/shared/go-utilities/null"
	"weavelab.xyz/monorail/shared/wlib/uuid"
	"weavelab.xyz/monorail/shared/wlib/wsql"
)

func TestOnboarderService_CreateOrUpdate(t *testing.T) {
	db := initDBConnection(t, psqlConnString)
	clearExistingData(db)

	userID := uuid.NewV4()
	id := uuid.NewV4()
	existingUserID := uuid.NewV4()

	//create existing onoarder to test update functionality
	query := `INSERT INTO insys_onboarding.onboarders
		(id, user_id, schedule_customization_link)
		VALUES ($1, $2, 'testing exsinting_schedule_customization_link')`
	_, err := db.ExecContext(context.Background(), query, id, existingUserID)
	if err != nil {
		t.Fatalf("could not create onboarder: %v\n", err)
	}

	type fields struct {
		DB *wsql.PG
	}
	type args struct {
		ctx context.Context
		onb *app.Onboarder
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *app.Onboarder
		wantErr bool
	}{
		{
			name:   "successfully creates an onboarder",
			fields: fields{DB: db},
			args: args{
				ctx: context.Background(),
				onb: &app.Onboarder{
					SalesforceUserID:             null.NewString("testing salesforce user id"),
					ScheduleCustomizationLink:    null.NewString("testing schedule customization link"),
					ScheduleNetworkLink:          null.NewString("testing schedule network link"),
					SchedulePhoneInstallLink:     null.NewString("testing schedule phone install link"),
					SchedulePhoneTrainingLink:    null.NewString("testing schedule phone training link"),
					SchedulePortingLink:          null.NewString("testing schedule porting link"),
					ScheduleSoftwareInstallLink:  null.NewString("testing software install link"),
					ScheduleSoftwareTrainingLink: null.NewString("testing software training link"),
					UserID:                       userID,
				},
			},
			want: &app.Onboarder{
				SalesforceUserID:             null.NewString("testing salesforce user id"),
				ScheduleCustomizationLink:    null.NewString("testing schedule customization link"),
				ScheduleNetworkLink:          null.NewString("testing schedule network link"),
				SchedulePhoneInstallLink:     null.NewString("testing schedule phone install link"),
				SchedulePhoneTrainingLink:    null.NewString("testing schedule phone training link"),
				SchedulePortingLink:          null.NewString("testing schedule porting link"),
				ScheduleSoftwareInstallLink:  null.NewString("testing software install link"),
				ScheduleSoftwareTrainingLink: null.NewString("testing software training link"),
				UserID:                       userID,
			},
			wantErr: false,
		},
		{
			name:   "successfully updates an existing onboarder",
			fields: fields{DB: db},
			args: args{
				ctx: context.Background(),
				onb: &app.Onboarder{
					SalesforceUserID:             null.NewString("testing salesforce user id"),
					ScheduleCustomizationLink:    null.NewString("testing schedule customization link"),
					ScheduleNetworkLink:          null.NewString("testing schedule network link"),
					SchedulePhoneInstallLink:     null.NewString("testing schedule phone install link"),
					SchedulePhoneTrainingLink:    null.NewString("testing schedule phone training link"),
					SchedulePortingLink:          null.NewString("testing schedule porting link"),
					ScheduleSoftwareInstallLink:  null.NewString("testing software install link"),
					ScheduleSoftwareTrainingLink: null.NewString("testing software training link"),
					UserID:                       existingUserID,
				},
			},
			want: &app.Onboarder{
				SalesforceUserID:             null.NewString("testing salesforce user id"),
				ScheduleCustomizationLink:    null.NewString("testing schedule customization link"),
				ScheduleNetworkLink:          null.NewString("testing schedule network link"),
				SchedulePhoneInstallLink:     null.NewString("testing schedule phone install link"),
				SchedulePhoneTrainingLink:    null.NewString("testing schedule phone training link"),
				SchedulePortingLink:          null.NewString("testing schedule porting link"),
				ScheduleSoftwareInstallLink:  null.NewString("testing software install link"),
				ScheduleSoftwareTrainingLink: null.NewString("testing software training link"),
				UserID:                       existingUserID,
			},
			wantErr: false,
		},
	}

	// custom functions to ignore fields in cmp.Equal comparison
	opts := []cmp.Option{
		cmpopts.IgnoreFields(app.Onboarder{}, "ID", "CreatedAt", "UpdatedAt"),
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &OnboarderService{
				DB: tt.fields.DB,
			}
			got, err := s.CreateOrUpdate(tt.args.ctx, tt.args.onb)
			if (err != nil) != tt.wantErr {
				t.Errorf("OnboarderService.CreateOrUpdate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !cmp.Equal(got, tt.want, opts...) {
				t.Errorf("OnboarderService.CreateOrUpdate().\n %v", cmp.Diff(got, tt.want, opts...))
			}
		})
	}
}

func TestOnboarderService_ReadByUserID(t *testing.T) {
	db := initDBConnection(t, psqlConnString)
	clearExistingData(db)

	id := uuid.NewV4()
	existingUserID := uuid.NewV4()

	query := `INSERT INTO insys_onboarding.onboarders
		(id, user_id, schedule_customization_link, salesforce_user_id)
		VALUES ($1, $2, 'testing existing_schedule_customization_link', 'testing salesforce_user_id')`
	_, err := db.ExecContext(context.Background(), query, id, existingUserID)
	if err != nil {
		t.Fatalf("could not create onboarder: %v\n", err)
	}

	type fields struct {
		DB *wsql.PG
	}
	type args struct {
		ctx    context.Context
		userID uuid.UUID
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *app.Onboarder
		wantErr bool
	}{
		{
			name:   "successfully reads an existing onboarder",
			fields: fields{DB: db},
			args:   args{context.Background(), existingUserID},
			want: &app.Onboarder{
				UserID:                    existingUserID,
				SalesforceUserID:          null.NewString("testing salesforce_user_id"),
				ScheduleCustomizationLink: null.NewString("testing existing_schedule_customization_link"),
			},
			wantErr: false,
		},
	}

	// custom functions to ignore fields in cmp.Equal comparison
	opts := []cmp.Option{
		cmpopts.IgnoreFields(app.Onboarder{}, "ID", "CreatedAt", "UpdatedAt"),
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &OnboarderService{
				DB: tt.fields.DB,
			}
			got, err := s.ReadByUserID(tt.args.ctx, tt.args.userID)
			if (err != nil) != tt.wantErr {
				t.Errorf("OnboarderService.ReadByUserID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !cmp.Equal(got, tt.want, opts...) {
				t.Errorf("OnboarderService.ReadByUserID() = %v, want %v", got, tt.want)
			}
		})
	}
}
