-- +goose Up
-- SQL in this section is executed when the migration is applied.
ALTER TABLE insys_onboarding.handoff_snapshots
  ADD handed_off_at timestamp with time zone,

  ADD point_of_contact uuid,
  ADD reason_for_purchase text,
  ADD customizations bool,
  ADD customization_setup text,
  ADD fax_port_submitted text,
  ADD router_type text,
  ADD router_make_and_model text,
  ADD network_decision text,
  ADD billing_notes text,
  ADD notes text
;

-- +goose Down
-- SQL in this section is executed when the migration is rolled back.
ALTER TABLE insys_onboarding.handoff_snapshots
  DROP COLUMN handed_off_at,

  DROP COLUMN point_of_contact,
  DROP COLUMN reason_for_purchase,
  DROP COLUMN customizations,
  DROP COLUMN customization_setup,
  DROP COLUMN fax_port_submitted,
  DROP COLUMN router_type,
  DROP COLUMN router_make_and_model,
  DROP COLUMN network_decision,
  DROP COLUMN billing_notes,
  DROP COLUMN notes
;