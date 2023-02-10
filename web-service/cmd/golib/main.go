package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/Chirag-Nayak/go-basics/web-service/handlers"
	"github.com/Chirag-Nayak/go-basics/web-service/repository"
	"github.com/Chirag-Nayak/go-basics/web-service/service"
)

func main() {

	logger := log.New(os.Stdout, "demo-api ", log.LstdFlags)
	hHandler := handlers.NewHello(logger)
	gHandler := handlers.NewGoodbye(logger)

	// Initialize service & repository to be used for Employee handlers
	eRepo := repository.NewEmployeeImplInMemory(logger)
	eService := service.NewEmployee(logger, eRepo)

	// Create employee handler
	eHandler := handlers.NewEmployee(logger, eService)

	// Create a ServerMux to handle http reuests,
	// ServeMux also implements http handler interface,
	// so it can be used with the default http server (with default options) in go as follows
	// ```` http.ListenAndServe(":9090", sMux)
	sMux := http.NewServeMux()
	sMux.Handle("/", hHandler)
	sMux.Handle("/goodbye", gHandler)
	sMux.Handle("/employee/", eHandler)

	server := &http.Server{
		Addr:         ":9090",           // Configure the bind address
		Handler:      sMux,              // Set the default handler
		ErrorLog:     logger,            // Set the logger for server
		ReadTimeout:  5 * time.Second,   // Max time to read request from client
		WriteTimeout: 10 * time.Second,  // Max time to write response to the client
		IdleTimeout:  120 * time.Second, // Max time for connections using TCP Keep-Alive
	}

	// Start a server without blocking the call
	go func() {
		logger.Printf("Starting demo web-api server on 9090 port ! ! !")
		err := server.ListenAndServe()
		if err != nil {
			logger.Fatal(err)
		}
	}()

	// Listen for os.Interrupt & os.Kill signals form OS
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt)
	signal.Notify(sigChan, syscall.SIGTERM)

	// Once any of the signal is recieved in the sigChan, close server & close the application
	sig := <-sigChan
	logger.Println("Sinal received: ", sig)

	tc, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	logger.Println("Shutting down the server.")
	server.Shutdown(tc)
}
