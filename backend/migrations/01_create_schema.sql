-- Create schema if it doesn't exist
CREATE SCHEMA IF NOT EXISTS naplex_data;

-- Set the search path to the new schema, this is for the postgres original user when run the init sql scripts
ALTER USER postgres SET search_path TO naplex_data;
