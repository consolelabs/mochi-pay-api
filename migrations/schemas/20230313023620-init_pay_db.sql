
-- +migrate Up
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

create table if not exists token (
    "id" uuid NOT NULL DEFAULT uuid_generate_v4(),
    "name" varchar NOT NULL,
    "symbol" varchar NOT NULL,
    "decimal" integer NOT NULL,
    "chain_id" text NOT NULL,
    "created_at" timestamptz DEFAULT now(),
    "updated_at" timestamptz DEFAULT now()
);

create table if not exists balance (
    "id" uuid NOT NULL DEFAULT uuid_generate_v4(),
    "profile_id" text NOT NULL,
    "token_id" text,
    "amount" float8,
    "created_at" timestamptz DEFAULT now(),
    "updated_at" timestamptz DEFAULT now(),
    "deleted_at" timestamptz
);

create unique index balance_profile_id_token_id_uidx ON balance (profile_id, token_id);

create table if not exists activity_log (
    "id" uuid NOT NULL DEFAULT uuid_generate_v4(),
    "profile_id" text,
    "receiver" varchar(255)[],
    "number_receiver" integer,
    "token_id" text,
    "amount" float8,
    "status" varchar,
    "created_at" timestamptz DEFAULT now(),
    "updated_at" timestamptz DEFAULT now(),
    "deleted_at" timestamptz,
    "note" text
);
-- +migrate Down
drop index if exists balance_profile_id_token_id_uidx;
drop table if exists token;
drop table if exists balance;
drop table if exists activity_log;