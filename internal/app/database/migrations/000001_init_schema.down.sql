DROP TABLE IF EXISTS users;
-- Note: pgcrypto extension is NOT dropped here because migration 000004
-- also depends on gen_random_uuid(). Drop the extension only after all
-- dependent migrations have been rolled back.
