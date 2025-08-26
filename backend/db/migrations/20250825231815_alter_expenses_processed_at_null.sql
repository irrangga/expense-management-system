-- migrate:up
ALTER TABLE
  expenses
ALTER COLUMN
  processed_at DROP NOT NULL,
ALTER COLUMN
  processed_at DROP DEFAULT;

-- migrate:down
ALTER TABLE
  expenses
ALTER COLUMN
  processed_at
SET
  NOT NULL,
ALTER COLUMN
  processed_at
SET
  DEFAULT now();
