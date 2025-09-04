-- migrate:up
ALTER TABLE
  expenses
ADD
  COLUMN external_id UUID;

-- migrate:down
ALTER TABLE
  expenses DROP COLUMN external_id;
