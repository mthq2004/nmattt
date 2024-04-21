-- +migrate Up
CREATE TABLE data (
  "id" text PRIMARY KEY,
  "content" text,
  "key" text,
  "created_at" TIMESTAMPTZ NOT NULL,
  "updated_at" TIMESTAMPTZ NOT NULL
);

-- +migrate Down
DROP TABLE data;
