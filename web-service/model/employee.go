package model

import (
	"encoding/json"
	"io"
)

// Employee defines an employee structure
// swagger:model
type Employee struct {
	// Extensions:
	// x-order: "0"
	ID int64 `json:"id"`

	// Extensions:
	// x-order: "1"
	FirstName string `json:"firstName"`

	// Extensions:
	// x-order: "2"
	LastName string `json:"lastName"`

	// Extensions:
	// x-order: "3"
	Email string `json:"email"`

	// Extensions:
	// x-order: "4"
	JoinDate string `json:"joinDate"`
}

type Employees []*Employee

func (e *Employee) ToJSON(w io.Writer) error {
	encoder := json.NewEncoder(w)
	return encoder.Encode(e)
}

func (e *Employees) ToJSON(w io.Writer) error {
	encoder := json.NewEncoder(w)
	return encoder.Encode(e)
}

func (e *Employee) FromJSON(r io.Reader) error {
	decoder := json.NewDecoder(r)
	return decoder.Decode(e)
}

func (e *Employees) FromJSON(r io.Reader) error {
	decoder := json.NewDecoder(r)
	return decoder.Decode(e)
}
