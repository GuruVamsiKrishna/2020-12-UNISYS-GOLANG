package dao

import (
	"database/sql"
	"errors"
	"strconv"

	"github.com/kayartaya-vinod/e19-rest/src/entity"
)

func getDb() *sql.DB {
	db, err := sql.Open("mysql", "root:Welcome#123@(localhost:3306)/vindb")
	if err != nil {
		panic(err.Error())
	}
	return db
}

func GetAllEmps() []entity.Employee {
	sql := "SELECT * FROM EMPLOYEES"
	db := getDb()
	defer db.Close()
	rows, err := db.Query(sql)
	if err != nil {
		panic(err.Error())
	}

	var emps = []entity.Employee{}
	for rows.Next() {
		emp := entity.Employee{}
		rows.Scan(&emp.Id, &emp.Name, &emp.Salary)
		emps = append(emps, emp)
	}

	return emps
}

func GetOneEmployee(id int) (entity.Employee, error) {
	sql := "SELECT * FROM EMPLOYEES WHERE ID=?"
	db := getDb()
	defer db.Close()

	rows, err := db.Query(sql, id)
	if err != nil {
		panic(err.Error())
	}

	emp := entity.Employee{}
	if rows.Next() {
		rows.Scan(&emp.Id, &emp.Name, &emp.Salary)
		return emp, nil
	}

	return emp, errors.New("No data found for id " + strconv.Itoa(id))
}

func AddNewEmployee(emp entity.Employee) (entity.Employee, error) {
	sql := "INSERT INTO EMPLOYEES(NAME, SALARY) VALUES (?, ?)"
	db := getDb()
	defer db.Close()

	result, err := db.Exec(sql, emp.Name, emp.Salary)
	if err != nil {
		return emp, err
	}

	generatedId, _ := result.LastInsertId()

	emp.Id = int(generatedId)
	return emp, nil
}
