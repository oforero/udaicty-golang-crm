# Udaicty Golang-CRM

This is the submission to the GOLANG CRM project by Oscar Forero

## Contents

The application is structured in various packages:

* db: Has the DB code, right now only supports SQLite, but could be extended to use PostgreSQL
* model: It has the Customer struct and the intializers to populate a Database with 3 initial records
* main.go: The entry point for the application
* main_test.go: The integration tests for the application, they run using SQLite
* css & scripts: Are static folders used to serve the API documentation
* docs: The Swagger generated API documentation, which is also served statically.

## How to run

To be able to run the application from the source code, the first step is to install all the required libraries:

```
$ go get -u -v -f all
```

This command will install all the dependencies to be able to build and run the application.

Tests can be run in the usual GO way:

```
$ go test
```

Running the application will start the server in port 3000, and will serve both the API and the documentation:

```
$ go run main.go
```

With the application running the API can be accessed using Postman or any other HTTP client, and the documentation witha Browser going to ** http://localhost:3000/ **