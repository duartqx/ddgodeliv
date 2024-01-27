CREATE TABLE IF NOT EXISTS Vehicles (
    id SERIAL PRIMARY KEY,
    model_id int NOT NULL REFERENCES VehicleModels(id),
    company_id int NOT NULL REFERENCES Companies(id),
    license_id VARCHAR(255) NOT NULL UNIQUE
);
