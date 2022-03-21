package employee

import "fmt"

type Employee struct {
	name       string
	age        int
	salary     int
	isFulltime bool
}

func EmployeeLogic() {
	jack := Employee{
		name:       "Jack Smith",
		age:        27,
		salary:     2000,
		isFulltime: false,
	}

	jill := Employee{
		name:       "Jill Smith",
		age:        27,
		salary:     2000,
		isFulltime: true,
	}

	var employees []Employee
	employees = append(employees, jack)
	employees = append(employees, jill)

	for _, x := range employees {
		if x.age > 30 && x.age < 50 {
			fmt.Println(x.name, "is 30 older")
		} else {
			fmt.Println(x.name, "is under 30")
		}
	}

}
