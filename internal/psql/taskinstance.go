package psql

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"weavelab.xyz/insys-onboarding-service/internal/app"

	"weavelab.xyz/monorail/shared/protorepo/dist/go/enums/insysenums"
	"weavelab.xyz/monorail/shared/wlib/uuid"
	"weavelab.xyz/monorail/shared/wlib/werror"
	"weavelab.xyz/monorail/shared/wlib/wsql"
)

type TaskInstanceService struct {
	DB *wsql.PG
}

func (t *TaskInstanceService) ByLocationID(ctx context.Context, locationID uuid.UUID) ([]app.TaskInstance, error) {
	var taskInstances []app.TaskInstance

	query := `
SELECT
	id, location_id, onboarding_category_id, onboarding_task_id, completed_at, completed_by, verified_at, verified_by, button_content, button_external_url, button_internal_url, content, display_order, status, status_updated_at, status_updated_by, title, explanation, created_at, updated_at
FROM insys_onboarding.onboarding_task_instances
WHERE location_id = $1
`

	rows, err := t.DB.QueryContext(ctx, query, locationID.String())
	if err != nil {
		return nil, werror.Wrap(err, "error exectuting onboarding tasks query")
	}
	defer rows.Close()

	for rows.Next() {
		var taskInstance app.TaskInstance

		err := rows.Scan(
			&taskInstance.ID,
			&taskInstance.LocationID,
			&taskInstance.CategoryID,
			&taskInstance.TaskID,
			&taskInstance.CompletedAt,
			&taskInstance.CompletedBy,
			&taskInstance.VerifiedAt,
			&taskInstance.VerifiedBy,
			&taskInstance.ButtonContent,
			&taskInstance.ButtonExternalURL,
			&taskInstance.ButtonInternalURL,
			&taskInstance.Content,
			&taskInstance.DisplayOrder,
			&taskInstance.Status,
			&taskInstance.StatusUpdatedAt,
			&taskInstance.StatusUpdatedBy,
			&taskInstance.Title,
			&taskInstance.Explanation,
			&taskInstance.CreatedAt,
			&taskInstance.UpdatedAt,
		)
		if err != nil {
			return nil, werror.Wrap(err, "failed to scan onboarding taskInstances")
		}

		taskInstances = append(taskInstances, taskInstance)
	}

	return taskInstances, nil
}

// CreateFromTasks will copy the template data from the tasks database table into the task_instances database table along with populating the location id.
func (t *TaskInstanceService) CreateFromTasks(ctx context.Context, locationID uuid.UUID) ([]app.TaskInstance, error) {
	var taskInstances []app.TaskInstance

	/*
		Copy information from the tasks table into the tasks instance table while inserting the location ID. Also use the location ID to find the assigned onboarder
		and use their personal scheduling links that they have setup as the button_external_url for specific tasks. Some of the tasks require a custom URL ,
		and some use a default URL (see the mapping in case statement).

		NOTE: "overlay(overlay(md5(random()::text || ':' || clock_timestamp()::text) placing '4' from 13) placing '8' from 17)::uuid" is generating a v4 uuid for each new task instance
	*/
	query := `
INSERT INTO insys_onboarding.onboarding_task_instances
	(id, location_id, title, content, button_content, button_external_url, button_internal_url, display_order, status, status_updated_at, status_updated_by, created_at, updated_at, onboarding_category_id, onboarding_task_id)
	SELECT overlay(overlay(md5(random()::text || ':' || clock_timestamp()::text) placing '4' from 13) placing '8' from 17)::uuid,
		$1, -- location id
		title,
		content,
		button_content,
		CASE id -- map task to specific onboarder url to use for scheduling for button_external_url field
			WHEN '2d2df285-9211-48fc-a057-74f7dee2d9a4' THEN COALESCE(NULLIF((SELECT schedule_customization_link FROM insys_onboarding.onboarders AS a INNER JOIN insys_onboarding.onboarders_location AS b ON a.id = b.onboarder_id WHERE b.location_id=$1), ''), button_external_url)
			WHEN '9aec502b-f8b8-4f10-9748-1fe4050eacde' THEN COALESCE(NULLIF((SELECT schedule_porting_link FROM insys_onboarding.onboarders AS a INNER JOIN insys_onboarding.onboarders_location AS b ON a.id = b.onboarder_id WHERE b.location_id=$1), ''), button_external_url)
			WHEN '7b15e061-8002-4edc-9bf4-f38c6eec6364' THEN COALESCE(NULLIF((SELECT schedule_network_link FROM insys_onboarding.onboarders AS a INNER JOIN insys_onboarding.onboarders_location AS b ON a.id = b.onboarder_id WHERE b.location_id=$1), ''), button_external_url)
			WHEN '16a6dc91-ec6b-4b09-b591-a5b0dfa92932' THEN COALESCE(NULLIF((SELECT schedule_software_install_link FROM insys_onboarding.onboarders AS a INNER JOIN insys_onboarding.onboarders_location AS b ON a.id = b.onboarder_id WHERE b.location_id=$1), ''), button_external_url)
			WHEN 'fd4f656c-c9f1-47b8-96ad-3080b999a843' THEN COALESCE(NULLIF((SELECT schedule_phone_install_link FROM insys_onboarding.onboarders AS a INNER JOIN insys_onboarding.onboarders_location AS b ON a.id = b.onboarder_id WHERE b.location_id=$1), ''), button_external_url)
			WHEN 'c20b65d8-e281-4e62-98f0-4aebf83e0bee' THEN COALESCE(NULLIF((SELECT schedule_software_training_link FROM insys_onboarding.onboarders AS a INNER JOIN insys_onboarding.onboarders_location AS b ON a.id = b.onboarder_id WHERE b.location_id=$1), ''), button_external_url)
			WHEN '47743fae-c775-45d5-8a51-dc7e3371dfa4' THEN COALESCE(NULLIF((SELECT schedule_phone_training_link FROM insys_onboarding.onboarders AS a INNER JOIN insys_onboarding.onboarders_location AS b ON a.id = b.onboarder_id WHERE b.location_id=$1), ''), button_external_url)
			ELSE button_external_url
		END,
		button_internal_url,
		display_order,
		0, -- default status
		now(),
		'Weave - default', -- default status_updated_by
		now(),
		now(),
		onboarding_category_id,
		id
	FROM insys_onboarding.onboarding_tasks
RETURNING id, location_id, onboarding_category_id, onboarding_task_id, completed_at, completed_by, verified_at, verified_by, button_content, button_external_url, button_internal_url, content, display_order, status, status_updated_at, status_updated_by, title, explanation, created_at, updated_at;
`
	rows, err := t.DB.QueryContext(ctx, query, locationID.String())
	if err != nil {
		return nil, werror.Wrap(err, "error exectuting create task instances from tasks query")
	}
	defer rows.Close()

	for rows.Next() {
		var taskInstance app.TaskInstance

		err := rows.Scan(
			&taskInstance.ID,
			&taskInstance.LocationID,
			&taskInstance.CategoryID,
			&taskInstance.TaskID,
			&taskInstance.CompletedAt,
			&taskInstance.CompletedBy,
			&taskInstance.VerifiedAt,
			&taskInstance.VerifiedBy,
			&taskInstance.ButtonContent,
			&taskInstance.ButtonExternalURL,
			&taskInstance.ButtonInternalURL,
			&taskInstance.Content,
			&taskInstance.DisplayOrder,
			&taskInstance.Status,
			&taskInstance.StatusUpdatedAt,
			&taskInstance.StatusUpdatedBy,
			&taskInstance.Title,
			&taskInstance.Explanation,
			&taskInstance.CreatedAt,
			&taskInstance.UpdatedAt,
		)
		if err != nil {
			return nil, werror.Wrap(err, "failed to scan onboarding taskInstances")
		}

		taskInstances = append(taskInstances, taskInstance)
	}

	return taskInstances, nil
}

// SyncTaskInstanceLinksFromOnboarderLinks will go to the assigned onbaorder in the onboarderslocation table and lookup the onboarder link
// for the list of task id's.
func (t *TaskInstanceService) SyncTaskInstanceLinksFromOnboarderLinks(ctx context.Context, locationID uuid.UUID) error {
	onboarderLinkToTaskID := []struct {
		onboarderLinkColumnName string
		taskID                  string
	}{
		{
			"schedule_customization_link",
			"2d2df285-9211-48fc-a057-74f7dee2d9a4",
		},
		{
			"schedule_porting_link",
			"9aec502b-f8b8-4f10-9748-1fe4050eacde",
		},
		{
			"schedule_network_link",
			"7b15e061-8002-4edc-9bf4-f38c6eec6364",
		},
		{
			"schedule_software_install_link",
			"16a6dc91-ec6b-4b09-b591-a5b0dfa92932",
		},
		{
			"schedule_phone_install_link",
			"fd4f656c-c9f1-47b8-96ad-3080b999a843",
		},
		{
			"schedule_software_training_link",
			"c20b65d8-e281-4e62-98f0-4aebf83e0bee",
		},
		{
			"schedule_phone_training_link",
			"47743fae-c775-45d5-8a51-dc7e3371dfa4",
		},
	}

	query := `
UPDATE insys_onboarding.onboarding_task_instances
SET button_external_url = (
	SELECT
			COALESCE(NULLIF((SELECT %s FROM insys_onboarding.onboarders AS a INNER JOIN insys_onboarding.onboarders_location AS b ON a.id = b.onboarder_id WHERE b.location_id=$1), ''), button_external_url)
	FROM insys_onboarding.onboarding_tasks
	WHERE id=$2
)
WHERE location_id=$1 AND onboarding_task_id=$2;
	`

	for _, mapping := range onboarderLinkToTaskID {
		// cannot use the column name as a bindvar in the postgres query. inject the name before executing the query.
		formattedQuery := fmt.Sprintf(query, mapping.onboarderLinkColumnName)
		_, err := t.DB.ExecContext(ctx, formattedQuery, locationID, mapping.taskID)
		if err != nil {
			return werror.Wrap(err, "could not sync task instance links for onboarder").Add("locationID", locationID).Add("taskID", mapping.taskID)
		}
	}

	return nil
}

const (
	defaultUpdateQuery = `
UPDATE insys_onboarding.onboarding_task_instances
SET status=$1, status_updated_at=$2, status_updated_by=$3
WHERE id=$4
RETURNING id, location_id, onboarding_category_id, onboarding_task_id, completed_at, completed_by, verified_at, verified_by, button_content, button_external_url, button_internal_url, content, display_order, status, status_updated_at, status_updated_by, title, explanation, created_at, updated_at
`
	completedUpdateQuery = `
UPDATE insys_onboarding.onboarding_task_instances
SET status=$1, status_updated_at=$2, status_updated_by=$3, completed_at=$2, completed_by=$3
WHERE id = $4
RETURNING id, location_id, onboarding_category_id, onboarding_task_id, completed_at, completed_by, verified_at, verified_by, button_content, button_external_url, button_internal_url, content, display_order, status, status_updated_at, status_updated_by, title, explanation, created_at, updated_at
`
	verifiedUpdateQuery = `
UPDATE insys_onboarding.onboarding_task_instances
SET status=$1, status_updated_at=$2, status_updated_by=$3, verified_at=$2, verified_by=$3
WHERE id = $4
RETURNING id, location_id, onboarding_category_id, onboarding_task_id, completed_at, completed_by, verified_at, verified_by, button_content, button_external_url, button_internal_url, content, display_order, status, status_updated_at, status_updated_by, title, explanation, created_at, updated_at
`
)

// Set the status and status_updated_by fields for the task_instance record
func (t *TaskInstanceService) Update(ctx context.Context, id uuid.UUID, status insysenums.OnboardingTaskStatus, statusUpdatedBy string) (*app.TaskInstance, error) {
	var taskInstance app.TaskInstance
	var row *sql.Row

	switch status {
	case 0: // waiting on customer
		row = t.DB.QueryRowContext(ctx, defaultUpdateQuery, status, time.Now(), statusUpdatedBy, id.String())
	case 1: // waiting on weave
		row = t.DB.QueryRowContext(ctx, defaultUpdateQuery, status, time.Now(), statusUpdatedBy, id.String())
	case 2: // completed
		row = t.DB.QueryRowContext(ctx, completedUpdateQuery, status, time.Now().UTC(), statusUpdatedBy, id.String())
	case 3: // verified
		row = t.DB.QueryRowContext(ctx, verifiedUpdateQuery, status, time.Now().UTC(), statusUpdatedBy, id.String())
	default:
		return nil, werror.New("Cannot update task instance. Not a valid status. Valid paremeters are 0, 1, 2, or 3")
	}

	err := row.Scan(
		&taskInstance.ID,
		&taskInstance.LocationID,
		&taskInstance.CategoryID,
		&taskInstance.TaskID,
		&taskInstance.CompletedAt,
		&taskInstance.CompletedBy,
		&taskInstance.VerifiedAt,
		&taskInstance.VerifiedBy,
		&taskInstance.ButtonContent,
		&taskInstance.ButtonExternalURL,
		&taskInstance.ButtonInternalURL,
		&taskInstance.Content,
		&taskInstance.DisplayOrder,
		&taskInstance.Status,
		&taskInstance.StatusUpdatedAt,
		&taskInstance.StatusUpdatedBy,
		&taskInstance.Title,
		&taskInstance.Explanation,
		&taskInstance.CreatedAt,
		&taskInstance.UpdatedAt,
	)
	if err != nil {
		return nil, werror.Wrap(err, "error updating taskInstance")
	}

	return &taskInstance, nil
}

func (t *TaskInstanceService) UpdateExplanation(ctx context.Context, id uuid.UUID, explanation string) (*app.TaskInstance, error) {
	var taskInstance app.TaskInstance

	query := `
UPDATE insys_onboarding.onboarding_task_instances
SET explanation=$1, updated_at=$2
WHERE id=$3
RETURNING id, location_id, onboarding_category_id, onboarding_task_id, completed_at, completed_by, verified_at, verified_by, button_content, button_external_url, button_internal_url, content, display_order, status, status_updated_at, status_updated_by, title, explanation, created_at, updated_at
`
	row := t.DB.QueryRowContext(ctx, query, explanation, time.Now(), id.String())
	err := row.Scan(
		&taskInstance.ID,
		&taskInstance.LocationID,
		&taskInstance.CategoryID,
		&taskInstance.TaskID,
		&taskInstance.CompletedAt,
		&taskInstance.CompletedBy,
		&taskInstance.VerifiedAt,
		&taskInstance.VerifiedBy,
		&taskInstance.ButtonContent,
		&taskInstance.ButtonExternalURL,
		&taskInstance.ButtonInternalURL,
		&taskInstance.Content,
		&taskInstance.DisplayOrder,
		&taskInstance.Status,
		&taskInstance.StatusUpdatedAt,
		&taskInstance.StatusUpdatedBy,
		&taskInstance.Title,
		&taskInstance.Explanation,
		&taskInstance.CreatedAt,
		&taskInstance.UpdatedAt,
	)
	if err != nil {
		return nil, werror.Wrap(err, "error updating taskInstance")
	}

	return &taskInstance, nil
}
