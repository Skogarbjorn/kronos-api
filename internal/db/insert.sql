INSERT INTO profile (kt, first_name, last_name)
VALUES (1907032070, 'ragnar bjoern', 'ingvarsson');

INSERT INTO profile (kt, first_name, last_name)
VALUES (1907032071, 'gamer epic', 'super');

INSERT INTO profile_pin_auth (profile_id, pin)
VALUES (1, '0311a8ebc9de5629f286b5861092cc55c0a63d2d24e1aa2f7aec3da9b2de41d9');

INSERT INTO profile_pin_auth (profile_id, pin)
VALUES (2, '0311a8ebc9de5629f286b5861092cc55c0a63d2d24e1aa2f7aec3da9b2de41d9');

INSERT INTO workspace (name)
VALUES ('test workspace');

INSERT INTO location (name, address, workspace_id)
VALUES ('test location', 'test address', 1);

INSERT INTO location (name, address, workspace_id)
VALUES ('test location 2', 'test address 2', 1);

INSERT INTO company (name, workspace_id)
VALUES ('test company', 1);

INSERT INTO company (name, workspace_id)
VALUES ('test company 2', 1);

INSERT INTO contract (hourly_rate, unpaid_lunch_minutes)
VALUES (4500, 30);

INSERT INTO employment (profile_id, company_id, contract_id, role, end_date)
VALUES (1, 1, 1, 'worker', now() + interval '30 days');

INSERT INTO employment (profile_id, company_id, contract_id, role, end_date)
VALUES (2, 2, 1, 'worker', now() + interval '30 days');

INSERT INTO task (location_id, company_id, name, description, is_completed)
VALUES (1, 1, 'test task', 'this is a dummy test task', false);

INSERT INTO task (location_id, company_id, name, description, is_completed)
VALUES (1, 1, 'test task 2', 'gamer', false);

INSERT INTO task (location_id, company_id, name, description, is_completed)
VALUES (2, 1, 'test task 3', 'epic', false);

INSERT INTO task (location_id, company_id, name, description, is_completed)
VALUES (1, 2, 'test task 4', 'super', false);
