-- Set the search path to our schema
SET search_path TO naplex_data;

-- Create the raw_questions table
CREATE TABLE IF NOT EXISTS raw_questions (
    id SERIAL PRIMARY KEY,
    title TEXT NOT NULL,
    raw_question TEXT NOT NULL,
    link TEXT
);

-- Grant permissions to our user
GRANT ALL PRIVILEGES ON TABLE raw_questions TO naplex_user;
GRANT USAGE, SELECT ON SEQUENCE raw_questions_id_seq TO naplex_user;
