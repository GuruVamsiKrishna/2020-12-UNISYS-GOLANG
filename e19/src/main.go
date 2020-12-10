package main

import (
	"log"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/kayartaya-vinod/e19-rest/src/handlers"
)

func main() {

	router := gin.Default()

	router.GET("/api/employees", handlers.GetAllEmployeesHandler)

	router.GET("/api/employees/:id", handlers.GetEmployeeByIdHandler)

	router.POST("/api/employees", handlers.AddNewEmployeeHandler)

	log.Println("Starting up REST server....")
	router.Run(":7070") // port defaults to 8080
}

/*
	ReST --> Representational State Transfer

	Diff representations --> XML/JSON/text/html/CSV

	request headers allow a client to negotiate for the kind of data
	expected (Accept) or being sent (Content-Type)

	HTTP METHODS --> actions GET/POST/PUT/PATCH/DELETE
	GET --> retrieve one or more resources (state) from the server
	POST --> create one ore more resources on the server
	PUT --> update one ore more resources on the server
	PATCH --> partial update one or more resources on the server
	DELETE --> delete one resource on the server

	URI --> identifies a resource
	http://..../api/employees

	GET / --> all employees
	GET /12 --> get a single employee with id 12
	POST / --> carries a payload (JSON) and adds a new employee to the exisiting resource
	PUT /12 --> carries a payload and replaces the existing employee#12 with the payload
	PATCH /12 --> carries a payload (partial content) and updates existing employee#12
	DELETE /12 --> delete the employee#12

	{
		"name": "Vinod",
		"age": 47,
		"isMarried": true,
		"address": {
			"city": "Bangalore",
			"phones": [ "9731424784", "9844083934" ]
		}
	}

	http://localhost:7070/blog
	http://localhost:7070/blog/2020
	http://localhost:7070/blog/2020/12
	http://localhost:7070/blog/2020/12/10

	r.GET("/blog", h1)
	r.GET("/blog/:year", h2)
	r.GET("/blog/:year/:month", h3)
	r.GET("/blog/:year/:month/:date", h4)

	c.Params --> {2020, 12, 10} --> {year: 2020, month: 12, date: 10}

*/
