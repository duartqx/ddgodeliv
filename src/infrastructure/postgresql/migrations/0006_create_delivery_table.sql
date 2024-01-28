CREATE TABLE IF NOT EXISTS Deliveries (
    id SERIAL PRIMARY KEY,
    driver_id int,
    sender_id int,
    origin TEXT NOT NULL,
    destination TEXT NOT NULL,
    deadline TIMESTAMP,
    status SMALLINT DEFAULT 0,
    FOREIGN KEY (driver_id) REFERENCES Drivers (id) ON DELETE SET NULL,
    FOREIGN KEY (sender_id) REFERENCES Users (id) ON DELETE SET NULL
);
