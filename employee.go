/*
Write a program that stores information about employees in a company.
Each employee has the following information:

	ID (a string)
	Name (a string)
	Age (an integer)
	Salary (a float64)

You should define a struct to represent an employee, and create a slice of employees.
The program should allow the user to:

	Add a new employee
	Remove an employee by ID
	Print the list of all employees, sorted by salary in descending order
*/

/*
package main

import "fmt"

	func main() {
		type emp struct {
			emp_id, emp_name string
			age              int
			salary           float64
		}

		Emp1 := emp{
			emp_id:   "ABC123",
			emp_name: "Roma",
			age:      30,
			salary:   35000,
		}
		Emp2 := emp{
			emp_id:   "XYZ123",
			emp_name: "Rajeesha",
			age:      20,
			salary:   55000,
		}
		Emp3 := emp{
			emp_id:   "DEF789",
			emp_name: "Lucian",
			age:      35,
			salary:   10000,
		}
		fmt.Println(Emp1)
		fmt.Println(Emp3)
		for _, employee := range Emp2 {
			fmt.Println(employee)
		}


		//slice := emp[]

}
*/
package main

import (
	"fmt"
	"sort"
)

type Employee struct {
	ID     string
	Name   string
	Age    int
	Salary float64
}

func main() {

	//
	employees := []Employee{
		Employee{ID: "1", Name: "John Smith", Age: 30, Salary: 45000.0},
		Employee{ID: "2", Name: "Jane Doe", Age: 25, Salary: 30000.0},
		Employee{ID: "3", Name: "Bob Johnson", Age: 35, Salary: 50000.0},
	}

	employees = append(employees, Employee{ID: "4", Name: "Alice Williams", Age: 28, Salary: 40000.0})

	id := "2"
	var indexToRemove int
	for i, employee := range employees {
		if employee.ID == id {
			indexToRemove = i
			break
		}
	}
	employees = append(employees[:indexToRemove], employees[indexToRemove+1:]...)

	sort.Slice(employees, func(i, j int) bool {
		return employees[i].Salary > employees[j].Salary
	})

	fmt.Println("Employees:")
	for _, employee := range employees {
		fmt.Printf("ID: %s, Name: %s, Age: %d, Salary: %.2f\n", employee.ID, employee.Name, employee.Age, employee.Salary)
	}
}
