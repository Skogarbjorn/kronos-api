INSERT INTO profile (kt, first_name, last_name)
VALUES (1, 'Ragnar Björn', 'Ingvarsson');

INSERT INTO profile (kt, first_name, last_name)
VALUES (2, 'Mímir Freyr', 'Gunnarsson');

INSERT INTO profile_pin_auth (profile_id, pin)
VALUES (1, '0311a8ebc9de5629f286b5861092cc55c0a63d2d24e1aa2f7aec3da9b2de41d9');

INSERT INTO profile_pin_auth (profile_id, pin)
VALUES (2, '0311a8ebc9de5629f286b5861092cc55c0a63d2d24e1aa2f7aec3da9b2de41d9');

INSERT INTO workspace (name)
VALUES ('Sample Workspace');

INSERT INTO location (name, address, workspace_id)
VALUES ('Sample Location', 'Sample Address', 1);

INSERT INTO location (name, address, workspace_id)
VALUES ('Sample Location 2', 'Sample Address 2', 1);

INSERT INTO company (name, workspace_id)
VALUES ('Sample Company', 1);

INSERT INTO company (name, workspace_id)
VALUES ('Sample Company 2', 1);

INSERT INTO contract (hourly_rate, unpaid_lunch_minutes)
VALUES (4500, 30);

INSERT INTO employment (profile_id, company_id, contract_id, role, end_date)
VALUES (1, 1, 1, 'worker', now() + interval '30 days');

INSERT INTO employment (profile_id, company_id, contract_id, role, end_date)
VALUES (2, 2, 1, 'worker', now() + interval '30 days');

INSERT INTO task (location_id, company_id, name, description, is_completed)
VALUES (1, 1, 'Sample Task', 'This is a dummy sample task', false);

INSERT INTO task (location_id, company_id, name, description, is_completed)
VALUES (1, 1, 'Sample Task 2', 'Some other description', false);

INSERT INTO task (location_id, company_id, name, description, is_completed)
VALUES (2, 1, 'Sample Task 3', 'I was born in the darkness', false);

INSERT INTO task (location_id, company_id, name, description, is_completed)
VALUES (1, 2, 'Sample Task 4', 'Lorem Ipsum', false);
