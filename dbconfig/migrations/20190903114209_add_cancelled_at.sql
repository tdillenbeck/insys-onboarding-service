-- +goose Up
-- +goose StatementBegin
ALTER TABLE insys_onboarding.chili_piper_schedule_events ADD cancelled_at timestamp with time zone;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE insys_onboarding.chili_piper_schedule_events DROP COLUMN cancelled_at;
-- +goose StatementEnd
