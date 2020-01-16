package entity

// Employee entity will be used to define the json format and available fields in Employee
type Employee struct {
	ID    int64  `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	Phone int    `json:"phone"`
}
