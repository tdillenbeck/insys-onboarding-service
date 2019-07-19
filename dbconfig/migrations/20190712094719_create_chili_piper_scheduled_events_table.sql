-- +goose Up
-- +goose StatementBegin
CREATE TABLE insys_onboarding.chili_piper_schedule_events (
  id uuid NOT NULL PRIMARY KEY,
  location_id uuid NOT NULL,

  event_id text,
  event_type text,
  route_id text,
  assignee_id text,
  contact_id text, -- salesforce_user_id

  start_at timestamp with time zone,
  end_at timestamp with time zone,

  created_at timestamp with time zone NOT NULL DEFAULT now(),
  updated_at timestamp with time zone NOT NULL DEFAULT now()
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE insys_onboarding.chili_piper_schedule_events;
-- +goose StatementEnd
