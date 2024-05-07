-- create "characters" table
CREATE TABLE "characters" (
  "id" bigserial NOT NULL,
  "created_at" timestamptz NULL,
  "updated_at" timestamptz NULL,
  "deleted_at" timestamptz NULL,
  "username" text NULL,
  "user" text NULL,
  "c_level" bigint NULL,
  "c_experience" bigint NULL,
  "c_health" bigint NULL,
  "m_health" bigint NULL,
  "b_health" bigint NULL,
  "s_strength" bigint NULL,
  "s_agility" bigint NULL,
  "s_constitution" bigint NULL,
  "s_intelligence" bigint NULL,
  "s_wisdom" bigint NULL,
  "ws_melee" bigint NULL,
  "we_melee" bigint NULL,
  "ws_sword" bigint NULL,
  "we_sword" bigint NULL,
  "ws_axe" bigint NULL,
  "we_axe" bigint NULL,
  "ws_spear" bigint NULL,
  "we_spear" bigint NULL,
  "p_state" text NULL,
  "c_area" text NULL,
  "ce_weapon" bigint NULL,
  "ce_armor" bigint NULL,
  PRIMARY KEY ("id")
);
-- create index "idx_characters_deleted_at" to table: "characters"
CREATE INDEX "idx_characters_deleted_at" ON "characters" ("deleted_at");
-- create "armors" table
CREATE TABLE "armors" (
  "id" bigserial NOT NULL,
  "created_at" timestamptz NULL,
  "updated_at" timestamptz NULL,
  "deleted_at" timestamptz NULL,
  "character_id" bigint NULL,
  "i_name" text NULL,
  "i_id" bigint NULL,
  "i_attack" bigint NULL,
  "i_strength" bigint NULL,
  "idefense" bigint NULL,
  "i_agility" bigint NULL,
  "i_constitution" bigint NULL,
  "i_type" text NULL,
  PRIMARY KEY ("id"),
  CONSTRAINT "fk_characters_armor" FOREIGN KEY ("character_id") REFERENCES "characters" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION
);
-- create index "idx_armors_deleted_at" to table: "armors"
CREATE INDEX "idx_armors_deleted_at" ON "armors" ("deleted_at");
-- create "creatures" table
CREATE TABLE "creatures" (
  "id" bigserial NOT NULL,
  "created_at" timestamptz NULL,
  "updated_at" timestamptz NULL,
  "deleted_at" timestamptz NULL,
  "character_id" bigint NULL,
  "c_name" text NULL,
  "c_id" bigint NULL,
  "c_level" bigint NULL,
  "c_experience" bigint NULL,
  "cc_health" bigint NULL,
  "cm_health" bigint NULL,
  "c_attack" bigint NULL,
  "c_defense" bigint NULL,
  PRIMARY KEY ("id"),
  CONSTRAINT "fk_characters_creature" FOREIGN KEY ("character_id") REFERENCES "characters" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION
);
-- create index "idx_creatures_deleted_at" to table: "creatures"
CREATE INDEX "idx_creatures_deleted_at" ON "creatures" ("deleted_at");
-- create "inventories" table
CREATE TABLE "inventories" (
  "id" bigserial NOT NULL,
  "created_at" timestamptz NULL,
  "updated_at" timestamptz NULL,
  "deleted_at" timestamptz NULL,
  "character_id" bigint NULL,
  "i_apple" bigint NULL,
  "ipotion" bigint NULL,
  "ipotionplus" bigint NULL,
  "c_gold" bigint NULL,
  "b_gold" bigint NULL,
  PRIMARY KEY ("id"),
  CONSTRAINT "fk_characters_inventory" FOREIGN KEY ("character_id") REFERENCES "characters" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION
);
-- create index "idx_inventories_deleted_at" to table: "inventories"
CREATE INDEX "idx_inventories_deleted_at" ON "inventories" ("deleted_at");
-- create "weapons" table
CREATE TABLE "weapons" (
  "id" bigserial NOT NULL,
  "created_at" timestamptz NULL,
  "updated_at" timestamptz NULL,
  "deleted_at" timestamptz NULL,
  "character_id" bigint NULL,
  "i_name" text NULL,
  "i_id" bigint NULL,
  "i_attack" bigint NULL,
  "i_strength" bigint NULL,
  "idefense" bigint NULL,
  "i_agility" bigint NULL,
  "i_constitution" bigint NULL,
  "i_type" text NULL,
  PRIMARY KEY ("id"),
  CONSTRAINT "fk_characters_weapon" FOREIGN KEY ("character_id") REFERENCES "characters" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION
);
-- create index "idx_weapons_deleted_at" to table: "weapons"
CREATE INDEX "idx_weapons_deleted_at" ON "weapons" ("deleted_at");
