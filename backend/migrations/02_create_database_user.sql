-- Create a new user (change 'naplex_user' and 'secure_password' as needed)
CREATE USER naplex_user WITH PASSWORD 'password';

-- Grant schema usage permission
GRANT USAGE ON SCHEMA naplex_data TO naplex_user;

-- Grant all privileges on all tables in schema (current and future)
GRANT ALL PRIVILEGES ON ALL TABLES IN SCHEMA naplex_data TO naplex_user;
GRANT ALL PRIVILEGES ON ALL SEQUENCES IN SCHEMA naplex_data TO naplex_user;

-- Allow the user to create new objects in the schema
GRANT CREATE ON SCHEMA naplex_data TO naplex_user;

-- Set default privileges for future tables
ALTER DEFAULT PRIVILEGES IN SCHEMA naplex_data
GRANT ALL PRIVILEGES ON TABLES TO naplex_user;

ALTER DEFAULT PRIVILEGES IN SCHEMA naplex_data
GRANT ALL PRIVILEGES ON SEQUENCES TO naplex_user;

-- Set search_path permanently for the user
ALTER USER naplex_user SET search_path TO naplex_data;
