-- +goose Up
-- SQL in this section is executed when the migration is applied.
CREATE TABLE insys_onboarding.handoff_snapshots (
  id uuid NOT NULL PRIMARY KEY,
  onboarders_location_id uuid NOT NULL UNIQUE,

  billing_notes text,
  csat_recipient_user_email text,
  csat_sent_at timestamp with time zone,
  customization_setup text,
  customizations boolean,
  disclaimer_type_sent text,
  fax_port_submitted text,
  handed_off_at timestamp with time zone,
  network_decision text,
  notes text,
  point_of_contact_email text,
  reason_for_purchase text,
  router_make_and_model text,
  router_type text,

  created_at timestamp with time zone DEFAULT now() NOT NULL,
  updated_at timestamp with time zone DEFAULT now() NOT NULL
);

ALTER TABLE ONLY insys_onboarding.handoff_snapshots
  ADD CONSTRAINT handoff_snapshots_onboarders_location_fkey FOREIGN KEY (onboarders_location_id) REFERENCES insys_onboarding.onboarders_location(id);

-- +goose Down
-- SQL in this section is executed when the migration is rolled back.
DROP TABLE insys_onboarding.handoff_snapshots;
