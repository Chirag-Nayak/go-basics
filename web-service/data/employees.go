package data

import (
	"encoding/json"
	"io"
	"time"
)

type Employee struct {
	ID        int
	FirstName string
	LastName  string
	Email     string
	JoinDate  string
}

type Employees []*Employee

func GetAllEmployees() Employees {
	return empList
}

// Encoding JSON directly to an io writer is faster and efficient
// than marshalling json into memory and sending that as string
func (e *Employees) ToJSON(w io.Writer) error {
	encoder := json.NewEncoder(w)
	return encoder.Encode(e)
}

// In case of marshalling into json,
// all the data will be stored in memory and than used / written to some response
func (e *Employees) ToJSONData() ([]byte, error) {
	return json.Marshal(e)
}

// Hard coded list is used as data source for the api demo
var empList = []*Employee{
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
