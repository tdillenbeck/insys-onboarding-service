package psql

import (
	"context"
	"io/ioutil"
	"path/filepath"
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"weavelab.xyz/insys-onboarding-service/internal/app"

	"weavelab.xyz/monorail/shared/go-utilities/null"
	"weavelab.xyz/monorail/shared/protorepo/dist/go/enums/insysenums"
	"weavelab.xyz/monorail/shared/wlib/uuid"
	"weavelab.xyz/monorail/shared/wlib/wsql"
)

func TestTaskInstanceService_ByLocationID(t *testing.T) {
	skipCI(t)

	db := initDBConnection(t, psqlConnString)
	clearExistingData(db)

	locationID := uuid.NewV4()

	// create a category
	categoryID := uuid.NewV4()
	query := `
INSERT INTO insys_onboarding.onboarding_categories
(id, display_text, display_order)
VALUES ($1, 'testing display text', 0)
`
	_, err := db.ExecContext(context.Background(), query, categoryID.String())
	if err != nil {
		t.Fatalf("could not create onboarding category: %v\n", err)
	}

	// create a task
	taskID := uuid.NewV4()
	query = `
INSERT INTO insys_onboarding.onboarding_tasks
(id, title, content, display_order, onboarding_category_id)
VALUES ($1, 'testing title', 'testing content', 0, $2)
`
	_, err = db.ExecContext(context.Background(), query, taskID.String(), categoryID.String())
	if err != nil {
		t.Fatalf("could not create onboarding task: %v\n", err)
	}

	// create a task instance
	taskInstanceID := uuid.NewV4()
	query = `
INSERT INTO insys_onboarding.onboarding_task_instances
(id, location_id, title, content, display_order, status, status_updated_at, onboarding_category_id, onboarding_task_id)
VALUES ($1, $2, 'testing title', 'testing content', 0, 0, now(), $3, $4)
`
	_, err = db.ExecContext(context.Background(), query, taskInstanceID.String(), locationID.String(), categoryID.String(), taskID.String())
	if err != nil {
		t.Fatalf("could not create onboarding task instance: %v\n", err)
	}

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
		{
			name:   "finds existing task instances for a location id",
			fields: fields{DB: db},
			args:   args{ctx: context.Background(), locationID: locationID},
			want: []app.TaskInstance{
				{
					LocationID: locationID,
					CategoryID: categoryID,
					TaskID:     taskID,

					ButtonContent:     null.String{},
					ButtonExternalURL: null.String{},
					ButtonInternalURL: null.String{},
					CompletedAt:       null.Time{},
					CompletedBy:       null.String{},
					VerifiedAt:        null.Time{},
					VerifiedBy:        null.String{},
					Content:           "testing content",
					DisplayOrder:      0,
					Status:            0,
					StatusUpdatedAt:   time.Now(),
					StatusUpdatedBy:   null.String{},
					Title:             "testing title",
					Explanation:       null.String{},

					CreatedAt: time.Now(),
					UpdatedAt: time.Now(),
				},
			},
			wantErr: false,
		},
		{
			name:    "attempts to to look up task instances that do not have the requested location id",
			fields:  fields{DB: db},
			args:    args{ctx: context.Background(), locationID: uuid.NewV4()},
			want:    nil,
			wantErr: false,
		},
	}

	// custom functions to ignore fields in cmp.Equal comparison
	opts := []cmp.Option{
		cmpopts.IgnoreFields(app.TaskInstance{}, "ID", "CompletedAt", "VerifiedAt", "StatusUpdatedAt", "CreatedAt", "UpdatedAt"),
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

			if !cmp.Equal(got, tt.want, opts...) {
				t.Errorf("TaskInstanceService.ByLocationID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTaskInstanceService_CreateFromTasks(t *testing.T) {
	skipCI(t)

	db := initDBConnection(t, psqlConnString)
	onboarderService := &OnboarderService{DB: db}
	onboardersLocationService := &OnboardersLocationService{DB: db}

	clearExistingData(db)

	// Setup Database values for test
	absPath, err := filepath.Abs("../../dbconfig/seed.sql")
	if err != nil {
		t.Fatalf("could not find file path for seed file")
	}
	seedFile, err := ioutil.ReadFile(absPath)
	if err != nil {
		t.Fatalf("could not open seed.sql file")
	}
	db.ExecContext(context.Background(), string(seedFile))

	// create an onboarder
	onb := &app.Onboarder{
		UserID:                       uuid.NewV4(),
		ScheduleCustomizationLink:    null.NewString("schedule customization link"),
		SchedulePortingLink:          null.NewString("schedule porting link"),
		ScheduleNetworkLink:          null.NewString("schedule network link"),
		ScheduleSoftwareInstallLink:  null.NewString("schedule software install link"),
		SchedulePhoneInstallLink:     null.NewString("schedule phone install link"),
		ScheduleSoftwareTrainingLink: null.NewString("schedule software training link"),
		SchedulePhoneTrainingLink:    null.NewString("schedule phone training link"),
	}
	onb, err = onboarderService.CreateOrUpdate(context.Background(), onb)
	if err != nil {
		t.Fatalf("could not create onboarder")
	}
	partialOnboarder, err := onboarderService.CreateOrUpdate(
		context.Background(),
		&app.Onboarder{
			UserID:                       uuid.NewV4(),
			ScheduleCustomizationLink:    null.NewString("schedule customization link"),
			SchedulePortingLink:          null.NewString("schedule porting link"),
			ScheduleNetworkLink:          null.NewString("schedule network link"),
			ScheduleSoftwareInstallLink:  null.NewString(""),
			SchedulePhoneInstallLink:     null.NewString(""),
			ScheduleSoftwareTrainingLink: null.NewString("schedule software training link"),
			SchedulePhoneTrainingLink:    null.NewString("schedule phone training link"),
		},
	)
	if err != nil {
		t.Fatalf("could not create onboarder")
	}

	assignedOnboarderLocationID := uuid.NewV4()
	unassignedOnboarderLocationID := uuid.NewV4()
	partialOnboarderLocationID := uuid.NewV4()

	// assign an onboarder to a location
	_, err = onboardersLocationService.CreateOrUpdate(
		context.Background(),
		&app.OnboardersLocation{
			OnboarderID: onb.ID,
			LocationID:  assignedOnboarderLocationID,
		},
	)
	if err != nil {
		t.Fatalf("could not assign onboarder to a location")
	}
	_, err = onboardersLocationService.CreateOrUpdate(context.Background(), &app.OnboardersLocation{
		OnboarderID: partialOnboarder.ID,
		LocationID:  partialOnboarderLocationID,
	})
	if err != nil {
		t.Fatalf("could not assign partial onboarder to a location")
	}

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
		{
			name:   "creates task instances with task data when an onboarder is not assigned",
			fields: fields{DB: db},
			args:   args{ctx: context.Background(), locationID: unassignedOnboarderLocationID},
			want: []app.TaskInstance{
				{
					LocationID: unassignedOnboarderLocationID,
					CategoryID: mustUUID(uuid.Parse("26ba2237-c452-42dd-95ca-a5e59dd2853b")),
					TaskID:     mustUUID(uuid.Parse("16a6dc91-ec6b-4b09-b591-a5b0dfa92932")),

					ButtonContent:     null.NewString("Schedule Call"),
					ButtonExternalURL: null.NewString("https://getweave.chilipiper.com/book/installation-scheduler?type=software-installation"),
					ButtonInternalURL: null.String{},
					CompletedAt:       null.Time{},
					CompletedBy:       null.String{},
					VerifiedAt:        null.Time{},
					VerifiedBy:        null.String{},
					Content:           `<div class="insys-content-body"><span class="insys-content-bold">What is this?</span> On this scheduled call, we will help get your patient or customer database syncing with Weave. This will allow you to start using many of Weave's software features.</div><div class="insys-content-body"><span class="insys-content-bold">How long will this take?</span> 30-60 minutes</div><div class="insys-content-body"><span class="insys-content-bold">Anything else?</span> We may need to access your office server - so make sure you have access to that computer and the right login information.</div>`,
					DisplayOrder:      2,
					Status:            0,
					StatusUpdatedAt:   time.Now(),
					StatusUpdatedBy:   null.NewString("Weave - default"),
					Title:             "Sync your patient data to Weave",
					Explanation:       null.String{},

					CreatedAt: time.Now(),
					UpdatedAt: time.Now(),
				},
				{
					LocationID: unassignedOnboarderLocationID,
					CategoryID: mustUUID(uuid.Parse("26ba2237-c452-42dd-95ca-a5e59dd2853b")),
					TaskID:     mustUUID(uuid.Parse("720af494-38a4-499f-8633-9c8d5169cd43")),

					ButtonContent:     null.NewString("Install Weave"),
					ButtonExternalURL: null.NewString("http://www.weavehelp.com/software-install"),
					ButtonInternalURL: null.String{},
					CompletedAt:       null.Time{},
					CompletedBy:       null.String{},
					VerifiedAt:        null.Time{},
					VerifiedBy:        null.String{},
					Content:           `<div class="insys-content-body"><span class="insys-content-bold">What is this?</span> A quick installation of Weave on other workstations in your office. Installing Weave on more workstations will help you make the most of features like Team Chat and two-way texting.</div><div class="insys-content-body"><span class="insys-content-bold">How long will this take?</span> 5 minutes per workstation</div><div class="insys-content-body"><span class="insys-content-bold">Anything else?</span> Use the link below to download Weave throughout your office - if you need help, just let us know!</div>`,
					DisplayOrder:      5,
					Status:            0,
					StatusUpdatedAt:   time.Now(),
					StatusUpdatedBy:   null.NewString("Weave - default"),
					Title:             "Install Weave on other workstations in your office",
					Explanation:       null.String{},

					CreatedAt: time.Now(),
					UpdatedAt: time.Now(),
				},
				{
					LocationID: unassignedOnboarderLocationID,
					CategoryID: mustUUID(uuid.Parse("26ba2237-c452-42dd-95ca-a5e59dd2853b")),
					TaskID:     mustUUID(uuid.Parse("c20b65d8-e281-4e62-98f0-4aebf83e0bee")),

					ButtonContent:     null.NewString("Watch Videos"),
					ButtonExternalURL: null.NewString("http://www.weavehelp.com/webinar-on-demand/"),
					ButtonInternalURL: null.String{},
					CompletedAt:       null.Time{},
					CompletedBy:       null.String{},
					VerifiedAt:        null.Time{},
					VerifiedBy:        null.String{},
					Content:           `<div class="insys-content-body"><span class="insys-content-bold">What is this?</span> An online webinar - available to schedule or on demand - that will guide you through all the key features you can be using now, like two-way texting, Team Chat, and much more.</div><div class="insys-content-body"><span class="insys-content-bold">How long will this take?</span> 30 minutes</div>`,
					DisplayOrder:      6,
					Status:            0,
					StatusUpdatedAt:   time.Now(),
					StatusUpdatedBy:   null.NewString("Weave - default"),
					Title:             "Watch our helpful software training videos",
					Explanation:       null.String{},

					CreatedAt: time.Now(),
					UpdatedAt: time.Now(),
				},
				{
					LocationID: unassignedOnboarderLocationID,
					CategoryID: mustUUID(uuid.Parse("26ba2237-c452-42dd-95ca-a5e59dd2853b")),
					TaskID:     mustUUID(uuid.Parse("1120842a-a24b-40e8-b29e-0e05e89af99f")),

					ButtonContent:     null.NewString("Learn How"),
					ButtonExternalURL: null.NewString("http://www.weavehelp.com/weave-mobile-app/"),
					ButtonInternalURL: null.String{},
					CompletedAt:       null.Time{},
					CompletedBy:       null.String{},
					VerifiedAt:        null.Time{},
					VerifiedBy:        null.String{},
					Content:           `<div class="insys-content-body"><span class="insys-content-bold">What is this?</span> A quick installation of our mobile app on your phone. Our mobile app is a great tool to see your schedule, field office calls, and chat with your team - all on-the-go, whenever you need access.</div><div class="insys-content-body"><span class="insys-content-bold">How long will this take?</span> 3-5 minutes</div>`,
					DisplayOrder:      7,
					Status:            0,
					StatusUpdatedAt:   time.Now(),
					StatusUpdatedBy:   null.NewString("Weave - default"),
					Title:             "Install the Weave mobile app",
					Explanation:       null.String{},

					CreatedAt: time.Now(),
					UpdatedAt: time.Now(),
				},
				{
					LocationID: unassignedOnboarderLocationID,
					CategoryID: mustUUID(uuid.Parse("ebc72a11-f1b3-40d5-888e-5b6aba66e871")),
					TaskID:     mustUUID(uuid.Parse("7b15e061-8002-4edc-9bf4-f38c6eec6364")),

					ButtonContent:     null.NewString("Schedule Call"),
					ButtonExternalURL: null.NewString("https://app.acuityscheduling.com/schedule.php?owner=14911380&appointmentType=5221365"),
					ButtonInternalURL: null.String{},
					CompletedAt:       null.Time{},
					CompletedBy:       null.String{},
					VerifiedAt:        null.Time{},
					VerifiedBy:        null.String{},
					Content:           `<div class="insys-content-body"><span class="insys-content-bold">What is this?</span> A scheduled call during which a Weave technician will remotely access your workstation to check your office network and make recommendations to have the best experience with Weave.</div><div class="insys-content-body"><span class="insys-content-bold">How long will this take?</span> 15 minutes</div>`,
					DisplayOrder:      3,
					Status:            0,
					StatusUpdatedAt:   time.Now(),
					StatusUpdatedBy:   null.NewString("Weave - default"),
					Title:             "Check your office network to ensure compatibility",
					Explanation:       null.String{},

					CreatedAt: time.Now(),
					UpdatedAt: time.Now(),
				},
				{
					LocationID: unassignedOnboarderLocationID,
					CategoryID: mustUUID(uuid.Parse("ebc72a11-f1b3-40d5-888e-5b6aba66e871")),
					TaskID:     mustUUID(uuid.Parse("fd4f656c-c9f1-47b8-96ad-3080b999a843")),

					ButtonContent:     null.NewString("Schedule Call"),
					ButtonExternalURL: null.NewString("https://getweave.chilipiper.com/book/installation-scheduler?type=phone-installation"),
					ButtonInternalURL: null.String{},
					CompletedAt:       null.Time{},
					CompletedBy:       null.String{},
					VerifiedAt:        null.Time{},
					VerifiedBy:        null.String{},
					Content:           `<div class="insys-content-body"><span class="insys-content-bold">What is this?</span> A scheduled call during which a Weave technician will work with you (or your IT professional) to guide you through physically connecting the new phones to internet and power.</div><div class="insys-content-body"><span class="insys-content-bold">How long will this take?</span> 60 minutes</div><div class="insys-content-body"><span class="insys-content-bold">Anything else?</span> Typically, we will want to help you connect the new phones side-by-side with the old phones while we work with your current phone company for a few days to transition your phones service to Weave.</div>`,
					DisplayOrder:      8,
					Status:            0,
					StatusUpdatedAt:   time.Now(),
					StatusUpdatedBy:   null.NewString("Weave - default"),
					Title:             "Install your new phones",
					Explanation:       null.String{},

					CreatedAt: time.Now(),
					UpdatedAt: time.Now(),
				},
				{
					LocationID: unassignedOnboarderLocationID,
					CategoryID: mustUUID(uuid.Parse("ebc72a11-f1b3-40d5-888e-5b6aba66e871")),
					TaskID:     mustUUID(uuid.Parse("47743fae-c775-45d5-8a51-dc7e3371dfa4")),

					ButtonContent:     null.NewString("Watch Videos"),
					ButtonExternalURL: null.NewString("http://www.weavehelp.com/webinar-on-demand/"),
					ButtonInternalURL: null.String{},
					CompletedAt:       null.Time{},
					CompletedBy:       null.String{},
					VerifiedAt:        null.Time{},
					VerifiedBy:        null.String{},
					Content:           `<div class="insys-content-body"><span class="insys-content-bold">What is this?</span> An online webinar - available to schedule or on demand - that will guide you through all the key features of your new phone system, including voicemail and auto-attendant setup, transferring calls, placing callers on hold, and much more.</div><div class="insys-content-body"><span class="insys-content-bold">How long will this take?</span> 30 minutes</div>`,
					DisplayOrder:      9,
					Status:            0,
					StatusUpdatedAt:   time.Now(),
					StatusUpdatedBy:   null.NewString("Weave - default"),
					Title:             "Watch our helpful phone training videos",
					Explanation:       null.String{},

					CreatedAt: time.Now(),
					UpdatedAt: time.Now(),
				},
				{
					LocationID: unassignedOnboarderLocationID,
					CategoryID: mustUUID(uuid.Parse("ebc72a11-f1b3-40d5-888e-5b6aba66e871")),
					TaskID:     mustUUID(uuid.Parse("2d2df285-9211-48fc-a057-74f7dee2d9a4")),

					ButtonContent:     null.NewString("Schedule Call"),
					ButtonExternalURL: null.NewString("https://getweave.chilipiper.com/book/customization-calls"),
					ButtonInternalURL: null.String{},
					CompletedAt:       null.Time{},
					CompletedBy:       null.String{},
					VerifiedAt:        null.Time{},
					VerifiedBy:        null.String{},
					Content:           `<div class="insys-content-body"><span class="insys-content-bold">What is this?</span> A scheduled call during which your onboarding agent will help customize various features of your phone system - like which phones ring on inbound calls, the name and extension of each phone, and any advanced call routing or auto attendant.</div><div class="insys-content-body"><span class="insys-content-bold">How long will this take?</span> 30 minutes</div>`,
					DisplayOrder:      10,
					Status:            0,
					StatusUpdatedAt:   time.Now(),
					StatusUpdatedBy:   null.NewString("Weave - default"),
					Title:             "Customize your phone system",
					Explanation:       null.String{},

					CreatedAt: time.Now(),
					UpdatedAt: time.Now(),
				},
				{
					LocationID: unassignedOnboarderLocationID,
					CategoryID: mustUUID(uuid.Parse("d0da53a9-fbdb-4d22-85c6-ed521f237349")),
					TaskID:     mustUUID(uuid.Parse("9aec502b-f8b8-4f10-9748-1fe4050eacde")),

					ButtonContent:     null.NewString("Let's Start"),
					ButtonExternalURL: null.NewString("https://getweave.chilipiper.com/book/start-porting-process-call"),
					ButtonInternalURL: null.NewString("/porting"),
					CompletedAt:       null.Time{},
					CompletedBy:       null.String{},
					VerifiedAt:        null.Time{},
					VerifiedBy:        null.String{},
					Content:           `<div class="insys-content-body"> <span class="insys-content-bold"> What is this?  </span> In order to port your phone numbers from your current provider to your new Weave phones, we need three things: <ul> <li>Info about your current phone account</li> <li>Your approval of the terms of service</li> <li>A copy of your current phone bill</li> </ul> </div> <div class="insys-content-body"> <span class="insys-content-bold">How long will this take?</span> 5-10 minutes </div>`,
					DisplayOrder:      4,
					Status:            0,
					StatusUpdatedAt:   time.Now(),
					StatusUpdatedBy:   null.NewString("Weave - default"),
					Title:             "Provide current phone account info",
					Explanation:       null.String{},

					CreatedAt: time.Now(),
					UpdatedAt: time.Now(),
				},
			},
			wantErr: false,
		},
		{
			name:   "creates task instances with onboarder scheduling links when an onboarder is assigned to a location",
			fields: fields{DB: db},
			args:   args{ctx: context.Background(), locationID: assignedOnboarderLocationID},
			want: []app.TaskInstance{
				{
					LocationID: assignedOnboarderLocationID,
					CategoryID: mustUUID(uuid.Parse("26ba2237-c452-42dd-95ca-a5e59dd2853b")),
					TaskID:     mustUUID(uuid.Parse("16a6dc91-ec6b-4b09-b591-a5b0dfa92932")),

					ButtonContent:     null.NewString("Schedule Call"),
					ButtonExternalURL: null.NewString("schedule software install link"),
					ButtonInternalURL: null.String{},
					CompletedAt:       null.Time{},
					CompletedBy:       null.String{},
					VerifiedAt:        null.Time{},
					VerifiedBy:        null.String{},
					Content:           `<div class="insys-content-body"><span class="insys-content-bold">What is this?</span> On this scheduled call, we will help get your patient or customer database syncing with Weave. This will allow you to start using many of Weave's software features.</div><div class="insys-content-body"><span class="insys-content-bold">How long will this take?</span> 30-60 minutes</div><div class="insys-content-body"><span class="insys-content-bold">Anything else?</span> We may need to access your office server - so make sure you have access to that computer and the right login information.</div>`,
					DisplayOrder:      2,
					Status:            0,
					StatusUpdatedAt:   time.Now(),
					StatusUpdatedBy:   null.NewString("Weave - default"),
					Title:             "Sync your patient data to Weave",
					Explanation:       null.String{},

					CreatedAt: time.Now(),
					UpdatedAt: time.Now(),
				},
				{
					LocationID: assignedOnboarderLocationID,
					CategoryID: mustUUID(uuid.Parse("26ba2237-c452-42dd-95ca-a5e59dd2853b")),
					TaskID:     mustUUID(uuid.Parse("720af494-38a4-499f-8633-9c8d5169cd43")),

					ButtonContent:     null.NewString("Install Weave"),
					ButtonExternalURL: null.NewString("http://www.weavehelp.com/software-install"),
					ButtonInternalURL: null.String{},
					CompletedAt:       null.Time{},
					CompletedBy:       null.String{},
					VerifiedAt:        null.Time{},
					VerifiedBy:        null.String{},
					Content:           `<div class="insys-content-body"><span class="insys-content-bold">What is this?</span> A quick installation of Weave on other workstations in your office. Installing Weave on more workstations will help you make the most of features like Team Chat and two-way texting.</div><div class="insys-content-body"><span class="insys-content-bold">How long will this take?</span> 5 minutes per workstation</div><div class="insys-content-body"><span class="insys-content-bold">Anything else?</span> Use the link below to download Weave throughout your office - if you need help, just let us know!</div>`,
					DisplayOrder:      5,
					Status:            0,
					StatusUpdatedAt:   time.Now(),
					StatusUpdatedBy:   null.NewString("Weave - default"),
					Title:             "Install Weave on other workstations in your office",
					Explanation:       null.String{},

					CreatedAt: time.Now(),
					UpdatedAt: time.Now(),
				},
				{
					LocationID: assignedOnboarderLocationID,
					CategoryID: mustUUID(uuid.Parse("26ba2237-c452-42dd-95ca-a5e59dd2853b")),
					TaskID:     mustUUID(uuid.Parse("c20b65d8-e281-4e62-98f0-4aebf83e0bee")),

					ButtonContent:     null.NewString("Watch Videos"),
					ButtonExternalURL: null.NewString("schedule software training link"),
					ButtonInternalURL: null.String{},
					CompletedAt:       null.Time{},
					CompletedBy:       null.String{},
					VerifiedAt:        null.Time{},
					VerifiedBy:        null.String{},
					Content:           `<div class="insys-content-body"><span class="insys-content-bold">What is this?</span> An online webinar - available to schedule or on demand - that will guide you through all the key features you can be using now, like two-way texting, Team Chat, and much more.</div><div class="insys-content-body"><span class="insys-content-bold">How long will this take?</span> 30 minutes</div>`,
					DisplayOrder:      6,
					Status:            0,
					StatusUpdatedAt:   time.Now(),
					StatusUpdatedBy:   null.NewString("Weave - default"),
					Title:             "Watch our helpful software training videos",
					Explanation:       null.String{},

					CreatedAt: time.Now(),
					UpdatedAt: time.Now(),
				},
				{
					LocationID: assignedOnboarderLocationID,
					CategoryID: mustUUID(uuid.Parse("26ba2237-c452-42dd-95ca-a5e59dd2853b")),
					TaskID:     mustUUID(uuid.Parse("1120842a-a24b-40e8-b29e-0e05e89af99f")),

					ButtonContent:     null.NewString("Learn How"),
					ButtonExternalURL: null.NewString("http://www.weavehelp.com/weave-mobile-app/"),
					ButtonInternalURL: null.String{},
					CompletedAt:       null.Time{},
					CompletedBy:       null.String{},
					VerifiedAt:        null.Time{},
					VerifiedBy:        null.String{},
					Content:           `<div class="insys-content-body"><span class="insys-content-bold">What is this?</span> A quick installation of our mobile app on your phone. Our mobile app is a great tool to see your schedule, field office calls, and chat with your team - all on-the-go, whenever you need access.</div><div class="insys-content-body"><span class="insys-content-bold">How long will this take?</span> 3-5 minutes</div>`,
					DisplayOrder:      7,
					Status:            0,
					StatusUpdatedAt:   time.Now(),
					StatusUpdatedBy:   null.NewString("Weave - default"),
					Title:             "Install the Weave mobile app",
					Explanation:       null.String{},

					CreatedAt: time.Now(),
					UpdatedAt: time.Now(),
				},
				{
					LocationID: assignedOnboarderLocationID,
					CategoryID: mustUUID(uuid.Parse("ebc72a11-f1b3-40d5-888e-5b6aba66e871")),
					TaskID:     mustUUID(uuid.Parse("7b15e061-8002-4edc-9bf4-f38c6eec6364")),

					ButtonContent:     null.NewString("Schedule Call"),
					ButtonExternalURL: null.NewString("schedule network link"),
					ButtonInternalURL: null.String{},
					CompletedAt:       null.Time{},
					CompletedBy:       null.String{},
					VerifiedAt:        null.Time{},
					VerifiedBy:        null.String{},
					Content:           `<div class="insys-content-body"><span class="insys-content-bold">What is this?</span> A scheduled call during which a Weave technician will remotely access your workstation to check your office network and make recommendations to have the best experience with Weave.</div><div class="insys-content-body"><span class="insys-content-bold">How long will this take?</span> 15 minutes</div>`,
					DisplayOrder:      3,
					Status:            0,
					StatusUpdatedAt:   time.Now(),
					StatusUpdatedBy:   null.NewString("Weave - default"),
					Title:             "Check your office network to ensure compatibility",
					Explanation:       null.String{},

					CreatedAt: time.Now(),
					UpdatedAt: time.Now(),
				},
				{
					LocationID: assignedOnboarderLocationID,
					CategoryID: mustUUID(uuid.Parse("ebc72a11-f1b3-40d5-888e-5b6aba66e871")),
					TaskID:     mustUUID(uuid.Parse("fd4f656c-c9f1-47b8-96ad-3080b999a843")),

					ButtonContent:     null.NewString("Schedule Call"),
					ButtonExternalURL: null.NewString("schedule phone install link"),
					ButtonInternalURL: null.String{},
					CompletedAt:       null.Time{},
					CompletedBy:       null.String{},
					VerifiedAt:        null.Time{},
					VerifiedBy:        null.String{},
					Content:           `<div class="insys-content-body"><span class="insys-content-bold">What is this?</span> A scheduled call during which a Weave technician will work with you (or your IT professional) to guide you through physically connecting the new phones to internet and power.</div><div class="insys-content-body"><span class="insys-content-bold">How long will this take?</span> 60 minutes</div><div class="insys-content-body"><span class="insys-content-bold">Anything else?</span> Typically, we will want to help you connect the new phones side-by-side with the old phones while we work with your current phone company for a few days to transition your phones service to Weave.</div>`,
					DisplayOrder:      8,
					Status:            0,
					StatusUpdatedAt:   time.Now(),
					StatusUpdatedBy:   null.NewString("Weave - default"),
					Title:             "Install your new phones",
					Explanation:       null.String{},

					CreatedAt: time.Now(),
					UpdatedAt: time.Now(),
				},
				{
					LocationID: assignedOnboarderLocationID,
					CategoryID: mustUUID(uuid.Parse("ebc72a11-f1b3-40d5-888e-5b6aba66e871")),
					TaskID:     mustUUID(uuid.Parse("47743fae-c775-45d5-8a51-dc7e3371dfa4")),

					ButtonContent:     null.NewString("Watch Videos"),
					ButtonExternalURL: null.NewString("schedule phone training link"),
					ButtonInternalURL: null.String{},
					CompletedAt:       null.Time{},
					CompletedBy:       null.String{},
					VerifiedAt:        null.Time{},
					VerifiedBy:        null.String{},
					Content:           `<div class="insys-content-body"><span class="insys-content-bold">What is this?</span> An online webinar - available to schedule or on demand - that will guide you through all the key features of your new phone system, including voicemail and auto-attendant setup, transferring calls, placing callers on hold, and much more.</div><div class="insys-content-body"><span class="insys-content-bold">How long will this take?</span> 30 minutes</div>`,
					DisplayOrder:      9,
					Status:            0,
					StatusUpdatedAt:   time.Now(),
					StatusUpdatedBy:   null.NewString("Weave - default"),
					Title:             "Watch our helpful phone training videos",
					Explanation:       null.String{},

					CreatedAt: time.Now(),
					UpdatedAt: time.Now(),
				},
				{
					LocationID: assignedOnboarderLocationID,
					CategoryID: mustUUID(uuid.Parse("ebc72a11-f1b3-40d5-888e-5b6aba66e871")),
					TaskID:     mustUUID(uuid.Parse("2d2df285-9211-48fc-a057-74f7dee2d9a4")),

					ButtonContent:     null.NewString("Schedule Call"),
					ButtonExternalURL: null.NewString("schedule customization link"),
					ButtonInternalURL: null.String{},
					CompletedAt:       null.Time{},
					CompletedBy:       null.String{},
					VerifiedAt:        null.Time{},
					VerifiedBy:        null.String{},
					Content:           `<div class="insys-content-body"><span class="insys-content-bold">What is this?</span> A scheduled call during which your onboarding agent will help customize various features of your phone system - like which phones ring on inbound calls, the name and extension of each phone, and any advanced call routing or auto attendant.</div><div class="insys-content-body"><span class="insys-content-bold">How long will this take?</span> 30 minutes</div>`,
					DisplayOrder:      10,
					Status:            0,
					StatusUpdatedAt:   time.Now(),
					StatusUpdatedBy:   null.NewString("Weave - default"),
					Title:             "Customize your phone system",
					Explanation:       null.String{},

					CreatedAt: time.Now(),
					UpdatedAt: time.Now(),
				},
				{
					LocationID: assignedOnboarderLocationID,
					CategoryID: mustUUID(uuid.Parse("d0da53a9-fbdb-4d22-85c6-ed521f237349")),
					TaskID:     mustUUID(uuid.Parse("9aec502b-f8b8-4f10-9748-1fe4050eacde")),

					ButtonContent:     null.NewString("Let's Start"),
					ButtonExternalURL: null.NewString("schedule porting link"),
					ButtonInternalURL: null.NewString("/porting"),
					CompletedAt:       null.Time{},
					CompletedBy:       null.String{},
					VerifiedAt:        null.Time{},
					VerifiedBy:        null.String{},
					Content:           `<div class="insys-content-body"> <span class="insys-content-bold"> What is this?  </span> In order to port your phone numbers from your current provider to your new Weave phones, we need three things: <ul> <li>Info about your current phone account</li> <li>Your approval of the terms of service</li> <li>A copy of your current phone bill</li> </ul> </div> <div class="insys-content-body"> <span class="insys-content-bold">How long will this take?</span> 5-10 minutes </div>`,
					DisplayOrder:      4,
					Status:            0,
					StatusUpdatedAt:   time.Now(),
					StatusUpdatedBy:   null.NewString("Weave - default"),
					Title:             "Provide current phone account info",
					Explanation:       null.String{},

					CreatedAt: time.Now(),
					UpdatedAt: time.Now(),
				},
			},
			wantErr: false,
		},
		{
			name:   "create links with an onboarder that has partial links setup",
			fields: fields{DB: db},
			args:   args{ctx: context.Background(), locationID: partialOnboarderLocationID},
			want: []app.TaskInstance{
				{
					LocationID: partialOnboarderLocationID,
					CategoryID: mustUUID(uuid.Parse("26ba2237-c452-42dd-95ca-a5e59dd2853b")),
					TaskID:     mustUUID(uuid.Parse("16a6dc91-ec6b-4b09-b591-a5b0dfa92932")),

					ButtonContent:     null.NewString("Schedule Call"),
					ButtonExternalURL: null.NewString("https://getweave.chilipiper.com/book/installation-scheduler?type=software-installation"),
					ButtonInternalURL: null.String{},
					CompletedAt:       null.Time{},
					CompletedBy:       null.String{},
					VerifiedAt:        null.Time{},
					VerifiedBy:        null.String{},
					Content:           `<div class="insys-content-body"><span class="insys-content-bold">What is this?</span> On this scheduled call, we will help get your patient or customer database syncing with Weave. This will allow you to start using many of Weave's software features.</div><div class="insys-content-body"><span class="insys-content-bold">How long will this take?</span> 30-60 minutes</div><div class="insys-content-body"><span class="insys-content-bold">Anything else?</span> We may need to access your office server - so make sure you have access to that computer and the right login information.</div>`,
					DisplayOrder:      2,
					Status:            0,
					StatusUpdatedAt:   time.Now(),
					StatusUpdatedBy:   null.NewString("Weave - default"),
					Title:             "Sync your patient data to Weave",
					Explanation:       null.String{},

					CreatedAt: time.Now(),
					UpdatedAt: time.Now(),
				},
				{
					LocationID: partialOnboarderLocationID,
					CategoryID: mustUUID(uuid.Parse("26ba2237-c452-42dd-95ca-a5e59dd2853b")),
					TaskID:     mustUUID(uuid.Parse("720af494-38a4-499f-8633-9c8d5169cd43")),

					ButtonContent:     null.NewString("Install Weave"),
					ButtonExternalURL: null.NewString("http://www.weavehelp.com/software-install"),
					ButtonInternalURL: null.String{},
					CompletedAt:       null.Time{},
					CompletedBy:       null.String{},
					VerifiedAt:        null.Time{},
					VerifiedBy:        null.String{},
					Content:           `<div class="insys-content-body"><span class="insys-content-bold">What is this?</span> A quick installation of Weave on other workstations in your office. Installing Weave on more workstations will help you make the most of features like Team Chat and two-way texting.</div><div class="insys-content-body"><span class="insys-content-bold">How long will this take?</span> 5 minutes per workstation</div><div class="insys-content-body"><span class="insys-content-bold">Anything else?</span> Use the link below to download Weave throughout your office - if you need help, just let us know!</div>`,
					DisplayOrder:      5,
					Status:            0,
					StatusUpdatedAt:   time.Now(),
					StatusUpdatedBy:   null.NewString("Weave - default"),
					Title:             "Install Weave on other workstations in your office",
					Explanation:       null.String{},

					CreatedAt: time.Now(),
					UpdatedAt: time.Now(),
				},
				{
					LocationID: partialOnboarderLocationID,
					CategoryID: mustUUID(uuid.Parse("26ba2237-c452-42dd-95ca-a5e59dd2853b")),
					TaskID:     mustUUID(uuid.Parse("c20b65d8-e281-4e62-98f0-4aebf83e0bee")),

					ButtonContent:     null.NewString("Watch Videos"),
					ButtonExternalURL: null.NewString("schedule software training link"),
					ButtonInternalURL: null.String{},
					CompletedAt:       null.Time{},
					CompletedBy:       null.String{},
					VerifiedAt:        null.Time{},
					VerifiedBy:        null.String{},
					Content:           `<div class="insys-content-body"><span class="insys-content-bold">What is this?</span> An online webinar - available to schedule or on demand - that will guide you through all the key features you can be using now, like two-way texting, Team Chat, and much more.</div><div class="insys-content-body"><span class="insys-content-bold">How long will this take?</span> 30 minutes</div>`,
					DisplayOrder:      6,
					Status:            0,
					StatusUpdatedAt:   time.Now(),
					StatusUpdatedBy:   null.NewString("Weave - default"),
					Title:             "Watch our helpful software training videos",
					Explanation:       null.String{},

					CreatedAt: time.Now(),
					UpdatedAt: time.Now(),
				},
				{
					LocationID: partialOnboarderLocationID,
					CategoryID: mustUUID(uuid.Parse("26ba2237-c452-42dd-95ca-a5e59dd2853b")),
					TaskID:     mustUUID(uuid.Parse("1120842a-a24b-40e8-b29e-0e05e89af99f")),

					ButtonContent:     null.NewString("Learn How"),
					ButtonExternalURL: null.NewString("http://www.weavehelp.com/weave-mobile-app/"),
					ButtonInternalURL: null.String{},
					CompletedAt:       null.Time{},
					CompletedBy:       null.String{},
					VerifiedAt:        null.Time{},
					VerifiedBy:        null.String{},
					Content:           `<div class="insys-content-body"><span class="insys-content-bold">What is this?</span> A quick installation of our mobile app on your phone. Our mobile app is a great tool to see your schedule, field office calls, and chat with your team - all on-the-go, whenever you need access.</div><div class="insys-content-body"><span class="insys-content-bold">How long will this take?</span> 3-5 minutes</div>`,
					DisplayOrder:      7,
					Status:            0,
					StatusUpdatedAt:   time.Now(),
					StatusUpdatedBy:   null.NewString("Weave - default"),
					Title:             "Install the Weave mobile app",
					Explanation:       null.String{},

					CreatedAt: time.Now(),
					UpdatedAt: time.Now(),
				},
				{
					LocationID: partialOnboarderLocationID,
					CategoryID: mustUUID(uuid.Parse("ebc72a11-f1b3-40d5-888e-5b6aba66e871")),
					TaskID:     mustUUID(uuid.Parse("7b15e061-8002-4edc-9bf4-f38c6eec6364")),

					ButtonContent:     null.NewString("Schedule Call"),
					ButtonExternalURL: null.NewString("schedule network link"),
					ButtonInternalURL: null.String{},
					CompletedAt:       null.Time{},
					CompletedBy:       null.String{},
					VerifiedAt:        null.Time{},
					VerifiedBy:        null.String{},
					Content:           `<div class="insys-content-body"><span class="insys-content-bold">What is this?</span> A scheduled call during which a Weave technician will remotely access your workstation to check your office network and make recommendations to have the best experience with Weave.</div><div class="insys-content-body"><span class="insys-content-bold">How long will this take?</span> 15 minutes</div>`,
					DisplayOrder:      3,
					Status:            0,
					StatusUpdatedAt:   time.Now(),
					StatusUpdatedBy:   null.NewString("Weave - default"),
					Title:             "Check your office network to ensure compatibility",
					Explanation:       null.String{},

					CreatedAt: time.Now(),
					UpdatedAt: time.Now(),
				},
				{
					LocationID: partialOnboarderLocationID,
					CategoryID: mustUUID(uuid.Parse("ebc72a11-f1b3-40d5-888e-5b6aba66e871")),
					TaskID:     mustUUID(uuid.Parse("fd4f656c-c9f1-47b8-96ad-3080b999a843")),

					ButtonContent:     null.NewString("Schedule Call"),
					ButtonExternalURL: null.NewString("https://getweave.chilipiper.com/book/installation-scheduler?type=phone-installation"),
					ButtonInternalURL: null.String{},
					CompletedAt:       null.Time{},
					CompletedBy:       null.String{},
					VerifiedAt:        null.Time{},
					VerifiedBy:        null.String{},
					Content:           `<div class="insys-content-body"><span class="insys-content-bold">What is this?</span> A scheduled call during which a Weave technician will work with you (or your IT professional) to guide you through physically connecting the new phones to internet and power.</div><div class="insys-content-body"><span class="insys-content-bold">How long will this take?</span> 60 minutes</div><div class="insys-content-body"><span class="insys-content-bold">Anything else?</span> Typically, we will want to help you connect the new phones side-by-side with the old phones while we work with your current phone company for a few days to transition your phones service to Weave.</div>`,
					DisplayOrder:      8,
					Status:            0,
					StatusUpdatedAt:   time.Now(),
					StatusUpdatedBy:   null.NewString("Weave - default"),
					Title:             "Install your new phones",
					Explanation:       null.String{},

					CreatedAt: time.Now(),
					UpdatedAt: time.Now(),
				},
				{
					LocationID: partialOnboarderLocationID,
					CategoryID: mustUUID(uuid.Parse("ebc72a11-f1b3-40d5-888e-5b6aba66e871")),
					TaskID:     mustUUID(uuid.Parse("47743fae-c775-45d5-8a51-dc7e3371dfa4")),

					ButtonContent:     null.NewString("Watch Videos"),
					ButtonExternalURL: null.NewString("schedule phone training link"),
					ButtonInternalURL: null.String{},
					CompletedAt:       null.Time{},
					CompletedBy:       null.String{},
					VerifiedAt:        null.Time{},
					VerifiedBy:        null.String{},
					Content:           `<div class="insys-content-body"><span class="insys-content-bold">What is this?</span> An online webinar - available to schedule or on demand - that will guide you through all the key features of your new phone system, including voicemail and auto-attendant setup, transferring calls, placing callers on hold, and much more.</div><div class="insys-content-body"><span class="insys-content-bold">How long will this take?</span> 30 minutes</div>`,
					DisplayOrder:      9,
					Status:            0,
					StatusUpdatedAt:   time.Now(),
					StatusUpdatedBy:   null.NewString("Weave - default"),
					Title:             "Watch our helpful phone training videos",
					Explanation:       null.String{},

					CreatedAt: time.Now(),
					UpdatedAt: time.Now(),
				},
				{
					LocationID: partialOnboarderLocationID,
					CategoryID: mustUUID(uuid.Parse("ebc72a11-f1b3-40d5-888e-5b6aba66e871")),
					TaskID:     mustUUID(uuid.Parse("2d2df285-9211-48fc-a057-74f7dee2d9a4")),

					ButtonContent:     null.NewString("Schedule Call"),
					ButtonExternalURL: null.NewString("schedule customization link"),
					ButtonInternalURL: null.String{},
					CompletedAt:       null.Time{},
					CompletedBy:       null.String{},
					VerifiedAt:        null.Time{},
					VerifiedBy:        null.String{},
					Content:           `<div class="insys-content-body"><span class="insys-content-bold">What is this?</span> A scheduled call during which your onboarding agent will help customize various features of your phone system - like which phones ring on inbound calls, the name and extension of each phone, and any advanced call routing or auto attendant.</div><div class="insys-content-body"><span class="insys-content-bold">How long will this take?</span> 30 minutes</div>`,
					DisplayOrder:      10,
					Status:            0,
					StatusUpdatedAt:   time.Now(),
					StatusUpdatedBy:   null.NewString("Weave - default"),
					Title:             "Customize your phone system",
					Explanation:       null.String{},

					CreatedAt: time.Now(),
					UpdatedAt: time.Now(),
				},
				{
					LocationID: partialOnboarderLocationID,
					CategoryID: mustUUID(uuid.Parse("d0da53a9-fbdb-4d22-85c6-ed521f237349")),
					TaskID:     mustUUID(uuid.Parse("9aec502b-f8b8-4f10-9748-1fe4050eacde")),

					ButtonContent:     null.NewString("Let's Start"),
					ButtonExternalURL: null.NewString("schedule porting link"),
					ButtonInternalURL: null.NewString("/porting"),
					CompletedAt:       null.Time{},
					CompletedBy:       null.String{},
					VerifiedAt:        null.Time{},
					VerifiedBy:        null.String{},
					Content:           `<div class="insys-content-body"> <span class="insys-content-bold"> What is this?  </span> In order to port your phone numbers from your current provider to your new Weave phones, we need three things: <ul> <li>Info about your current phone account</li> <li>Your approval of the terms of service</li> <li>A copy of your current phone bill</li> </ul> </div> <div class="insys-content-body"> <span class="insys-content-bold">How long will this take?</span> 5-10 minutes </div>`,
					DisplayOrder:      4,
					Status:            0,
					StatusUpdatedAt:   time.Now(),
					StatusUpdatedBy:   null.NewString("Weave - default"),
					Title:             "Provide current phone account info",
					Explanation:       null.String{},

					CreatedAt: time.Now(),
					UpdatedAt: time.Now(),
				},
			},
			wantErr: false,
		},
	}

	opts := []cmp.Option{
		cmpopts.IgnoreFields(app.TaskInstance{}, "ID", "CompletedAt", "VerifiedAt", "StatusUpdatedAt", "CreatedAt", "UpdatedAt"),
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

			if !cmp.Equal(got, tt.want, opts...) {
				t.Errorf("TaskInstanceService.CreateFromTasks() compare failed. diff: %v", cmp.Diff(got, tt.want, opts...))
			}
		})
	}
}

func TestTaskInstanceService_Update(t *testing.T) {
	skipCI(t)

	db := initDBConnection(t, psqlConnString)
	clearExistingData(db)

	locationID := uuid.NewV4()

	// create a category
	categoryID := uuid.NewV4()
	query := `
INSERT INTO insys_onboarding.onboarding_categories
(id, display_text, display_order)
VALUES ($1, 'testing display text', 0)
`
	_, err := db.ExecContext(context.Background(), query, categoryID.String())
	if err != nil {
		t.Fatalf("could not create onboarding category: %v\n", err)
	}

	// create a task
	taskID := uuid.NewV4()
	query = `
INSERT INTO insys_onboarding.onboarding_tasks
(id, title, content, display_order, onboarding_category_id)
VALUES ($1, 'testing title', 'testing content', 0, $2)
`
	_, err = db.ExecContext(context.Background(), query, taskID.String(), categoryID.String())
	if err != nil {
		t.Fatalf("could not create onboarding task: %v\n", err)
	}

	// create a task instance
	taskInstanceID := uuid.NewV4()
	query = `
INSERT INTO insys_onboarding.onboarding_task_instances
(id, location_id, title, content, display_order, status, status_updated_at, onboarding_category_id, onboarding_task_id)
VALUES ($1, $2, 'testing title', 'testing content', 0, 0, now(), $3, $4)
`
	_, err = db.ExecContext(context.Background(), query, taskInstanceID.String(), locationID.String(), categoryID.String(), taskID.String())
	if err != nil {
		t.Fatalf("could not create onboarding task instance: %v\n", err)
	}

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
		{
			name:   "update the status to completed",
			fields: fields{DB: db},
			args: args{
				ctx:             context.Background(),
				id:              taskInstanceID,
				status:          2,
				statusUpdatedBy: "test",
			},
			want: &app.TaskInstance{
				ID:         taskInstanceID,
				LocationID: locationID,
				CategoryID: categoryID,
				TaskID:     taskID,

				ButtonContent:     null.String{},
				ButtonExternalURL: null.String{},
				ButtonInternalURL: null.String{},
				CompletedAt:       null.Time{},
				CompletedBy:       null.NewString("test"),
				VerifiedAt:        null.Time{},
				VerifiedBy:        null.String{},
				Content:           "testing content",
				DisplayOrder:      0,
				Status:            2,
				StatusUpdatedAt:   time.Now(),
				StatusUpdatedBy:   null.NewString("test"),
				Title:             "testing title",
				Explanation:       null.String{},

				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
			},
			wantErr: false,
		},
	}

	opts := []cmp.Option{
		cmpopts.IgnoreFields(app.TaskInstance{}, "ID", "CompletedAt", "VerifiedAt", "StatusUpdatedAt", "CreatedAt", "UpdatedAt"),
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

			if !cmp.Equal(*got, *tt.want, opts...) {
				t.Errorf("TaskInstanceService.Update() = %v, want %v", got, tt.want)
			}
		})
	}
}

// musUUID is a helper function that simply returns the value without checking the errors. This is usefule when using uuid.Parse directly in a struct definition.
func mustUUID(u uuid.UUID, err error) uuid.UUID {
	return u
}
