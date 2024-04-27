-- +migrate Up
CREATE TABLE data (
  "id" text PRIMARY KEY,
  "content" text,
  "key" text,
  "public_key" text,
  "private_key" text,
  "type" text,
  "created_at" TIMESTAMPTZ NOT NULL
);

-- +migrate Down
DROP TABLE data;
