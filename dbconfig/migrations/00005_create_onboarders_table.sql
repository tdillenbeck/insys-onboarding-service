-- +goose Up
-- SQL in this section is executed when the migration is applied.
CREATE TABLE insys_onboarding.onboarders (
  id uuid NOT NULL PRIMARY KEY,
  user_id uuid NOT NULL,
  schedule_custimization_link text, -- task id: '2d2df285-9211-48fc-a057-74f7dee2d9a4'
  schedule_porting_link text, --task id: '9aec502b-f8b8-4f10-9748-1fe4050eacde'
  schedule_network_link text,-- task id: '7b15e061-8002-4edc-9bf4-f38c6eec6364'
  schedule_software_install_link text, -- task id: '16a6dc91-ec6b-4b09-b591-a5b0dfa92932'
  schedule_phone_install_link text, -- task id: 'fd4f656c-c9f1-47b8-96ad-3080b999a843'
  schedule_software_training_link text, -- task id: 'c20b65d8-e281-4e62-98f0-4aebf83e0bee'
  schedule_phone_training_link text, -- task id: '47743fae-c775-45d5-8a51-dc7e3371dfa4'

  created_at timestamp without time zone NOT NULL DEFAULT now(),
  updated_at timestamp without time zone NOT NULL DEFAULT now()
);
CREATE UNIQUE INDEX index_onboarders_on_user_id ON insys_onboarding.onboarders USING btree (user_id);


-- +goose Down
-- SQL in this section is executed when the migration is rolled back.
DROP TABLE insys_onboarding.onboarders;
