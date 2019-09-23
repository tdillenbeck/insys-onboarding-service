-- +goose Up
-- SQL in this section is executed when the migration is applied.
ALTER TABLE insys_onboarding.onboarders_location ADD user_first_logged_in_at timestamp with time zone;


-- +goose Down
-- SQL in this section is executed when the migration is rolled back.
ALTER TABLE insys_onboarding.onboarders_location DROP COLUMN user_first_logged_in_at;
