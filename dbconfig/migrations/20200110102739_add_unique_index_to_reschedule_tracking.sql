-- +goose NO TRANSACTION

-- +goose Up
CREATE UNIQUE INDEX CONCURRENTLY index_reschedule_event_tracking_on_location_id_and_event_type ON insys_onboarding.reschedule_tracking USING btree (location_id, event_type);

-- +goose Down
DROP INDEX CONCURRENTLY IF EXISTS  index_reschedule_event_tracking_on_location_id_and_event_type;