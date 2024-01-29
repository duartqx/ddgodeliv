CREATE TABLE IF NOT EXISTS VehicleModels (
    id SERIAL PRIMARY KEY,
    manufacturer VARCHAR(255) NOT NULL,
    year INT NOT NULL,
    max_load INT NOT NULL
);
