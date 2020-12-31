-- Your SQL goes here
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE UNITS (
    id text PRIMARY KEY DEFAULT uuid_generate_v4(),
    "name" text NOT NULL,
    element int NOT NULL,
    "type" int NOT NULL
);

CREATE TABLE FIGHTERS (
    id text PRIMARY KEY DEFAULT uuid_generate_v4(),
    "user_id" text NOT NULL, -- no fk as that data is not 
    -- stored in this db, but we associate it with a user still
    unit_id text NOT NULL,
    "level" int NOT NULL DEFAULT 0,
    CONSTRAINT fk_fighter_fighter_unit
      FOREIGN KEY(unit_id) 
	      REFERENCES UNITS(id)
        ON DELETE CASCADE
);

CREATE TABLE MOVES (
  id text PRIMARY KEY DEFAULT uuid_generate_v4(),
  "name" text NOT NULL,
  element int NOT NULL,
  targets int NOT NULL,
  "type" int NOT NULL
);

CREATE TABLE UNIT_MOVES (
  id text PRIMARY KEY DEFAULT uuid_generate_v4(),
  unit_id text NOT NULL,
  move_id text NOT NULL,
  level_learned_at int NOT NULL DEFAULT 0,
  CONSTRAINT fk_fighter_moves_unit
      FOREIGN KEY(unit_id)
	      REFERENCES UNITS(id)
        ON DELETE CASCADE,
  CONSTRAINT fk_fighter_moves_move
      FOREIGN KEY(move_id) 
	      REFERENCES MOVES(id)
        ON DELETE CASCADE
)