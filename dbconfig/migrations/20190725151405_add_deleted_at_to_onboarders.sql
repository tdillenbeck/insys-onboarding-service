-- +goose Up
-- +goose StatementBegin
ALTER TABLE insys_onboarding.onboarders ADD deleted_at timestamp with time zone;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE insys_onboarding.onboarders DROP COLUMN deleted_at;
-- +goose StatementEnd
