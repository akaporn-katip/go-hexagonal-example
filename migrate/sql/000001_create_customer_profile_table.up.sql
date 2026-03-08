CREATE TABLE customer_profile (
    id TEXT PRIMARY KEY,
    title TEXT,
    first_name TEXT,
    last_name TEXT,
    email TEXT UNIQUE,
    date_of_birth DATE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);