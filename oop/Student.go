package oop

import (
	"errors"
	"fmt"
)

// private member
type student struct {
	nim string
	name string
	address string
	department string
}


// constructor
func NewStudent(nim string, name string, address string, department string) (student, error) {

	if (nim=="" || name=="" || address=="" || department=="") {
		return student{}, errors.New("Missing Parameters")
	}

	return student{
		nim: nim,
		name: name,
		address: address,
		department: department,
	}, nil
}

// getter
func (s student) GetNim() string {
	return s.nim
}

func (s student) GetName() string {
	return s.name
}

func (s student) GetAddress() string {
	return s.address
}

func (s student) GetDepartment() string {
	return s.department
}

// setter
func (s * student) SetNim(nim string) {
	s.nim = nim
}

func (s * student) SetName(name string) {
	s.name = name
}

func (s * student) SetAddress(address string) {
	s.address = address
}

func (s * student) SetDepartment(department string) {
	s.department = department
}

// method
func (s student) PrintStudent() {
	fmt.Printf("Nim : %s\nName : %s\nAddress : %s\nDepartment : %s\n", s.nim, s.name, s.address, s.department)
}