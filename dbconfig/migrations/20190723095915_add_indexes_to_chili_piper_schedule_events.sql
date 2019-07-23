-- +goose NO TRANSACTION

-- +goose Up
ALTER TABLE insys_onboarding.chili_piper_schedule_events ALTER COLUMN event_id SET NOT NULL;
CREATE UNIQUE INDEX CONCURRENTLY index_chili_piper_schedule_events_on_event_id ON insys_onboarding.chili_piper_schedule_events USING btree (event_id);
CREATE INDEX CONCURRENTLY index_chili_piper_schedule_events_on_location_id ON insys_onboarding.chili_piper_schedule_events USING btree (location_id);

-- +goose Down
DROP INDEX CONCURRENTLY IF EXISTS index_chili_piper_schedule_events_on_location_id;
DROP INDEX CONCURRENTLY IF EXISTS index_chili_piper_schedule_events_on_event_id;
