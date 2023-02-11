package handlers

import (
	"errors"
	"log"
	"net/http"
	"strconv"

	"github.com/Chirag-Nayak/go-basics/web-service/service"
	"github.com/gorilla/mux"
)

func NewGorillaEmployee(l *log.Logger, eServ *service.Employee) *Employee {
	return &Employee{
		logger:    l,
		eService:  eServ,
		getIdFunc: getIDFromRequest_gorilla,
	}
}

// ---------------------------------------------------------------
// Package only methods
// ---------------------------------------------------------------

// Get employee ID from the request using gorilla framework
func getIDFromRequest_gorilla(r *http.Request, l *log.Logger) (int64, error) {
	l.Println("Getting ID from the request using Gorilla.")

	vars := mux.Vars(r)

	if idString, containsID := vars["id"]; containsID {
		// Parse the string ID as a decimal (10 base) 64 bit Integer
		id, err := strconv.ParseInt(idString, 10, 64)
		if err != nil {
			l.Printf("Invalid URI unable to convert ID:%s to int64, err: %#+v\n", idString, err)
			return -1, errors.New("invalid URI string to int64 parsing error")
		}
		return id, nil
	} else {
		l.Printf("ID is not available in the URL.")
		return -1, nil
	}
}
