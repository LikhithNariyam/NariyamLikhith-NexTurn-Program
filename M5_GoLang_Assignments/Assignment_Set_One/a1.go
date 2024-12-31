package main

import (
	"errors"
	"fmt"
)

// Define departments as constants
const (
	HR = "HR"
	IT = "IT"
)

// Define the Employee struct
type Employee struct {
	ID         int
	Name       string
	Age        int
	Department string
}

// Slice to store employee data
var employees []Employee

// Function to add an employee
func addEmployee(id int, name string, age int, department string) error {
	// Check if ID is unique
	for _, emp := range employees {
		if emp.ID == id {
			return errors.New("employee ID must be unique")
		}
	}

	// Check age
	if age <= 18 {
		return errors.New("age must be greater than 18")
	}

	// Add employee to the slice
	newEmployee := Employee{ID: id, Name: name, Age: age, Department: department}
	employees = append(employees, newEmployee)
	return nil
}

// Function to search for an employee by ID
func searchEmployeeByID(id int) (Employee, error) {
	for _, emp := range employees {
		if emp.ID == id {
			return emp, nil
		}
	}
	return Employee{}, errors.New("employee not found")
}

// Function to search for an employee by name
func searchEmployeeByName(name string) (Employee, error) {
	for _, emp := range employees {
		if emp.Name == name {
			return emp, nil
		}
	}
	return Employee{}, errors.New("employee not found")
}

// Function to list employees in a specific department
func listEmployeesByDepartment(department string) []Employee {
	var result []Employee
	for _, emp := range employees {
		if emp.Department == department {
			result = append(result, emp)
		}
	}
	return result
}

// Function to count employees in a specific department
func countEmployees(department string) int {
	count := 0
	for _, emp := range employees {
		if emp.Department == department {
			count++
		}
	}
	return count
}

// Main function
func main() {
	// Adding employees
	err := addEmployee(1, "Likhith", 25, HR)
	if err != nil {
		fmt.Println("Error:", err)
	}
	err = addEmployee(2, "Pranav", 22, IT)
	if err != nil {
		fmt.Println("Error:", err)
	}
	err = addEmployee(3, "Maurya", 17, IT) // Should return an error (age <= 18)
	if err != nil {
		fmt.Println("Error:", err)
	}

	// Searching for an employee by ID
	emp, err := searchEmployeeByID(1)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Printf("Found Employee by ID: %+v\n", emp)
	}

	// Searching for an employee by name
	emp, err = searchEmployeeByName("Pranav")
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Printf("Found Employee by Name: %+v\n", emp)
	}

	// Listing employees in IT department
	itEmployees := listEmployeesByDepartment(IT)
	fmt.Println("Employees in IT Department:", itEmployees)

	// Counting employees in HR department
	hrCount := countEmployees(HR)
	fmt.Printf("Number of employees in HR Department: %d\n", hrCount)
}
