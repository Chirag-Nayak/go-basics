package handlers

import (
	"errors"
	"log"
	"net/http"
	"regexp"
	"strconv"

	"github.com/Chirag-Nayak/go-basics/web-service/model"
	"github.com/Chirag-Nayak/go-basics/web-service/repository"
	"github.com/Chirag-Nayak/go-basics/web-service/service"
)

type GetIDFromReq func(*http.Request, *log.Logger) (int64, error)

type Employee struct {
	logger    *log.Logger
	eService  *service.Employee
	getIdFunc GetIDFromReq
}

func NewEmployee(l *log.Logger, eServ *service.Employee) *Employee {
	emp := Employee{}
	emp.eService = eServ
	emp.logger = l
	emp.getIdFunc = getIDFromRequest_golib
	return &emp
}

func (e *Employee) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		e.GetEmployees(w, r)
	} else if r.Method == http.MethodPost {
		e.AddEmployee(w, r)
	} else if r.Method == http.MethodPut {
		e.logger.Println("PUT request: ", r.URL.Path)
		e.UpdateEmployee(w, r)
	} else if r.Method == http.MethodDelete {
		w.WriteHeader(http.StatusMethodNotAllowed)
	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func (e *Employee) GetEmployees(w http.ResponseWriter, r *http.Request) {
	id, err := e.getIdFunc(r, e.logger)
	if err != nil {
		http.Error(w, "Error while getting ID from URL", http.StatusBadRequest)
		return
	}
	if id == -1 {
		// If ID is not present in the URL, return all the employees
		e.getAllEmployees(w, r)
	} else {
		e.getEmployeeByID(id, w, r)
	}
}

// Add new employee to the data store by handling POST requests
func (e *Employee) AddEmployee(w http.ResponseWriter, r *http.Request) {
	e.logger.Println("Received POST request on the employee URI.")

	emp := &model.Employee{}
	err := emp.FromJSON(r.Body)
	if err != nil {
		e.logger.Printf("Error while decoding employee data from POST request: %#+v\n", err)
		http.Error(w, "Unable to marshal JSON from POST request", http.StatusBadRequest)
		return
	}

	e.eService.AddEmployee(r.Context(), emp)
	w.WriteHeader(http.StatusCreated)
}

// Update employee informations by handling PUT request
func (e *Employee) UpdateEmployee(w http.ResponseWriter, r *http.Request) {
	e.logger.Println("Received PUT request on the employee URI.")

	// expect the id in the URI & get the ID
	id, err := e.getIdFunc(r, e.logger)
	if err != nil {
		http.Error(w, "Invalid URI", http.StatusBadRequest)
		return
	}

	emp := &model.Employee{}
	err = emp.FromJSON(r.Body)
	if err != nil {
		e.logger.Printf("Error while decoding employee data from POST request: %#+v\n", err)
		http.Error(w, "Unable to marshal JSON from POST request", http.StatusBadRequest)
		return
	}

	_, err = e.eService.UpdateEmployee(r.Context(), id, emp)
	if err == repository.ErrRecordNotFound {
		http.Error(w, "Employee with given ID not found", http.StatusNotFound)
		return
	} else if err != nil {
		http.Error(w, "Error while updating the Employee with given ID", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

// ---------------------------------------------------------------
// Package only methods
// ---------------------------------------------------------------

// Retrieve all the employees from the data store & send them as JSON response
func (e *Employee) getAllEmployees(w http.ResponseWriter, r *http.Request) {
	e.logger.Println("Received GET request on the employee URI.")
	emps, err := e.eService.GetAll(r.Context())
	if err != nil {
		e.logger.Printf("Error while getting employee list data: %#+v\n", err)
		http.Error(w, "Unable to read employee information", http.StatusInternalServerError)
		return
	}

	err = emps.ToJSON(w)
	if err != nil {
		e.logger.Printf("Error while encoding employee list data: %#+v\n", err)
		http.Error(w, "Unable to marshal employees into JSON", http.StatusInternalServerError)
	}
}

// Retrive employee information by ID
func (e *Employee) getEmployeeByID(id int64, w http.ResponseWriter, r *http.Request) {
	e.logger.Printf("Receieved GET with ID on employee URI.")
	emp, err := e.eService.GetByID(r.Context(), id)

	if err != nil {
		e.logger.Printf("Error returned from servie: %#+v\n", err)
		http.Error(w, "Unable to get employee information for given ID.", http.StatusBadRequest)
		return
	}

	err = emp.ToJSON(w)
	if err != nil {
		e.logger.Printf("Error while encoding employee information to JSON: %#+v\n", err)
		http.Error(w, "Unable to marshal JSON", http.StatusInternalServerError)
		return
	}
}

// Get employee ID from the request using built-in go libraries
func getIDFromRequest_golib(r *http.Request, l *log.Logger) (int64, error) {
	l.Println("Getting ID from the request using built-in Go Libraries.")
	reg, err := regexp.Compile(`/([0-9]+)`)
	if err != nil {
		l.Printf("Unable to compile regex, ID may not be present in the URI, err: %#+v\n", err)
		return -1, nil
	}

	g := reg.FindAllStringSubmatch(r.URL.Path, -1)

	if len(g) == 0 {
		l.Println("ID is not available in the URL")
		return -1, nil
	}

	if len(g) != 1 {
		l.Println("Invalid URI more than one id")
		return -1, errors.New("invalid URI more than one id")
	}

	if len(g[0]) != 2 {
		l.Println("Invalid URI more than one capture group")
		return -1, errors.New("invalid URI more than one capture group")
	}

	idString := g[0][1]
	// Parse the string ID as a decimal (10 base) 64 bit Integer
	id, err := strconv.ParseInt(idString, 10, 64)
	if err != nil {
		l.Printf("Invalid URI unable to convert ID:%s to int64, err: %#+v\n", idString, err)
		return -1, errors.New("invalid URI string to int64 parsing error")
	}
	return id, nil
}
