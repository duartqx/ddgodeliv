CREATE TABLE IF NOT EXISTS Drivers (
    id serial PRIMARY KEY,
    user_id int REFERENCES users(id),
    company_id int REFERENCES companies(id),
    license_id VARCHAR(255) NOT NULL UNIQUE
);
