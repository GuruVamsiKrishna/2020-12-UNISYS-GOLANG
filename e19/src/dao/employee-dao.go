package dao

import (
	"errors"
	"strconv"

	"github.com/kayartaya-vinod/e19-rest/src/entity"
)

// temp
var emps = []entity.Employee{
	{Id: 101, Name: "John Doe", Salary: 24000},
	{Id: 122, Name: "Jane Doe", Salary: 27500},
	{Id: 678, Name: "Robert Miller", Salary: 23400},
}

func GetAllEmps() []entity.Employee {
	return emps
}

func GetOneEmployee(id int) (entity.Employee, error) {
	for _, emp := range emps {
		if emp.Id == id {
			return emp, nil
		}
	}
	return entity.Employee{}, errors.New("No data found for id: " + strconv.Itoa(id))
}

func AddNewEmployee(emp entity.Employee) (entity.Employee, error) {
	for _, e := range emps {
		if e.Id == emp.Id {
			return emp, errors.New("Employee already exists with this id")
		}
	}
	emps = append(emps, emp)
	return emp, nil
}
