-- +goose Up
-- +goose StatementBegin
CREATE TABLE insys_onboarding.chili_piper_schedule_events (
  id uuid NOT NULL PRIMARY KEY,

  event_id text,
  route_id text,
  assignee_id text,
  start_at timestamp with time zone,
  end_at timestamp with time zone,
  contact_id text, -- salesforce_user_id
  location_id uuid,

  created_at timestamp with time zone NOT NULL DEFAULT now(),
  updated_at timestamp with time zone NOT NULL DEFAULT now()
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE insys_onboarding.chili_piper_schedule_events;
-- +goose StatementEnd
