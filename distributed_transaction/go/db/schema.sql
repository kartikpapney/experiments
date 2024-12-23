-- Drop tables if they exist to start fresh
DROP TABLE IF EXISTS seats;
DROP TABLE IF EXISTS trips;
DROP TABLE IF EXISTS flights;
DROP TABLE IF EXISTS airlines;
DROP TABLE IF EXISTS users;

-- Table: User
-- Stores information about users.
CREATE TABLE users (
    id SERIAL PRIMARY KEY,  -- Unique identifier for each user, auto-incremented
    name VARCHAR(255) NOT NULL  -- Name of the user
);

-- Table: Airline
-- Represents airline companies.
CREATE TABLE airlines (
    id SERIAL PRIMARY KEY,  -- Unique identifier for each airline
    name VARCHAR(255) NOT NULL  -- Name of the airline
);

-- Table: Flight
-- Defines flights operated by airlines.
CREATE TABLE flights (
    id SERIAL PRIMARY KEY,  -- Unique identifier for each flight
    airline_id INT NOT NULL,  -- Foreign key referencing the `airlines` table
    name VARCHAR(255) NOT NULL,  -- Name or code of the flight (e.g., "AA101")
    FOREIGN KEY (airline_id) REFERENCES airlines(id) ON DELETE CASCADE  -- Foreign key constraint
);

-- Table: Trip
-- Represents scheduled trips for flights.
CREATE TABLE trips (
    id SERIAL PRIMARY KEY,  -- Unique identifier for each trip
    flight_id INT NOT NULL,  -- Foreign key referencing the `flights` table
    start_time TIMESTAMP NOT NULL,  -- Scheduled start time of the trip
    end_time TIMESTAMP NOT NULL,  -- Scheduled end time of the trip
    FOREIGN KEY (flight_id) REFERENCES flights(id) ON DELETE CASCADE  -- Foreign key constraint
);

-- Table: Seat
-- Tracks seat reservations for users on trips.
CREATE TABLE seats (
    id SERIAL PRIMARY KEY,  -- Unique identifier for each seat
    name VARCHAR(10) NOT NULL,  -- Seat label (e.g., "12A")
    trip_id INT NOT NULL,  -- Foreign key referencing the `trips` table
    user_id INT,  -- Foreign key referencing the `users` table (nullable for available seats)
    FOREIGN KEY (trip_id) REFERENCES trips(id) ON DELETE CASCADE,  -- Foreign key constraint
    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE SET NULL  -- Foreign key constraint (nullable)
);