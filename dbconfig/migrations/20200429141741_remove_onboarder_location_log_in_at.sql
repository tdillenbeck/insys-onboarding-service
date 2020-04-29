-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
ALTER TABLE insys_onboarding.onboarders_location DROP COLUMN user_first_logged_in_at;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
ALTER TABLE insys_onboarding.onboarders_location ADD user_first_logged_in_at timestamp with time zone;
-- +goose StatementEnd
