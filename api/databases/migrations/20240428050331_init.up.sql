-- create "attacks" table
CREATE TABLE "attacks" (
  "id" bigserial NOT NULL,
  "created_at" timestamptz NULL,
  "updated_at" timestamptz NULL,
  "deleted_at" timestamptz NULL,
  "attack_type" text NULL,
  PRIMARY KEY ("id")
);
-- create index "idx_attacks_deleted_at" to table: "attacks"
CREATE INDEX "idx_attacks_deleted_at" ON "attacks" ("deleted_at");
