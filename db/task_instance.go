package db

import (
	"context"
	"database/sql"
	"time"

	"weavelab.xyz/go-utilities/null"
	"weavelab.xyz/protorepo/dist/go/enums/insysenums"
	"weavelab.xyz/wlib/uuid"
	"weavelab.xyz/wlib/werror"
)

type OnboardingTaskInstance struct {
	ID         uuid.UUID
	LocationID uuid.UUID
	CategoryID uuid.UUID
	TaskID     uuid.UUID

	ButtonContent   null.String
	CompletedAt     null.Time
	CompletedBy     null.String
	VerifiedAt      null.Time
	VerifiedBy      null.String
	Content         string
	DisplayOrder    int
	Status          int
	StatusUpdatedAt time.Time
	StatusUpdatedBy string
	Title           string

	CreatedAt time.Time
	UpdatedAt time.Time
}

func TaskInstances(ctx context.Context, locationID uuid.UUID) ([]OnboardingTaskInstance, error) {
	var tasks []OnboardingTaskInstance

	query := `SELECT
		id, location_id, onboarding_category_id, onboarding_task_id, completed_at, completed_by, verified_at, verified_by, button_content, content, display_order, status, status_updated_at, status_updated_by, title, created_at, updated_at
	FROM insys_onboarding.onboarding_task_instances
	WHERE location_id = $1`

	rows, err := Conn.QueryContext(ctx, query, locationID.String())
	if err != nil {
		return tasks, werror.Wrap(err, "error exectuting onboarding tasks query")
	}
	defer rows.Close()

	for rows.Next() {
		var taskInstanceID, locationID, categoryID, taskID string
		var task OnboardingTaskInstance

		err := rows.Scan(
			&taskInstanceID,
			&locationID,
			&categoryID,
			&taskID,
			&task.CompletedAt,
			&task.CompletedBy,
			&task.VerifiedAt,
			&task.VerifiedBy,
			&task.ButtonContent,
			&task.Content,
			&task.DisplayOrder,
			&task.Status,
			&task.StatusUpdatedAt,
			&task.StatusUpdatedBy,
			&task.Title,
			&task.CreatedAt,
			&task.UpdatedAt,
		)
		if err != nil {
			return nil, werror.Wrap(err, "failed to scan onboarding tasks")
		}

		taskInstanceUUID, err := uuid.Parse(taskInstanceID)
		if err != nil {
			return nil, werror.Wrap(err, "failed to parse task instance id into uuid")
		}
		task.ID = taskInstanceUUID

		locationUUID, err := uuid.Parse(locationID)
		if err != nil {
			return nil, werror.Wrap(err, "failed to parse location id into uuid")
		}
		task.LocationID = locationUUID

		categoryUUID, err := uuid.Parse(categoryID)
		if err != nil {
			return nil, werror.Wrap(err, "failed to parse category id into uuid")
		}
		task.CategoryID = categoryUUID

		taskUUID, err := uuid.Parse(taskID)
		if err != nil {
			return nil, werror.Wrap(err, "failed to parse task id into uuid")
		}
		task.TaskID = taskUUID

		tasks = append(tasks, task)
	}

	return tasks, nil
}

const (
	defaultUpdateQuery   = `UPDATE insys_onboarding.onboarding_task_instances SET status=$1, status_updated_at=$2, status_updated_by=$3 WHERE id = $4 RETURNING id, location_id, onboarding_category_id, onboarding_task_id, completed_at, completed_by, verified_at, verified_by, button_content, content, display_order, status, status_updated_at, status_updated_by, title, created_at, updated_at`
	completedUpdateQuery = `UPDATE insys_onboarding.onboarding_task_instances SET status=$1, status_updated_at=$2, status_updated_by=$3, completed_at=$2, completed_by=$3 WHERE id = $4 RETURNING id, location_id, onboarding_category_id, onboarding_task_id, completed_at, completed_by, verified_at, verified_by, button_content, content, display_order, status, status_updated_at, status_updated_by, title, created_at, updated_at`
	verifiedUpdateQuery  = `UPDATE insys_onboarding.onboarding_task_instances SET status=$1, status_updated_at=$2, status_updated_by=$3, verified_at=$2, verified_by=$3 WHERE id = $4 RETURNING id, location_id, onboarding_category_id, onboarding_task_id, completed_at, completed_by, verified_at, verified_by, button_content, content, display_order, status, status_updated_at, status_updated_by, title, created_at, updated_at`
)

func UpdateTaskInstance(ctx context.Context, id uuid.UUID, status insysenums.OnboardingTaskStatus, statusUpdatedBy string) (OnboardingTaskInstance, error) {
	var taskInstanceID, locationID, categoryID, taskID string

	var task OnboardingTaskInstance
	var row *sql.Row

	switch status {
	case 0: // waiting on customer
		row = Conn.QueryRowContext(ctx, defaultUpdateQuery, status, time.Now(), statusUpdatedBy, id.String())
	case 1: // waiting on weave
		row = Conn.QueryRowContext(ctx, defaultUpdateQuery, status, time.Now(), statusUpdatedBy, id.String())
	case 2: // completed
		row = Conn.QueryRowContext(ctx, completedUpdateQuery, status, time.Now().UTC(), statusUpdatedBy, id.String())
	case 3: // verified
		row = Conn.QueryRowContext(ctx, verifiedUpdateQuery, status, time.Now().UTC(), statusUpdatedBy, id.String())
	default:
		return task, werror.New("Cannot update task instance. Not a valid status. Valid paremeters are 0, 1, 2, or 3")
	}

	err := row.Scan(
		&taskInstanceID,
		&locationID,
		&categoryID,
		&taskID,
		&task.CompletedAt,
		&task.CompletedBy,
		&task.VerifiedAt,
		&task.VerifiedBy,
		&task.ButtonContent,
		&task.Content,
		&task.DisplayOrder,
		&task.Status,
		&task.StatusUpdatedAt,
		&task.StatusUpdatedBy,
		&task.Title,
		&task.CreatedAt,
		&task.UpdatedAt,
	)
	if err != nil {
		return task, werror.Wrap(err, "error updating task")
	}

	taskInstanceUUID, err := uuid.Parse(taskInstanceID)
	if err != nil {
		return OnboardingTaskInstance{}, werror.Wrap(err, "failed to parse task instance id into uuid")
	}
	task.ID = taskInstanceUUID

	locationUUID, err := uuid.Parse(locationID)
	if err != nil {
		return OnboardingTaskInstance{}, werror.Wrap(err, "failed to parse location id into uuid")
	}
	task.LocationID = locationUUID

	categoryUUID, err := uuid.Parse(categoryID)
	if err != nil {
		return OnboardingTaskInstance{}, werror.Wrap(err, "failed to parse category id into uuid")
	}
	task.CategoryID = categoryUUID

	taskUUID, err := uuid.Parse(taskID)
	if err != nil {
		return task, werror.Wrap(err, "failed to parse task id into uuid")
	}
	task.TaskID = taskUUID

	return task, nil
}

func CreateTaskInstancesFromTasks(ctx context.Context, locationID uuid.UUID) ([]OnboardingTaskInstance, error) {
	var taskInstances []OnboardingTaskInstance

	query := `INSERT INTO insys_onboarding.onboarding_task_instances
		(id, location_id, title, content, button_content, display_order, status, status_updated_at, status_updated_by, created_at, updated_at, onboarding_category_id, onboarding_task_id )
		SELECT overlay(overlay(md5(random()::text || ':' || clock_timestamp()::text) placing '4' from 13) placing '8' from 17)::uuid, $1, title, content, button_content, display_order, 0, $2, 'Weave - default', $2, $2, onboarding_category_id, id FROM insys_onboarding.onboarding_tasks
		RETURNING id, location_id, onboarding_category_id, onboarding_task_id, completed_at, completed_by, verified_at, verified_by, button_content, content, display_order, status, status_updated_at, status_updated_by, title, created_at, updated_at;`

	rows, err := Conn.QueryContext(ctx, query, locationID.String(), time.Now().UTC())
	if err != nil {
		return nil, werror.Wrap(err, "error exectuting create task instances from tasks query")
	}
	defer rows.Close()

	for rows.Next() {
		var taskInstanceID, locationID, categoryID, taskID string
		var taskInstance OnboardingTaskInstance

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
