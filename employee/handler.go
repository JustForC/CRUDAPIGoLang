package employee

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type handler struct {
	service Service
}

func NewHandler(service Service) *handler {
	return &handler{service}
}

func (h *handler) New(c *gin.Context) {
	var request EmployeeRequest

	err := c.ShouldBindJSON(&request)

	if err != nil {
		log.Fatal(err)
	}

	employee, _ := h.service.Create(request)

	var response EmployeeResponse

	response.Name = employee.Name
	response.Position = employee.Position
	response.Salary = employee.Salary

	c.JSON(http.StatusOK, gin.H{
		"data": response,
	})
}

func (h *handler) Get(c *gin.Context) {
	employees, _ := h.service.GetAll()

	c.JSON(http.StatusOK, gin.H{
		"data": employees,
	})
}

func (h *handler) Find(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	employee, _ := h.service.GetID(id)

	var response EmployeeResponse

	response.Name = employee.Name
	response.Position = employee.Position
	response.Salary = employee.Salary

	c.JSON(http.StatusOK, gin.H{
		"data": response,
	})
}

func (h *handler) Put(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	var request EmployeeRequest

	err := c.ShouldBindJSON(&request)

	if err != nil {
		log.Fatal(err)
	}

	employee, _ := h.service.Update(id, request)

	var response EmployeeResponse

	response.Name = employee.Name
	response.Position = employee.Position
	response.Salary = employee.Salary

	c.JSON(http.StatusOK, gin.H{
		"data": response,
	})
}

func (h *handler) Destroy(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	employee, _ := h.service.Delete(id)

	var response EmployeeResponse

	response.Name = employee.Name
	response.Position = employee.Position
	response.Salary = employee.Salary

	c.JSON(http.StatusOK, gin.H{
		"data": response,
	})
}
