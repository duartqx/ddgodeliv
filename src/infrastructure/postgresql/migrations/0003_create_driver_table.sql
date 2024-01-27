CREATE TABLE IF NOT EXISTS Drivers (
    id SERIAL PRIMARY KEY,
    user_id int NOT NULL,
    company_id int NOT NULL,
    license_id VARCHAR(255) NOT NULL UNIQUE,
    FOREIGN KEY (user_id) REFERENCES Users (id) ON DELETE CASCADE,
    FOREIGN KEY (company_id) REFERENCES Companies (id) ON DELETE CASCADE
);
