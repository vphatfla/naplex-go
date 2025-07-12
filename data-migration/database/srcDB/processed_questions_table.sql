-- Create the processed_questions table
CREATE TABLE IF NOT EXISTS processed_questions (
    id SERIAL PRIMARY KEY,
    title TEXT NOT NULL,
    question TEXT NOT NULL,
    multiple_choices TEXT NOT NULL,
    correct_answer TEXT NOT NULL,
    explanation TEXT,
    keywords TEXT,
    link TEXT
);
