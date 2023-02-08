package model

import (
	"encoding/json"
	"io"
)

type Employee struct {
	ID        int64  `json:"id"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
	JoinDate  string `json:"joinDate"`
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
