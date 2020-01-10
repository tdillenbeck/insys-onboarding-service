-- +goose Up
CREATE TABLE insys_onboarding.reschedule_tracking
(
    id uuid NOT NULL PRIMARY KEY,
    location_id uuid NOT NULL,
    event_type text NOT NULL,
    rescheduled_events_count int NOT NULL,
    rescheduled_events_calculated_at timestamp with time zone NOT NULL,
    created_at timestamp with time zone NOT NULL DEFAULT now(),
    updated_at timestamp with time zone NOT NULL DEFAULT now()
);
-- +goose Down
DROP TABLE IF EXISTS insys_onboarding.reschedule_tracking;
