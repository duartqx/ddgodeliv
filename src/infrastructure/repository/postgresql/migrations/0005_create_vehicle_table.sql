CREATE TABLE IF NOT EXISTS Vehicles (
    id SERIAL PRIMARY KEY,
    model_id int NOT NULL,
    company_id int NOT NULL,
    license_id VARCHAR(255) NOT NULL UNIQUE,
    FOREIGN KEY (model_id) REFERENCES VehicleModels (id) ON DELETE CASCADE,
    FOREIGN KEY (company_id) REFERENCES Companies (id) ON DELETE CASCADE
);
