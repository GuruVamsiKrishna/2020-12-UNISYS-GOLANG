package handlers

import (
	"log"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/kayartaya-vinod/e19-rest/src/entity"
	"github.com/kayartaya-vinod/e19-rest/src/mysql/dao"
)

func GetAllEmployeesHandler(c *gin.Context) {
	emps := dao.GetAllEmps()

	switch c.GetHeader("Accept") {
	case "application/json":
		c.JSON(200, emps)
	case "application/xml":
		c.XML(200, emps)
	default:
		c.Status(406)
	}
}

func GetEmployeeByIdHandler(c *gin.Context) {
	strId := c.Param("id")
	id, err := strconv.Atoi(strId)
	if err != nil {
		c.JSON(400, gin.H{"message": "Invalid id format"})
	} else {
		emp, err := dao.GetOneEmployee(id)
		if err != nil {
			c.JSON(404, gin.H{"message": err.Error()})
		} else {
			c.JSON(200, emp)
		}
	}
}

func AddNewEmployeeHandler(c *gin.Context) {
	emp := entity.Employee{}
	log.Printf("BEFORE: emp = %v\n", emp)
	c.BindJSON(&emp)
	log.Printf("AFTER: emp = %v\n", emp)
	emp, err := dao.AddNewEmployee(emp)
	if err != nil {
		c.JSON(400, gin.H{"message": err.Error()})
	} else {
		c.JSON(201, emp)
	}
}
