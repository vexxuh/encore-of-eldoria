-- reverse: drop "attacks" table
CREATE TABLE "attacks" (
  "id" bigserial NOT NULL,
  "created_at" timestamptz NULL,
  "updated_at" timestamptz NULL,
  "deleted_at" timestamptz NULL,
  "attack_type" text NULL,
  PRIMARY KEY ("id")
);
CREATE INDEX "idx_attacks_deleted_at" ON "attacks" ("deleted_at");
-- reverse: create index "idx_characters_deleted_at" to table: "characters"
DROP INDEX "idx_characters_deleted_at";
-- reverse: create "characters" table
DROP TABLE "characters";
-- reverse: create index "idx_inventories_deleted_at" to table: "inventories"
DROP INDEX "idx_inventories_deleted_at";
-- reverse: create "inventories" table
DROP TABLE "inventories";
