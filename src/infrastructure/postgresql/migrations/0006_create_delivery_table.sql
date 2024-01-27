CREATE TABLE IF NOT EXISTS Deliveries (
    id SERIAL PRIMARY KEY,
    driver_id int,
    sender_id int,
    destination TEXT NOT NULL,
    deadline TIMESTAMP,
    completed BOOLEAN DEFAULT FALSE,
    FOREIGN KEY (driver_id) REFERENCES Drivers (id) ON DELETE SET NULL,
    FOREIGN KEY (sender_id) REFERENCES Users (id) ON DELETE SET NULL
);
