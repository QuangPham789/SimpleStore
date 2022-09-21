CREATE TABLE IF NOT EXISTS accounts (
  "id" bigserial PRIMARY KEY,
  "username" varchar NOT NULL,
  "password" bigint NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);