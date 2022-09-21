// schema.sql
drop table if exists accounts;

CREATE TABLE "accounts" (
  "id" bigserial PRIMARY KEY,
  "username" varchar NOT NULL,
  "password" bigint NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);