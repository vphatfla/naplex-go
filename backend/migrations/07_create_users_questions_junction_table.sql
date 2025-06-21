-- Create enum type for the status
CREATE TYPE question_status AS ENUM ('FAILED', 'PASSED', 'NA');

-- Create the table
CREATE TABLE users_questions (
    uid INT NOT NULL,
    qid INT NOT NULL,
    status question_status DEFAULT 'NA',
    attempts INT DEFAULT 0,
    saved BOOLEAN DEFAULT FALSE,
    hidden BOOLEAN DEFAULT FALSE,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    -- composite PK
    PRIMARY KEY (uid, qid),
    -- FK
    CONSTRAINT fk_user FOREIGN KEY (uid) REFERENCES users(id) ON DELETE CASCADE,
    CONSTRAINT fk_question FOREIGN KEY (qid) REFERENCES questions(id) ON DELETE CASCADE
);

-- Trigger functions for updated_at
CREATE OR REPLACE FUNCTION update_updated_at_column()
RETURNS TRIGGER AS $$
BEGIN
    NEW.updated_at = NOW();
    RETURN NEW;
END;
$$ language 'plpgsql';

-- Create trigger to automatically update updated_at
CREATE TRIGGER update_users_questions_updated_at
    BEFORE UPDATE ON users_questions
    FOR EACH ROW
    EXECUTE FUNCTION update_updated_at_column();

-- Create indexes for better query performance
CREATE INDEX idx_users_questions_uid ON users_questions(uid);
CREATE INDEX idx_users_questions_qid ON users_questions(qid);
CREATE INDEX idx_users_questions_status ON users_questions(status);
CREATE INDEX idx_users_questions_saved ON users_questions(saved) WHERE saved = TRUE;
CREATE INDEX idx_users_questions_hidden ON users_questions(hidden) WHERE hidden = FALSE;
