package psql

import (
	"context"
	"database/sql"
	"time"

	app "weavelab.xyz/insys-onboarding"

	"weavelab.xyz/protorepo/dist/go/enums/insysenums"
	"weavelab.xyz/wlib/uuid"
	"weavelab.xyz/wlib/werror"
	"weavelab.xyz/wlib/wsql"
)

type TaskInstanceService struct {
	DB *wsql.PG
}

func (t *TaskInstanceService) ByLocationID(ctx context.Context, locationID uuid.UUID) ([]app.TaskInstance, error) {
	var taskInstances []app.TaskInstance

	query := `SELECT
							id, location_id, onboarding_category_id, onboarding_task_id, completed_at, completed_by, verified_at, verified_by, button_content, button_external_url, content, display_order, status, status_updated_at, status_updated_by, title, created_at, updated_at
						FROM insys_onboarding.onboarding_task_instances
						WHERE location_id = $1`

	rows, err := t.DB.QueryContext(ctx, query, locationID.String())
	if err != nil {
		return nil, werror.Wrap(err, "error exectuting onboarding tasks query")
	}
	defer rows.Close()

	for rows.Next() {
		var taskInstanceID, locationID, categoryID, taskID string
		var taskInstance app.TaskInstance

		err := rows.Scan(
			&taskInstanceID,
			&locationID,
			&categoryID,
			&taskID,
			&taskInstance.CompletedAt,
			&taskInstance.CompletedBy,
			&taskInstance.VerifiedAt,
			&taskInstance.VerifiedBy,
			&taskInstance.ButtonContent,
			&taskInstance.ButtonExternalURL,
			&taskInstance.Content,
			&taskInstance.DisplayOrder,
			&taskInstance.Status,
			&taskInstance.StatusUpdatedAt,
			&taskInstance.StatusUpdatedBy,
			&taskInstance.Title,
			&taskInstance.CreatedAt,
			&taskInstance.UpdatedAt,
		)
		if err != nil {
			return nil, werror.Wrap(err, "failed to scan onboarding taskInstances")
		}

		taskInstanceUUID, err := uuid.Parse(taskInstanceID)
		if err != nil {
			return nil, werror.Wrap(err, "failed to parse task instance id into uuid")
		}
		taskInstance.ID = taskInstanceUUID

		locationUUID, err := uuid.Parse(locationID)
		if err != nil {
			return nil, werror.Wrap(err, "failed to parse location id into uuid")
		}
		taskInstance.LocationID = locationUUID

		categoryUUID, err := uuid.Parse(categoryID)
		if err != nil {
			return nil, werror.Wrap(err, "failed to parse category id into uuid")
		}
		taskInstance.CategoryID = categoryUUID

		taskUUID, err := uuid.Parse(taskID)
		if err != nil {
			return nil, werror.Wrap(err, "failed to parse task id into uuid")
		}
		taskInstance.TaskID = taskUUID

		taskInstances = append(taskInstances, taskInstance)
	}

	return taskInstances, nil
}

// CreateFromTasks will copy the template data from the tasks database table into the task_instances database table along with populating the location id.
func (t *TaskInstanceService) CreateFromTasks(ctx context.Context, locationID uuid.UUID) ([]app.TaskInstance, error) {
	var taskInstances []app.TaskInstance

	query := `INSERT INTO insys_onboarding.onboarding_task_instances
							(id, location_id, title, content, button_content, button_external_url, display_order, status, status_updated_at, status_updated_by, created_at, updated_at, onboarding_category_id, onboarding_task_id )
						SELECT overlay(overlay(md5(random()::text || ':' || clock_timestamp()::text) placing '4' from 13) placing '8' from 17)::uuid, $1, title, content, button_content, button_external_url, display_order, 0, $2, 'Weave - default', $2, $2, onboarding_category_id, id FROM insys_onboarding.onboarding_tasks
						RETURNING id, location_id, onboarding_category_id, onboarding_task_id, completed_at, completed_by, verified_at, verified_by, button_content, button_external_url, content, display_order, status, status_updated_at, status_updated_by, title, created_at, updated_at;`

	rows, err := t.DB.QueryContext(ctx, query, locationID.String(), time.Now().UTC())
	if err != nil {
		return nil, werror.Wrap(err, "error exectuting create task instances from tasks query")
	}
	defer rows.Close()

	for rows.Next() {
		var taskInstanceID, locationID, categoryID, taskID string
		var taskInstance app.TaskInstance

		err := rows.Scan(
			&taskInstanceID,
			&locationID,
			&categoryID,
			&taskID,
			&taskInstance.CompletedAt,
			&taskInstance.CompletedBy,
			&taskInstance.VerifiedAt,
			&taskInstance.VerifiedBy,
			&taskInstance.ButtonContent,
			&taskInstance.ButtonExternalURL,
			&taskInstance.Content,
			&taskInstance.DisplayOrder,
			&taskInstance.Status,
			&taskInstance.StatusUpdatedAt,
			&taskInstance.StatusUpdatedBy,
			&taskInstance.Title,
			&taskInstance.CreatedAt,
			&taskInstance.UpdatedAt,
		)
		if err != nil {
			return nil, werror.Wrap(err, "failed to scan onboarding taskInstances")
		}

		taskInstanceUUID, err := uuid.Parse(taskInstanceID)
		if err != nil {
			return nil, werror.Wrap(err, "failed to parse task instance id into uuid")
		}
		taskInstance.ID = taskInstanceUUID

		locationUUID, err := uuid.Parse(locationID)
		if err != nil {
			return nil, werror.Wrap(err, "failed to parse location id into uuid")
		}
		taskInstance.LocationID = locationUUID

		categoryUUID, err := uuid.Parse(categoryID)
		if err != nil {
			return nil, werror.Wrap(err, "failed to parse category id into uuid")
		}
		taskInstance.CategoryID = categoryUUID

		taskUUID, err := uuid.Parse(taskID)
		if err != nil {
			return nil, werror.Wrap(err, "failed to parse task id into uuid")
		}
		taskInstance.TaskID = taskUUID

		taskInstances = append(taskInstances, taskInstance)
	}

	return taskInstances, nil
}

const (
	defaultUpdateQuery   = `UPDATE insys_onboarding.onboarding_task_instances SET status=$1, status_updated_at=$2, status_updated_by=$3 WHERE id = $4 RETURNING id, location_id, onboarding_category_id, onboarding_task_id, completed_at, completed_by, verified_at, verified_by, button_content, button_external_url, content, display_order, status, status_updated_at, status_updated_by, title, created_at, updated_at`
	completedUpdateQuery = `UPDATE insys_onboarding.onboarding_task_instances SET status=$1, status_updated_at=$2, status_updated_by=$3, completed_at=$2, completed_by=$3 WHERE id = $4 RETURNING id, location_id, onboarding_category_id, onboarding_task_id, completed_at, completed_by, verified_at, verified_by, button_content, button_external_url, content, display_order, status, status_updated_at, status_updated_by, title, created_at, updated_at`
	verifiedUpdateQuery  = `UPDATE insys_onboarding.onboarding_task_instances SET status=$1, status_updated_at=$2, status_updated_by=$3, verified_at=$2, verified_by=$3 WHERE id = $4 RETURNING id, location_id, onboarding_category_id, onboarding_task_id, completed_at, completed_by, verified_at, verified_by, button_content, button_external_url, content, display_order, status, status_updated_at, status_updated_by, title, created_at, updated_at`
)

func (t *TaskInstanceService) Update(ctx context.Context, id uuid.UUID, status insysenums.OnboardingTaskStatus, statusUpdatedBy string) (*app.TaskInstance, error) {
	var taskInstanceID, locationID, categoryID, taskID string

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
		&taskInstanceID,
		&locationID,
		&categoryID,
		&taskID,
		&taskInstance.CompletedAt,
		&taskInstance.CompletedBy,
		&taskInstance.VerifiedAt,
		&taskInstance.VerifiedBy,
		&taskInstance.ButtonContent,
		&taskInstance.ButtonExternalURL,
		&taskInstance.Content,
		&taskInstance.DisplayOrder,
		&taskInstance.Status,
		&taskInstance.StatusUpdatedAt,
		&taskInstance.StatusUpdatedBy,
		&taskInstance.Title,
		&taskInstance.CreatedAt,
		&taskInstance.UpdatedAt,
	)
	if err != nil {
		return nil, werror.Wrap(err, "error updating taskInstance")
	}

	taskInstanceUUID, err := uuid.Parse(taskInstanceID)
	if err != nil {
		return nil, werror.Wrap(err, "failed to parse taskInstance id into uuid")
	}
	taskInstance.ID = taskInstanceUUID

	locationUUID, err := uuid.Parse(locationID)
	if err != nil {
		return nil, werror.Wrap(err, "failed to parse taskInstance location id into uuid")
	}
	taskInstance.LocationID = locationUUID

	categoryUUID, err := uuid.Parse(categoryID)
	if err != nil {
		return nil, werror.Wrap(err, "failed to parse taskInstance category id into uuid")
	}
	taskInstance.CategoryID = categoryUUID

	taskUUID, err := uuid.Parse(taskID)
	if err != nil {
		return nil, werror.Wrap(err, "failed to parse ltask id into uuid")
	}
	taskInstance.TaskID = taskUUID

	return &taskInstance, nil
}
