CREATE TABLE IF NOT EXISTS VehicleModels (
    id SERIAL PRIMARY KEY,
    manufacturer VARCHAR(255) NOT NULL,
    year INT NOT NULL,
    name VARCHAR(30) NOT NULL,
    transmission VARCHAR(30) NOT NULL,
    type VARCHAR(30) NOT NULL
);
