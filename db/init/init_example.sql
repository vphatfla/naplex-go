CREATE SCHEMA IF NOT EXISTS naplex_data;

CREATE TABLE IF NOT EXISTS naplex_data.questions (
    id SERIAL PRIMARY KEY,
    question VARCHAR(255)
);


CREATE USER app_user WITH PASSWORD 'app_password';
GRANT USAGE ON SCHEMA naplex_data TO app_user;
GRANT SELECT, INSERT, UPDATE, DELETE ON ALL TABLES IN SCHEMA naplex_data TO app_user;
GRANT USAGE, SELECT ON ALL SEQUENCES IN SCHEMA naplex_data TO app_user;
