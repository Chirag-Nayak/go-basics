package repository

import (
	"context"
	"log"
	"time"

	"github.com/Chirag-Nayak/go-basics/web-service/model"
)

type EmployeeImplInMemory struct {
	logger *log.Logger
}

func NewEmployeeImplInMemory(l *log.Logger) *EmployeeImplInMemory {
	return &EmployeeImplInMemory{
		logger: l,
	}
}

// Get all the employees information from data store
func (e *EmployeeImplInMemory) GetAll(ctx context.Context) (model.Employees, error) {
	// Since this is In-Memory data store, we do not need to consider context
	// We can directly return the current list contents
	return empList, nil
}

// Get an employee information by ID
func (e *EmployeeImplInMemory) GetByID(ctx context.Context, id int64) (*model.Employee, error) {
	// Context will be ignored since this is In-Memory data store
	for _, emp := range empList {
		if emp.ID == id {
			return emp, nil
		}
	}
	return nil, ErrRecordNotFound
}

// Add new employee information to the data store
func (e *EmployeeImplInMemory) AddEmployee(ctx context.Context, emp *model.Employee) (*model.Employee, error) {
	// Context will be ignored & there should be no errors as this is In-memory data store
	emp.ID = getNextID()
	empList = append(empList, emp)
	return emp, nil
}

func (e *EmployeeImplInMemory) UpdateEmployee(ctx context.Context, id int64, emp *model.Employee) (*model.Employee, error) {
	// Since this is in memory data store, find the employee's index & update the in memory list
	_, pos, err := findEmployee(id)
	if err != nil {
		return nil, err
	}

	emp.ID = id
	empList[pos] = emp

	return emp, nil
}

// ---------------------------------------------------------------
// Package only methods
// ---------------------------------------------------------------

// Dummy ID generation logic for the demo purpose.
func getNextID() int64 {
	lastEmp := empList[len(empList)-1]
	return lastEmp.ID + 1
}

// Find the employee from list of employees
func findEmployee(id int64) (*model.Employee, int64, error) {
	for i, p := range empList {
		if p.ID == id {
			return p, int64(i), nil
		}
	}
	return nil, -1, ErrRecordNotFound
}

// Hard coded list is used as data source for the api demo
var empList = []*model.Employee{
	{
		ID:        1,
		FirstName: "FirstName1",
		LastName:  "LastName1",
		Email:     "email-1@fakedomain.com",
		JoinDate:  time.Now().UTC().String(),
	},
	{
		ID:        2,
		FirstName: "FirstName2",
		LastName:  "LastName2",
		Email:     "email-2@fakedomain.com",
		JoinDate:  time.Now().UTC().String(),
	},
}
