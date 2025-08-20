-- migrate:up
CREATE TABLE expenses(
  id BIGSERIAL PRIMARY KEY,
  user_id BIGINT NOT NULL,
  amount_idr BIGINT NOT NULL,
  description TEXT NOT NULL,
  receipt_url TEXT,
  status TEXT,
  submitted_at TIMESTAMPTZ NOT NULL DEFAULT now(),
  processed_at TIMESTAMPTZ NOT NULL DEFAULT now(),
  FOREIGN KEY (user_id) REFERENCES users (id) ON DELETE CASCADE
);

-- migrate:down
DROP TABLE expenses;
