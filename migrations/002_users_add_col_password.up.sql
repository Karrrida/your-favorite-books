ALTER TABLE IF EXISTS users
    ADD COLUMN IF NOT EXISTS password_hash varchar(255);
