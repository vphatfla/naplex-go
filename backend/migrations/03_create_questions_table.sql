-- Create the questions table
CREATE TABLE IF NOT EXISTS questions (
    id SERIAL PRIMARY KEY,
    title TEXT NOT NULL,
    question TEXT NOT NULL,
    multiple_choices TEXT NOT NULL,
    correct_answer TEXT NOT NULL,
    explanation TEXT,
    keywords TEXT,
    link TEXT
);

-- Grant permissions to our user
GRANT ALL PRIVILEGES ON TABLE questions TO naplex_user;
GRANT USAGE, SELECT ON SEQUENCE questions_id_seq TO naplex_user;
