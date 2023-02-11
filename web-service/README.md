# Read Me
Demo application to show how you can create a web service (REST API end-point) in Go. The application mainly demonstrates: 
* Creating REST API using Built-In Go Libraries (Using http handlers)
* Creating REST API using Gorilla framework (Utilizing the same http handlers by using a modified version of "Template Method" design pattern by having a function type inside a struct)

## Prerequisite for the developlent
* Development envorinment
	- Go Lang 1.19
	
* Go modules & pacages
	- Gorilla (github.com/gorilla/mux)

## Executing the Application
The API is still in development phase & not production ready, so it is not yet bundled in the docker container,
to execute the API please follow below steps:

* Build & Execute the demo of REST API using Built-In Go Libraries.
	- Execute the [main.go](./cmd/golib/main.go), from the application's root directory (that is where the go.mod file is located) using below command. 
```
	go build -o web-service.exe .\cmd\golib\main.go; if($?)  { .\web-service.exe }
```

* Build & Execute the demo of REST API using Gorilla framework.
	- Execute the [main.go](./cmd/gorilla/main.go), from the application's root directory (that is where the go.mod file is located) using below command. 
```
	go build -o web-service.exe .\cmd\gorilla\main.go; if($?)  { .\web-service.exe }
```

## Testing the Application
You can use any 3rd party tool of your choice, I am using Postman for my testing. And I have also uploaded the Postman's collection [here](./TestEmployeeAPI.postman_collection.json) in case if you want to test using Postman.