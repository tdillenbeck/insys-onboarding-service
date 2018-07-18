-- +goose Up
-- SQL in this section is executed when the migration is applied.
ALTER TABLE insys_onboarding.onboarding_task_instances ADD explanation text;

-- +goose Down
-- SQL in this section is executed when the migration is rolled back.
ALTER TABLE insys_onboarding.onboarding_task_instances DROP COLUMN explanation;
