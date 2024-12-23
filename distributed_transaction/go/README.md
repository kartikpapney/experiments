# Flight Booking System

## Overview
This project is a representation of how to handle concurrent request during **Airline Checkine**

---

## Database Schema

### Tables and Relationships

1. **User**
   - Stores information about users.
   - Columns:
     - `id`: Unique identifier for each user.
     - `name`: Name of the user.

2. **Airline**
   - Represents airline companies.
   - Columns:
     - `id`: Unique identifier for each airline.
     - `name`: Name of the airline.

3. **Flight**
   - Defines flights operated by airlines.
   - Columns:
     - `id`: Unique identifier for each flight.
     - `airline_id`: References the `Airline` table.
     - `name`: Name or code of the flight (e.g., "AA101").

4. **Trip**
   - Represents scheduled trips for flights.
   - Columns:
     - `id`: Unique identifier for each trip.
     - `flight_id`: References the `Flight` table.
     - `start_time`: Scheduled start time of the trip.
     - `end_time`: Scheduled end time of the trip.

5. **Seat**
   - Tracks seat reservations for users on trips.
   - Columns:
     - `id`: Unique identifier for each seat.
     - `name`: Seat label (e.g., "12A").
     - `trip_id`: References the `Trip` table.
     - `user_id`: References the `User` table.

---

## Relationships
- **Airline ↔ Flight**: One airline can operate multiple flights.
- **Flight ↔ Trip**: A flight can have multiple scheduled trips.
- **Trip ↔ Seat**: A trip can have multiple seats available for booking.
- **Seat ↔ User**: A seat is reserved by a user.

---

## Getting Started
1. Clone the repository:
   ```bash
   git clone https://github.com/yourusername/flight-booking-system.git
