package psql

import (
	"context"
	"io/ioutil"
	"path/filepath"
	"reflect"
	"sort"
	"testing"
	"time"

	"weavelab.xyz/go-utilities/null"
	"weavelab.xyz/insys-onboarding/internal/app"
	"weavelab.xyz/protorepo/dist/go/enums/insysenums"
	"weavelab.xyz/wlib/uuid"
	"weavelab.xyz/wlib/wsql"
)

func TestTaskInstanceService_ByLocationID(t *testing.T) {
	skipCI(t)

	db := initDBConnection(t, psqlConnString)
	clearExistingData(t, db)

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
			want:    []app.TaskInstance{},
			wantErr: false,
		},
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
			if !compareTaskInstances(got, tt.want) {
				t.Errorf("TaskInstanceService.CreateFromTasks() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTaskInstanceService_CreateFromTasks(t *testing.T) {
	skipCI(t)

	db := initDBConnection(t, psqlConnString)
	onboarderService := &OnboarderService{DB: db}
	onboardersLocationService := &OnboardersLocationService{DB: db}

	clearExistingData(t, db)

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
		t.Fatalf("could not insert seed data into database")
	}

	// create an onboarder
	onb := &app.Onboarder{
		UserID: uuid.NewV4(),
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

	assignedOnboarderLocationID := uuid.NewV4()
	unassignedOnboarderLocationID := uuid.NewV4()

	// assign an onboarder to a location
	onbl := &app.OnboardersLocation{
		OnboarderID: onb.ID,
		LocationID:  assignedOnboarderLocationID,
	}
	_, err = onboardersLocationService.CreateOrUpdate(context.Background(), onbl)
	if err != nil {
		t.Fatalf("could not assign onboarder to a location")
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
					ButtonExternalURL: null.NewString("https://app.acuityscheduling.com/schedule.php?owner=14911380&appointmentType=5221365"),
					CompletedAt:       null.Time{},
					CompletedBy:       null.String{},
					VerifiedAt:        null.Time{},
					VerifiedBy:        null.String{},
					Content:           "In order to start using all our software features, we need to sync your patient data to Weave. Typically, we will do this by installing a sync application on your office server - so make sure you have access to that computer and the right login information. If your practice management system is cloud-based, we will instead walk you through how to generate the credentials we need. Click below to schedule a 30-60 minute phone call with a Weave technician.",
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
					CompletedAt:       null.Time{},
					CompletedBy:       null.String{},
					VerifiedAt:        null.Time{},
					VerifiedBy:        null.String{},
					Content:           "In order to make the most of features like Team Chat and two-way text messaging, we recommend installing Weave on most workstations throughout your office. Click below to find helpful instructions to download Weave throughout your office.",
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
					ButtonExternalURL: null.NewString("http://www.weavehelp.com/webinar-registration-page/"),
					CompletedAt:       null.Time{},
					CompletedBy:       null.String{},
					VerifiedAt:        null.Time{},
					VerifiedBy:        null.String{},
					Content:           "In order to get the most out of the Weave software, we recommend tuning in to our training videos online. Our helpful training staff will walk you through setting up and using two-way texting, automated appointment reminders, and all our other features. Click below to watch our online webinars",
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
					CompletedAt:       null.Time{},
					CompletedBy:       null.String{},
					VerifiedAt:        null.Time{},
					VerifiedBy:        null.String{},
					Content:           "In order to see your schedule, field calls, and chat with your team all on the go, we recommend installing the Weave app on your mobile device. Click below to find helpful instructions to download Weave on your phone.",
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
					CompletedAt:       null.Time{},
					CompletedBy:       null.String{},
					VerifiedAt:        null.Time{},
					VerifiedBy:        null.String{},
					Content:           `In order to ensure the highest quality phone service on Weave's internet-based phones, we need to verify your network compatibility. Click below to schedule a 15 minute phone call with a Weave technician. <a href="http://www.weavehelp.com/network-specs">Learn More</a>`,
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
					ButtonExternalURL: null.NewString("https://app.acuityscheduling.com/schedule.php?owner=14911380&appointmentType=category:Installs"),
					CompletedAt:       null.Time{},
					CompletedBy:       null.String{},
					VerifiedAt:        null.Time{},
					VerifiedBy:        null.String{},
					Content:           `You're making great progress! Click below to schedule a 60 minute phone call with a Weave technician who will guide you through the process of connecting the new phones. Typically, we will aim to help you plug the new phones in side-by-side with your old phones. After finishing, you will continue to use your old phones while we work with your current phone service provider for a few days to officially move your phone numbers to our system. <a href="http://www.weavehelp.com/phone-install">Learn more.</a>`,
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
					ButtonExternalURL: null.NewString("http://www.weavehelp.com/webinar-registration-page/"),
					CompletedAt:       null.Time{},
					CompletedBy:       null.String{},
					VerifiedAt:        null.Time{},
					VerifiedBy:        null.String{},
					Content:           `In order to get the most out of the Weave phones, we recommend tuning in to our training videos online. Our helpful training staff will walk you through setting up voicemail messages, placing callers on hold, transferring calls, and all our other phone features. Click below to watch our online webinars.`,
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
					CompletedAt:       null.Time{},
					CompletedBy:       null.String{},
					VerifiedAt:        null.Time{},
					VerifiedBy:        null.String{},
					Content:           `In order to have the best experience with your new phones, we recommend customizing the system to suit your needs. Your onboarding agent will help you change which phones ring when a call comes in, adjust the names and extensions on each phone, and set up any advanced call routing or auto attendant you may need. Click below to schedule a call.`,
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

					ButtonContent:     null.NewString("Schedule Call"),
					ButtonExternalURL: null.NewString("https://getweave.chilipiper.com/book/start-porting-process-call"),
					CompletedAt:       null.Time{},
					CompletedBy:       null.String{},
					VerifiedAt:        null.Time{},
					VerifiedBy:        null.String{},
					Content:           `One of your first steps is to work with your onboarding agent to verify your office phone numbers. Your onboarding agent will then handle some initial processing with your current phone service provider to ensure that when your office is ready to start using the new phones, we are ready to activate them for you. Click below to schedule a call with your onboarding agent.`,
					DisplayOrder:      4,
					Status:            0,
					StatusUpdatedAt:   time.Now(),
					StatusUpdatedBy:   null.NewString("Weave - default"),
					Title:             "Verify your phone numbers",
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
					CompletedAt:       null.Time{},
					CompletedBy:       null.String{},
					VerifiedAt:        null.Time{},
					VerifiedBy:        null.String{},
					Content:           "In order to start using all our software features, we need to sync your patient data to Weave. Typically, we will do this by installing a sync application on your office server - so make sure you have access to that computer and the right login information. If your practice management system is cloud-based, we will instead walk you through how to generate the credentials we need. Click below to schedule a 30-60 minute phone call with a Weave technician.",
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
					CompletedAt:       null.Time{},
					CompletedBy:       null.String{},
					VerifiedAt:        null.Time{},
					VerifiedBy:        null.String{},
					Content:           "In order to make the most of features like Team Chat and two-way text messaging, we recommend installing Weave on most workstations throughout your office. Click below to find helpful instructions to download Weave throughout your office.",
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
					CompletedAt:       null.Time{},
					CompletedBy:       null.String{},
					VerifiedAt:        null.Time{},
					VerifiedBy:        null.String{},
					Content:           "In order to get the most out of the Weave software, we recommend tuning in to our training videos online. Our helpful training staff will walk you through setting up and using two-way texting, automated appointment reminders, and all our other features. Click below to watch our online webinars",
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
					CompletedAt:       null.Time{},
					CompletedBy:       null.String{},
					VerifiedAt:        null.Time{},
					VerifiedBy:        null.String{},
					Content:           "In order to see your schedule, field calls, and chat with your team all on the go, we recommend installing the Weave app on your mobile device. Click below to find helpful instructions to download Weave on your phone.",
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
					CompletedAt:       null.Time{},
					CompletedBy:       null.String{},
					VerifiedAt:        null.Time{},
					VerifiedBy:        null.String{},
					Content:           `In order to ensure the highest quality phone service on Weave's internet-based phones, we need to verify your network compatibility. Click below to schedule a 15 minute phone call with a Weave technician. <a href="http://www.weavehelp.com/network-specs">Learn More</a>`,
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
					CompletedAt:       null.Time{},
					CompletedBy:       null.String{},
					VerifiedAt:        null.Time{},
					VerifiedBy:        null.String{},
					Content:           `You're making great progress! Click below to schedule a 60 minute phone call with a Weave technician who will guide you through the process of connecting the new phones. Typically, we will aim to help you plug the new phones in side-by-side with your old phones. After finishing, you will continue to use your old phones while we work with your current phone service provider for a few days to officially move your phone numbers to our system. <a href="http://www.weavehelp.com/phone-install">Learn more.</a>`,
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
					CompletedAt:       null.Time{},
					CompletedBy:       null.String{},
					VerifiedAt:        null.Time{},
					VerifiedBy:        null.String{},
					Content:           `In order to get the most out of the Weave phones, we recommend tuning in to our training videos online. Our helpful training staff will walk you through setting up voicemail messages, placing callers on hold, transferring calls, and all our other phone features. Click below to watch our online webinars.`,
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
					CompletedAt:       null.Time{},
					CompletedBy:       null.String{},
					VerifiedAt:        null.Time{},
					VerifiedBy:        null.String{},
					Content:           `In order to have the best experience with your new phones, we recommend customizing the system to suit your needs. Your onboarding agent will help you change which phones ring when a call comes in, adjust the names and extensions on each phone, and set up any advanced call routing or auto attendant you may need. Click below to schedule a call.`,
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

					ButtonContent:     null.NewString("Schedule Call"),
					ButtonExternalURL: null.NewString("schedule porting link"),
					CompletedAt:       null.Time{},
					CompletedBy:       null.String{},
					VerifiedAt:        null.Time{},
					VerifiedBy:        null.String{},
					Content:           `One of your first steps is to work with your onboarding agent to verify your office phone numbers. Your onboarding agent will then handle some initial processing with your current phone service provider to ensure that when your office is ready to start using the new phones, we are ready to activate them for you. Click below to schedule a call with your onboarding agent.`,
					DisplayOrder:      4,
					Status:            0,
					StatusUpdatedAt:   time.Now(),
					StatusUpdatedBy:   null.NewString("Weave - default"),
					Title:             "Verify your phone numbers",
					Explanation:       null.String{},

					CreatedAt: time.Now(),
					UpdatedAt: time.Now(),
				},
			},
			wantErr: false,
		},
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
			if !compareTaskInstances(got, tt.want) {
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

// musUUID is a helper function that simply returns the value without checking the errors. This is usefule when using uuid.Parse directly in a struct definition.
func mustUUID(u uuid.UUID, err error) uuid.UUID {
	return u
}

func compareTaskInstances(a, b []app.TaskInstance) bool {
	if len(a) != len(b) {
		return false
	}

	sort.Slice(a, func(i, j int) bool { return a[i].DisplayOrder < a[j].DisplayOrder })
	sort.Slice(b, func(i, j int) bool { return b[i].DisplayOrder < b[j].DisplayOrder })

	for i := range a {
		if !compareTaskInstance(a[i], b[i]) {
			return false
		}
	}

	return true
}

func compareTaskInstance(a, b app.TaskInstance) bool {
	return a.LocationID == b.LocationID &&
		a.CategoryID == b.CategoryID &&
		a.TaskID == b.TaskID &&
		a.ButtonContent == b.ButtonContent &&
		a.ButtonExternalURL == b.ButtonExternalURL &&
		a.CompletedBy == b.CompletedBy &&
		a.VerifiedBy == b.VerifiedBy &&
		a.Content == b.Content &&
		a.DisplayOrder == b.DisplayOrder &&
		a.StatusUpdatedBy == b.StatusUpdatedBy &&
		a.Title == b.Title &&
		a.Explanation == b.Explanation
}
