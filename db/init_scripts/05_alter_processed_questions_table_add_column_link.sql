-- Set the search path
SET search_path TO naplex_data;

-- Add a new column into the table to keep track of the link
ALTER TABLE processed_questions
ADD COLUMN link TEXT;
