-- +goose Up
-- SQL in this section is executed when the migration is applied.
ALTER TABLE insys_onboarding.handoff_snapshots
    ALTER COLUMN point_of_contact TYPE text,
    ALTER COLUMN csat_recipient_user_id TYPE text,
    ADD disclaimer_type_sent text
;

ALTER TABLE insys_onboarding.handoff_snapshots
    rename column point_of_contact to point_of_contact_email;

ALTER TABLE insys_onboarding.handoff_snapshots
    rename column csat_recipient_user_id to csat_recipient_user_email;

-- +goose Down
-- SQL in this section is executed when the migration is rolled back.
ALTER TABLE insys_onboarding.handoff_snapshots
    drop column point_of_contact_email,
    drop column csat_recipient_user_email,
    ADD point_of_contact uuid,
    ADD csat_recipient_user_id uuid,
    DROP COLUMN disclaimer_type_sent
;