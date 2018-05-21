-- +goose Up
ALTER TABLE insys_onboarding.onboarding_tasks ADD button_content text;
ALTER TABLE insys_onboarding.onboarding_task_instances ADD button_content text;

-- +goose Down
ALTER TABLE insys_onboarding.onboarding_tasks DROP COLUMN button_content;
ALTER TABLE insys_onboarding.onboarding_task_instances DROP COLUMN button_content;
