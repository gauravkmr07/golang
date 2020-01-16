package models

import (
	"database/sql"
	"errors"

	db "github.com/user/app/databases"
	e "github.com/user/app/entity"
)

//ListEmployees -
func ListEmployees(page, limit int) ([]e.Employee, error) {
	var employee e.Employee
	var emplist []e.Employee

	page = (page*limit - limit)

	rows, err := db.DB.Query("SELECT * FROM employees LIMIT ?, ?", page, limit)

	if err != nil {
		return emplist, err
	}

	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&employee.ID, &employee.Name, &employee.Email, &employee.Phone)
		if err != nil {
			return emplist, err
		}
		emplist = append(emplist, employee)
	}

	return emplist, nil

}

//CountEmployees -
func CountEmployees() (int, error) {
	var totalEmps int
	err := db.DB.QueryRow("SELECT COUNT(id) AS `total` FROM `employees`").Scan(&totalEmps)

	if err == sql.ErrNoRows {
		return 0, errors.New("Zero results found")
	}

	if err != nil {
		return 0, err
	}

	return totalEmps, nil
}

// UpdateEmployee -
func UpdateEmployee(emp *e.Employee) (int64, error) {

	const query = `UPDATE employees SET name=?, email=?,  phone=? WHERE id=?`

	tx, err := db.DB.Begin()
	if err != nil {
		return 0, err
	}

	result, err := tx.Exec(query, emp.Name, emp.Email, emp.Phone, emp.ID)
	if err != nil {
		tx.Rollback()
		return 0, err
	}

	affectedRows, err := result.RowsAffected()
	if err != nil {
		tx.Rollback()
		return 0, err
	}

	if affectedRows == 0 {
		tx.Rollback()
		return 0, errors.New("No record found")
	}

	tx.Commit()

	return affectedRows, nil

}

// CreateNewEmployee -
func CreateNewEmployee(emp *e.Employee) (int64, error) {

	const query = "INSERT INTO employees(name, email, phone) VALUES (?, ?, ?)"

	insForm, err := db.DB.Prepare(query)

	if err != nil {
		return 0, err
	}
	var res sql.Result
	res, err = insForm.Exec(emp.Name, emp.Email, emp.Phone)

	if err != nil {
		return 0, err
	}

	id, err := res.LastInsertId()

	if err != nil {
		return 0, err
	}

	return id, nil
}
