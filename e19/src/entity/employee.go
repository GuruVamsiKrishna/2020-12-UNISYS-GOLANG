package entity

type Employee struct {
	Id     int     `xml:"emp-id" json:"id"`
	Name   string  `xml:"name" json:"name"`
	Salary float64 `xml:"salary" json:"salary"`
}
/*

// ON LINUX TABLE NAMES ARE CASE-SENSITIVE
// =======================================

CREATE TABLE EMPLOYEES(
	ID INT PRIMARY KEY AUTO_INCREMENT,
	NAME VARCHAR(50) NOT NULL,
	SALARY DOUBLE
);

INSERT INTO EMPLOYEES (NAME, SALARY) VALUES
('Suresh Kumar', 45000),
('Harish Rao', 55000),
('Veena', 75000),
('Ravi', 47000);

*/