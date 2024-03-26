-- +goose Up
-- +goose StatementBegin

BEGIN;
CREATE SCHEMA IF NOT EXISTS GophKepper;
CREATE TABLE  IF NOT EXISTS GophKepper.users
(
    id    serial PRIMARY KEY,
    username   varchar(255) UNIQUE NOT NULL,
    password_hash   text NOT NULL,
    created_at timestamp NOT NULL DEFAULT now(),
    updated_at  timestamp NOT NULL DEFAULT now()
);

CREATE OR REPLACE TRIGGER timestamp_trigger_users
    BEFORE UPDATE ON GophKepper.users
    FOR EACH ROW
EXECUTE FUNCTION trigger_timestamp()
;

CREATE TABLE IF NOT EXISTS GophKepper.pairs
(
    id         SERIAL PRIMARY KEY,
    user_id	   INT REFERENCES GophKepper.users (id) ON DELETE CASCADE,
    login      VARCHAR NOT NULL,
    password_hash   VARCHAR NOT NULL,
    nonce_hex text NOT NULL,
    metadata   TEXT,
    created_at TIMESTAMP  NOT NULL DEFAULT NOW(),
    updated_at  timestamp NOT NULL DEFAULT NOW()
);
CREATE OR REPLACE TRIGGER timestamp_trigger_pairs
    BEFORE UPDATE ON GophKepper.pairs
    FOR EACH ROW
EXECUTE FUNCTION trigger_timestamp()
;

CREATE TABLE IF NOT EXISTS GophKepper.text
(
    id         SERIAL PRIMARY KEY,
    user_id    INT REFERENCES GophKepper.users (id) ON DELETE CASCADE,
    note       TEXT,
    metadata   TEXT,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at  timestamp NOT NULL DEFAULT NOW()
);

CREATE OR REPLACE TRIGGER timestamp_trigger_text
    BEFORE UPDATE ON GophKepper.text
    FOR EACH ROW
EXECUTE FUNCTION trigger_timestamp()
;

CREATE TABLE  IF NOT EXISTS GophKepper.bank
(
    id     serial PRIMARY KEY,
    user_id     int       NOT NULL references GophKepper.users (id) on delete cascade,
    number        int     NOT NULL,
    card_holder varchar NOT NULL,
    expiration_date timestamp NOT NULL,
    security_code varchar NOT NULL,
    nonce_hex text NOT NULL,
    metadata jsonb,
    created_at  timestamp NOT NULL DEFAULT now(),
    updated_at  timestamp NOT NULL DEFAULT now()
);

CREATE OR REPLACE TRIGGER timestamp_trigger_bank
    BEFORE UPDATE ON GophKepper.bank
    FOR EACH ROW
EXECUTE FUNCTION trigger_timestamp();

CREATE INDEX  IF NOT EXISTS idx_username_id ON GophKepper.users USING brin (id);

CREATE UNIQUE INDEX  IF NOT EXISTS idx_username_unique ON GophKepper.users (username);

CREATE INDEX  IF NOT EXISTS idx_pairs_id ON GophKepper.pairs USING brin (id);

CREATE INDEX  IF NOT EXISTS idx_pairs_user_id ON GophKepper.pairs  (user_id);

CREATE INDEX  IF NOT EXISTS idx_username_text  ON GophKepper.text USING brin (id) ;

CREATE INDEX  IF NOT EXISTS idx_entity_text_user_id ON GophKepper.text (user_id);

CREATE INDEX  IF NOT EXISTS idx_bank_id ON GophKepper.bank USING brin (id);

CREATE INDEX  IF NOT EXISTS idx_bank_user_id ON GophKepper.bank (user_id);

commit;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
BEGIN;
DROP SCHEMA GophKepper CASCADE ;
DROP TABLE IF EXISTS users CASCADE;
DROP TABLE IF EXISTS text CASCADE;


commit;
-- +goose StatementEnd
