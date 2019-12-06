-- +goose Up
-- SQL in this section is executed when the migration is applied.
CREATE TABLE insys_onboarding.handoff_snapshots (
  id uuid NOT NULL PRIMARY KEY,
  onboarders_location_id uuid NOT NULL UNIQUE,
  csat_recipient_user_id uuid,
  csat_sent_at timestamp with time zone,

  created_at timestamp with time zone NOT NULL DEFAULT now(),
  updated_at timestamp with time zone NOT NULL DEFAULT now()
);

ALTER TABLE ONLY insys_onboarding.handoff_snapshots
    ADD CONSTRAINT handoff_snapshots_onboarders_location_fkey FOREIGN KEY (onboarders_location_id) REFERENCES insys_onboarding.onboarders_location(id);

-- +goose Down
-- SQL in this section is executed when the migration is rolled back.
DROP TABLE insys_onboarding.handoff_snapshots;