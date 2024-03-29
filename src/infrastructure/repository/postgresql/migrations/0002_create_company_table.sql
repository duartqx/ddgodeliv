CREATE TABLE IF NOT EXISTS Companies (
    id SERIAL PRIMARY KEY,
    owner_id int NOT NULL,
    name VARCHAR(255) NOT NULL UNIQUE,
    FOREIGN KEY (owner_id) REFERENCES Users (id) ON DELETE CASCADE
);
