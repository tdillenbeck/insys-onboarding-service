-- How to run:
-- psql -U postgres -h localhost -p 5432 -d insys_onboarding_local -f dbconfig/seed.sql

-- Clear all existing data previous data
DELETE FROM insys_onboarding.onboarding_task_instances;
DELETE FROM insys_onboarding.onboarding_tasks;
DELETE FROM insys_onboarding.onboarding_categories;

-- Categories
-- INSERT INTO insys_onboarding.onboarding_categories VALUES (id, display_text, display_order, created_at, updated_at);
  INSERT INTO insys_onboarding.onboarding_categories VALUES ('43d6553c-608f-4b45-a351-2e98b660980a', 'First test category', 1, default, default);

-- Tasks
-- INSERT INTO insys_onboarding.onboarding_tasks VALUES (id, title, content, display_order, created_at, updated_at, onboarding_category_id, button_content);
  INSERT INTO insys_onboarding.onboarding_tasks VALUES ('934b38d9-398c-4248-ada2-f33abe9ac2b8', 'First test onboarding task', 'Lorem Ipsum', 1, default, default, '43d6553c-608f-4b45-a351-2e98b660980a', 'button lorem ipsum');

-- Task Instances
--INSERT INTO insys_onboarding.onboarding_task_instances VALUES (id, location_id, title, content, display_order, status, status_updated_at, status_updated_by, completed_at, completed_by, verified_at, verified_by, created_at, updated_at, onboarding_category_id, onboarding_task_id);

-- Both completed_at and verified_at are null
  INSERT INTO insys_onboarding.onboarding_task_instances VALUES ('c897711a-7bec-4ffd-8686-c1d18332e8d7', '38fd93c3-78d0-4ef4-9466-99421eccf600', 'First test onboarding task intstance', 'Lorem Ipsum', 0, 0, now(), 'default', null, '', null, '', default, default, '43d6553c-608f-4b45-a351-2e98b660980a', '934b38d9-398c-4248-ada2-f33abe9ac2b8', 'button lorem ipsum');

-- Only verified_at is null
  INSERT INTO insys_onboarding.onboarding_task_instances VALUES ('efc659ac-6925-43a0-bc22-0f37abc9e9c8', '38fd93c3-78d0-4ef4-9466-99421eccf600', 'First test onboarding task intstance', 'Lorem Ipsum', 2, 0, now(), 'Weave - zach', now(), 'Weave - zach', null, '', default, default, '43d6553c-608f-4b45-a351-2e98b660980a', '934b38d9-398c-4248-ada2-f33abe9ac2b8', 'button lorem ipsum');

-- Both completed_at and verified_at have values
  INSERT INTO insys_onboarding.onboarding_task_instances VALUES ('e71fd324-4075-412b-bb0d-c7bb99e5fbb8', '38fd93c3-78d0-4ef4-9466-99421eccf600', 'First test onboarding task intstance', 'Lorem Ipsum', 3, 0, now(), 'Weave - zach', now(), 'Weave - zach', now(), 'Weave - gazza', default, default, '43d6553c-608f-4b45-a351-2e98b660980a', '934b38d9-398c-4248-ada2-f33abe9ac2b8', 'button lorem ipsum');
