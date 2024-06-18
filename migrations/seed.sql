CREATE TABLE users (
    id serial PRIMARY KEY,
    email VARCHAR(50),
    password_hash VARCHAR(50)
);