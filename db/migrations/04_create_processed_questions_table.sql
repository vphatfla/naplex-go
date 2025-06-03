-- Set the search path to our schema
SET search_path TO naplex_data;

-- Create the processed_questions table
CREATE TABLE IF NOT EXISTS processed_questions (
    id SERIAL PRIMARY KEY,
    title TEXT NOT NULL,
    question TEXT NOT NULL,
    multiple_choices TEXT NOT NULL,
    correct_answer TEXT NOT NULL,
    explanation TEXT,
    keywords TEXT
);

-- Grant permissions to our user
GRANT ALL PRIVILEGES ON TABLE processed_questions TO naplex_user;
GRANT USAGE, SELECT ON SEQUENCE processed_questions_id_seq TO naplex_user;
