CREATE TABLE "employee" (
  "id" bigserial PRIMARY KEY,
  "firstname" varchar NOT NULL,
  "lastname" varchar NOT NULL,
  "password" varchar NOT NULL,
  "role" varchar NOT NULL,
  "created_at" timestamptz DEFAULT 'now()',
  "is_admin" boolean NOT NULL
);

CREATE TABLE "supplier" (
  "id" bigserial PRIMARY KEY,
  "name" varchar NOT NULL,
  "company" varchar NOT NULL,
  "password" varchar NOT NULL,
  "created_at" timestamptz DEFAULT 'now()'
);

CREATE TABLE "coffee" (
  "id" bigserial PRIMARY KEY,
  "type" varchar NOT NULL,
  "quantity" int NOT NULL,
  "buyed_at" timestamptz DEFAULT 'now()',
  "stocked_at" timestamptz DEFAULT 'now()',
  "is_outstocked" boolean NOT NULL
);

CREATE TABLE "machine" (
  "id" integer PRIMARY KEY,
  "sector" varchar NOT NULL,
  "company" varchar NOT NULL,
  "coffee_id" bigint NOT NULL,
  "last_restocked_at" timestamptz DEFAULT 'now()'
);

CREATE TABLE "logs" (
  "id" bigserial PRIMARY KEY,
  "from_employee" bigint NOT NULL,
  "coffee" bigint NOT NULL,
  "made_at" timestamptz DEFAULT 'now()'
);

CREATE TABLE "stocklogs" (
  "id" bigserial PRIMARY KEY,
  "from_supplier" bigint NOT NULL,
  "from_employee" bigint NOT NULL,
  "coffee" bigint NOT NULL,
  "made_at" timestamptz DEFAULT 'now()'
);

CREATE INDEX ON "coffee" ("type");

CREATE INDEX ON "machine" ("sector");

ALTER TABLE "machine" ADD FOREIGN KEY ("coffee_id") REFERENCES "coffee" ("id");

ALTER TABLE "logs" ADD FOREIGN KEY ("from_employee") REFERENCES "employee" ("id");

ALTER TABLE "logs" ADD FOREIGN KEY ("coffee") REFERENCES "coffee" ("id");

ALTER TABLE "stocklogs" ADD FOREIGN KEY ("from_supplier") REFERENCES "supplier" ("id");

ALTER TABLE "stocklogs" ADD FOREIGN KEY ("coffee") REFERENCES "coffee" ("id");

ALTER TABLE "stocklogs" ADD FOREIGN KEY ("from_employee") REFERENCES "employee" ("id");
