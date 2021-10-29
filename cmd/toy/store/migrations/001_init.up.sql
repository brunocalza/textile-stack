CREATE TABLE IF NOT EXISTS people (
    id INT PRIMARY KEY,
    name TEXT NOT NULL,
    email TEXT,
    pb_data BYTEA NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);