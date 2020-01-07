-- +goose Up
-- SQL in this section is executed when the migration is applied.
ALTER TABLE insys_onboarding.handoff_snapshots
    rename column point_of_contact to point_of_contact_email;

ALTER TABLE insys_onboarding.handoff_snapshots
    rename column csat_recipient_user_id to csat_recipient_user_email;

ALTER TABLE insys_onboarding.handoff_snapshots
    ADD disclaimer_type_sent text;

-- +goose Down
-- SQL in this section is executed when the migration is rolled back.
ALTER TABLE insys_onboarding.handoff_snapshots
    rename column point_of_contact_email to point_of_contact;

ALTER TABLE insys_onboarding.handoff_snapshots
    rename column csat_recipient_user_email to csat_recipient_user_id;


ALTER TABLE insys_onboarding.handoff_snapshots
    DROP COLUMN disclaimer_type_sent
;