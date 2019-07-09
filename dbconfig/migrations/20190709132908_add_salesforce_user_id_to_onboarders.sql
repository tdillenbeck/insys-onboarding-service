-- +goose Up
-- +goose StatementBegin
ALTER TABLE insys_onboarding.onboarders ADD salesforce_user_id text;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE insys_onboarding.onboarders DROP COLUMN salesforce_user_id;
-- +goose StatementEnd
