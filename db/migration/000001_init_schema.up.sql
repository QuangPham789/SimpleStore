CREATE TABLE "accounts" (
  "id" bigserial PRIMARY KEY,
  "username" varchar NOT NULL,
  "password" bigint NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "products" (
  "id" bigserial PRIMARY KEY,
  "code" varchar NOT NULL,
  "name" varchar NOT NULL,
  "description" varchar,
  "category" varchar,
  "price" bigserial,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);


CREATE INDEX ON "products" ("code");

CREATE INDEX ON "accounts" ("username");