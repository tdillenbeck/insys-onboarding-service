-- +goose NO TRANSACTION

-- +goose Up
CREATE INDEX CONCURRENTLY index_onboarding_task_instances_on_location_id ON insys_onboarding.onboarding_task_instances USING btree (location_id);

-- +goose Down
DROP INDEX CONCURRENTLY IF EXISTS index_onboarding_task_instances_on_location_id;
