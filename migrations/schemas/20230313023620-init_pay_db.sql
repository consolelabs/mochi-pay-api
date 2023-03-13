
-- +migrate Up
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

create table if not exists tokens (
    "id" uuid NOT NULL DEFAULT uuid_generate_v4(),
    "name" varchar NOT NULL,
    "symbol" varchar NOT NULL,
    "decimal" integer NOT NULL,
    "chain_id" text NOT NULL,
    "created_at" timestamptz DEFAULT now(),
    "updated_at" timestamptz DEFAULT now()
);

create table if not exists balances (
    "id" uuid NOT NULL DEFAULT uuid_generate_v4(),
    "profile_id" text NOT NULL,
    "token_id" text,
    "amount" text,
    "created_at" timestamptz DEFAULT now(),
    "updated_at" timestamptz DEFAULT now(),
    "deleted_at" timestamptz
);

create unique index balances_profile_id_token_id_uidx ON balances (profile_id, token_id);

create type trasfer_log_status as enum ('success', 'failed');

create table if not exists transfer_logs (
    "id" uuid NOT NULL DEFAULT uuid_generate_v4(),
    "sender_profile_id" text,
    "recipients_profile_id" varchar(255)[],
    "number_receiver" integer,
    "token_id" text,
    "amount" float8,
    "status" trasfer_log_status,
    "created_at" timestamptz DEFAULT now(),
    "updated_at" timestamptz DEFAULT now(),
    "deleted_at" timestamptz,
    "note" text
);
-- +migrate Down
drop index if exists balances_profile_id_token_id_uidx;
drop table if exists tokens;
drop table if exists balances;
drop table if exists transfer_logs;
drop type if exists trasfer_log_status;