package controllers

import (
	"log"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
	e "github.com/user/app/entity"
	m "github.com/user/app/models"
)

//UpdateEmployee --
func UpdateEmployee(c echo.Context) error {

	var emp e.Employee
	id, err := strconv.Atoi(c.Param("id"))

	if id <= 0 || err != nil {
		return c.JSON(http.StatusBadRequest, e.SetResponse(http.StatusBadRequest, "Please enter a valid employee id", nil))
	}

	err = c.Bind(&emp)

	if err != nil {
		return c.JSON(http.StatusBadRequest, e.SetResponse(http.StatusBadRequest, err.Error(), emp))
	}

	emp.ID = int64(id)

	_, err = m.UpdateEmployee(&emp)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, e.SetResponse(http.StatusUnprocessableEntity, err.Error(), nil))
	}

	return c.JSON(http.StatusOK, e.SetResponse(http.StatusOK, "employee updated successfully", emp))

}

//EmptyValue --
var EmptyValue = make([]int, 0)

// CreateEmployee to create a new employee
func CreateEmployee(c echo.Context) error {
	var employee e.Employee
	var id int64

	err := c.Bind(&employee)

	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, err.Error())
	}

	id, err = m.CreateNewEmployee(&employee)
	if err != nil {
		log.Println("**************Error" + err.Error())
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	employee.ID = id

	return c.JSON(http.StatusOK, e.SetResponse(http.StatusOK, "employee created Successfully", employee))

}

// ListEmployee -
func ListEmployee(c echo.Context) error {

	var page, limit string

	page = c.QueryParam("page")
	limit = c.QueryParam("limit")

	if page == "" {
		page = "1"
	}
	if limit == "" {
		limit = "10"
	}

	pagenum, err := strconv.Atoi(page)
	limitnum, err := strconv.Atoi(limit)

	if err != nil || pagenum <= 0 || limitnum <= 0 {
		return c.JSON(http.StatusBadRequest, e.SetResponse(http.StatusBadRequest, "page and limit values must be of type integer greater than 0", EmptyValue))
	}

	employees, err := m.ListEmployees(pagenum, limitnum)

	if err != nil {
		return c.JSON(http.StatusBadRequest, e.SetResponse(http.StatusBadRequest, err.Error(), EmptyValue))
	}

	return c.JSON(http.StatusOK, e.SetResponse(http.StatusOK, "employees listed successfully", employees))
}
