# go.grades #

Simple distributed system in Go from 'Building Distributed Applications with Go' by Mike Van Sickle on Pluralsight.

Uses the Go standard library only.

Has the following components:
- Service Registry: Responsible for registering and de-registering services & health monitoring.
- Portal: A very basic web application that is essentially an API gateway to backend services.
- Log Service: Centralised logging.
- Grading Service: Business logic for grading. Uses service discovery. In memory data persistence.

This is not meant to be a production ready system. It is a learning exercise.

Each component has a server.go file where most of the web server code is located. The server.go file
contains some common code starts each web service.


Might look at re-building this using something like https://github.com/go-micro/go-micro


## To Run ##

Make sure to start the registry service first.

Run the registry service: go build grades/cmd/registryservice && ./registryservice

Run the log service: go build grades/cmd/logservice && ./logservice

Run the grading service: go build grades/cmd/gradingservice && ./gradingservice

Run the teacher portal: go build grades/cmd/teacherportal && ./teacherportal

Send a request to the log service: curl -XPOST -d "Log this" http://localhost:4000/log

The teacher portal is at: http://localhost:5000

