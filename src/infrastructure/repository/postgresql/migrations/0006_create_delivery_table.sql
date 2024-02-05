CREATE TABLE IF NOT EXISTS Deliveries (
    id SERIAL PRIMARY KEY,
    loadout VARCHAR(255) NOT NULL,
    weight INT NOT NULL,
    driver_id INT,
    sender_id INT,
    origin TEXT NOT NULL,
    destination TEXT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deadline TIMESTAMP,
    status SMALLINT DEFAULT 0,
    FOREIGN KEY (driver_id) REFERENCES Drivers (id) ON DELETE SET NULL,
    FOREIGN KEY (sender_id) REFERENCES Users (id) ON DELETE SET NULL
);
