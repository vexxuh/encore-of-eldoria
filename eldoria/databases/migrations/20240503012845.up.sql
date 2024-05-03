-- create "inventories" table
CREATE TABLE "inventories" (
  "id" bigserial NOT NULL,
  "created_at" timestamptz NULL,
  "updated_at" timestamptz NULL,
  "deleted_at" timestamptz NULL,
  "i_apple" bigint NULL,
  "i_potion" bigint NULL,
  "i_potion_plus" bigint NULL,
  "c_gold" bigint NULL,
  "b_gold" bigint NULL,
  PRIMARY KEY ("id")
);
-- create index "idx_inventories_deleted_at" to table: "inventories"
CREATE INDEX "idx_inventories_deleted_at" ON "inventories" ("deleted_at");
-- create "characters" table
CREATE TABLE "characters" (
  "id" bigserial NOT NULL,
  "created_at" timestamptz NULL,
  "updated_at" timestamptz NULL,
  "deleted_at" timestamptz NULL,
  "iventory_id" bigint NULL,
  "username" text NULL,
  "user" text NULL,
  "c_level" bigint NULL,
  "c_health" bigint NULL,
  "m_health" bigint NULL,
  "b_health" bigint NULL,
  "s_strength" bigint NULL,
  "s_agility" bigint NULL,
  "s_constitution" bigint NULL,
  "s_intelligence" bigint NULL,
  "s_wisdom" bigint NULL,
  "w_s_sword" bigint NULL,
  "w_s_axe" bigint NULL,
  "w_s_spear" bigint NULL,
  "p_state" text NULL,
  "c_area" text NULL,
  "c_e_weapon" bigint NULL,
  "c_e_armor" bigint NULL,
  PRIMARY KEY ("id"),
  CONSTRAINT "fk_characters_inventory" FOREIGN KEY ("iventory_id") REFERENCES "inventories" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION
);
-- create index "idx_characters_deleted_at" to table: "characters"
CREATE INDEX "idx_characters_deleted_at" ON "characters" ("deleted_at");
-- drop "attacks" table
DROP TABLE "attacks";
