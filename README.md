# go.grades #

Simple distributed system using Go only.

Has the following components:
- Service Registry: Responsible for registering and de-registering services & health monitoring.
- Portal: A very basic web application that is essentially an API gateway to backend services.
- Log Service: Centralised logging.
- Grading Service: Business logic for grading. In memory data persistence.

This is not meant to be a production ready system. It is a learning exercise.

Might look at re-building this using something like https://github.com/go-micro/go-micro


## To Run ##

Run the log service: go build grades/cmd/logservice && ./logservice

Send a request to the log service: curl -XPOST -d "Log this" http://localhost:4000/log

