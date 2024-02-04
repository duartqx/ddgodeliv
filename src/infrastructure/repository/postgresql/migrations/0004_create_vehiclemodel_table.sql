CREATE TABLE IF NOT EXISTS VehicleModels (
    id SERIAL PRIMARY KEY,
    name VARCHAR(55) NOT NULL,
    manufacturer VARCHAR(255) NOT NULL,
    year INT NOT NULL,
    max_load INT NOT NULL
);
