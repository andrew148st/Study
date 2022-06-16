package main

import "fmt"

type employee struct {
	name       string
	surname    string
	age        int
	occupation string
	salary     int
	inhabitant string
}

func main() {
	emp := employee{
		name: "Andrew", surname: "Stepanchenko", age: 52, occupation: "developer", salary: 3000, inhabitant: "St. Petersburg"}
	fmt.Printf("Name: %s %s\n", emp.name, emp.surname)
	fmt.Printf("Age: %d y.o.\n", emp.age)
	fmt.Printf("Occupation: %s\n", emp.occupation)
	fmt.Printf("Salary: %d dollars\n", emp.salary)
	fmt.Printf("Inhabitant: %s\n", emp.inhabitant)
}

//, emp.salary, emp.inhabitant
