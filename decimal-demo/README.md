# Read Me
Demo application to show decimal usage in Go. The application demostrates mainly below functions.
* Basic decimal operations using shopspring's decimal package
* Basic CRUD operations using PostgresSQL DB for decimal type
* Basic CRUD operations using MongoDB for decimal type
	- Work In Progress / Not fully working, because of the Open Issue in shopspring's decimal package. [Issue#168](https://github.com/shopspring/decimal/issues/168)

## Prerequisite for the developlent
* Development envorinment
	- Go Lang 1.19
	
* Go modules & pacages
	- decimal (github.com/shopspring/decimal)
	- godotenv (github.com/joho/godotenv)
	- Postgress Library (github.com/lib/pq)
	- MongoDB Library (go.mongodb.org/mongo-driver)

## Executing the Application
The API is still in development phase & not production ready, so it is not yet bundled in the docker container,
to execute the API please do following steps:
* Create "app.env" in the application's root directory (that is where the go.mod file is located) to store the necessary environment variables. You can refer to [example.env](./example.env) to  check which environment variables are needed for this API.

* Start the docker containers for PostgresDB using docker compose file [docker-compose.yml](./docker-compose.yml). Do check the port numbers in the docker compose file before starting. 
```
	docker-compose -f .\docker-compose.yml up
```

* Build & start the demo application.
	- To execute the demo with PostgreSQL DB, execute the [main.go](./cmd/postgres/main.go), from the application's root directory (that is where the go.mod file is located) using below command. 
```
	go build -o demo.exe .\cmd\postgres\main.go; if($?)  { .\demo.exe }
```

	- To execute the demo with MongoDB, execute the [main.go](./cmd/mongo/main.go), from the application's root directory (that is where the go.mod file is located) using below command. 
```
	go build -o demo.exe .\cmd\mongo\main.go; if($?)  { .\demo.exe }
```
