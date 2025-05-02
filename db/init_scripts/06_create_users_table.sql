-- Set the search path
SET search_path TO naplex_data;

-- Add user table
CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    email VARCHAR(255) UNIQUE NOT NULL,
    password BYTEA NOT NULL,
    name VARCHAR(255) NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW()
);

-- Create index on user email
CREATE INDEX idx_users_email ON users(email);
