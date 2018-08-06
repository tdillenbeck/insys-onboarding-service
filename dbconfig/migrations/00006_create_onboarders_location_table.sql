-- +goose Up
-- SQL in this section is executed when the migration is applied.
CREATE TABLE insys_onboarding.onboarders_location (
  id uuid NOT NULL PRIMARY KEY,
  onboarder_id uuid NOT NULL REFERENCES onboarders(id),
  location_id uuid NOT NULL,

  created_at timestamp without time zone NOT NULL DEFAULT now(),
  updated_at timestamp without time zone NOT NULL DEFAULT now()
);
CREATE UNIQUE INDEX index_onboarders_location_on_location_id ON insys_onboarding.onboarders_location USING btree (location_id);

-- +goose Down
-- SQL in this section is executed when the migration is rolled back.
DROP TABLE insys_onboarding.onboarders_location;
