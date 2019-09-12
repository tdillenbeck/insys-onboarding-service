-- +goose Up
-- +goose StatementBegin
ALTER TABLE insys_onboarding.chili_piper_schedule_events ADD canceled_at timestamp with time zone;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE insys_onboarding.chili_piper_schedule_events DROP COLUMN canceled_at;
-- +goose StatementEnd
