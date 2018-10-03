-- +goose Up
-- SQL in this section is executed when the migration is applied.
ALTER TABLE insys_onboarding.onboarding_tasks ADD button_internal_url text;
ALTER TABLE insys_onboarding.onboarding_task_instances ADD button_internal_url text;

-- +goose Down
-- SQL in this section is executed when the migration is rolled back.
ALTER TABLE insys_onboarding.onboarding_tasks DROP COLUMN button_internal_url;
ALTER TABLE insys_onboarding.onboarding_task_instances DROP COLUMN button_internal_url;
