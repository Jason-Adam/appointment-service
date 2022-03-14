# Appointment Service  
[![Tests](https://github.com/Jason-Adam/appointment-service/actions/workflows/test.yaml/badge.svg)](https://github.com/Jason-Adam/appointment-service/actions/workflows/test.yaml)  
A service for viewing and booking appointments between a trainer and client.

## How do I run the service?  
The service can be run locally using the command:  

```bash  
make local
```  

This requires `make`, `docker-compose`, & `docker` to be installed locally.  

To spin down the resources:  

```bash  
make compose-down
```

## Running Tests  
Tests can be run with the following command:  

```bash  
make test
```  

## Calling the service  
There are currently 3 endpoints available.  

### GET Available Appointments by Trainer Between 2 Dates  

`GET /api/applications`  

This endpoint will return all available appointments for a trainer within a timespan. If no timespan is given, sensible defaults (now to now + 4 weeks) will be used.  

There are three expected query parameters:  

1. `trainerID`:int64 (required)
2. `start`: ISO8601 timestamp  
3. `end`: ISO8601 timestamp  

Valid calls to this endpoint could look like:  

```bash  
curl 'http://localhost:8080/api/appointments?trainerID=3'
curl 'http://localhost:8080/api/appointments?trainerID=3&start=2022-03-26T14:30:00-08:00'
curl 'http://localhost:8080/api/appointments?trainerID=3&start=2022-03-26T14:30:00-08:00&end=2022-03-28T14:30:00-08:00'
```  

### GET Booked Appointments by Trainer  

`GET /api/applications/:trainerID`  

This endpoint will return all `BOOKED` appointments that a trainer has. Valid calls to this endpoint look like:  

```bash  
curl http://localhost:8080/api/appointments/1
```  

### POST Appointment (Book an Appointment)  

`POST /api/appointments`  

The body of the request needs to look like the following:  

```json  
{
  "id": 1,
  "user_id": 2
}
```  

`id` is the appointment ID you are attempting to book. `user_id` is the client.  

A sample call looks like:  

```bash  
curl -X POST http://localhost:8080/api/appointments -H "Content-Type: application/json" -d '{"id": 13, "user_id": 2}'
```

## Next Steps & Notes  
* The current implementation uses PostgreSQL as it's persistance layer. The appointment bookings are done with Transactions utilizing a Repeatable Read isolation level. The database is initialized with the SQL code in `sql/structure.sql`. It contains some initial inserts to seed appointments.
* Some additional work could be done to the repository layer to make the calls a little more generic and pass in optional parameters (perhaps use Go's template package).  
* It would be nice to seed the database with default appointment records for future business days within normal business hours. The API could be extended to include an endpoint for trainers to update their availability.  
* Replace IDs (trainer, user, appointment) with UUID instead of int64.  
* Change dates to Unix Epoch time for easier manipulation.  
* More testing is always nice.
