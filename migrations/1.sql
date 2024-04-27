-- +migrate Up
CREATE TABLE data (
  "id" text PRIMARY KEY,
  "type" text,
  "encrypted_content" text,
  "content" text,
  "key" text,
  "public_key" text,
  "private_key" text,
  "created_at" TIMESTAMPTZ NOT NULL
);

-- +migrate Down
DROP TABLE data;
