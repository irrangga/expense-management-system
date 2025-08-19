-- migrate:up
INSERT INTO
  users(email, name, role)
VALUES
  (
    'employee1@example.com',
    'Arie Untung',
    'employee'
  ),
  (
    'employee2@example.com',
    'Bandung Bondowoso',
    'employee'
  ),
  (
    'employee3@example.com',
    'Cecep Santoso',
    'employee'
  ),
  (
    'manager@example.com',
    'Dono Warkop',
    'manager'
  );

-- migrate:down
DELETE FROM
  users
WHERE
  email IN (
    'employee1@example.com',
    'employee2@example.com',
    'employee3@example.com',
    'manager@example.com'
  );
