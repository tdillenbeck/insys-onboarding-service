-- +goose Up
-- SQL in this section is executed when the migration is applied.
ALTER TABLE insys_onboarding.onboarders_location ADD region text;
ALTER TABLE insys_onboarding.onboarders_location ADD salesforce_opportunity_id text;


-- +goose Down
-- SQL in this section is executed when the migration is rolled back.
ALTER TABLE insys_onboarding.onboarders_location DROP COLUMN region;
ALTER TABLE insys_onboarding.onboarders_location DROP COLUMN salesforce_opportunity_id;