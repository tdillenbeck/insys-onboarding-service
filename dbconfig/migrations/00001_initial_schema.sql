-- +goose Up
CREATE TABLE insys_onboarding.onboarding_categories (
  id uuid NOT NULL PRIMARY KEY,
  display_text text NOT NULL,
  display_order integer NOT NULL,

  created_at timestamp without time zone NOT NULL DEFAULT now(),
  updated_at timestamp without time zone NOT NULL DEFAULT now()
);


CREATE TABLE insys_onboarding.onboarding_tasks (
  id uuid NOT NULL PRIMARY KEY,
  title text NOT NULL,
  content text NOT NULL,
  display_order integer NOT NULL,

  created_at timestamp without time zone NOT NULL DEFAULT now(),
  updated_at timestamp without time zone NOT NULL DEFAULT now(),

  onboarding_category_id uuid NOT NULL REFERENCES onboarding_categories(id)
);


CREATE TABLE insys_onboarding.onboarding_task_instances (
  id uuid NOT NULL PRIMARY KEY,
  location_id uuid NOT NULL,
  title text NOT NULL,
  content text NOT NULL,
  display_order integer NOT NULL,
  status integer NOT NULL,
  status_updated_at timestamp without time zone NOT NULL,
  status_updated_by text,
  completed_at timestamp without time zone,
  completed_by text,
  verified_at timestamp without time zone,
  verified_by text,

  created_at timestamp without time zone NOT NULL DEFAULT now(),
  updated_at timestamp without time zone NOT NULL DEFAULT now(),

  onboarding_category_id uuid NOT NULL REFERENCES onboarding_categories(id),
  onboarding_task_id uuid NOT NULL REFERENCES onboarding_tasks(id)
);

-- +goose Down
DROP TABLE insys_onboarding.onboarding_task_instances;
DROP TABLE insys_onboarding.onboarding_tasks;
DROP TABLE insys_onboarding.onboarding_categories;
