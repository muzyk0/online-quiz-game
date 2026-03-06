DROP TABLE IF EXISTS users;

-- NOTE: pgcrypto extension is NOT dropped here because it's shared with
-- later migrations (e.g., 000004_quiz_schema uses gen_random_uuid())
