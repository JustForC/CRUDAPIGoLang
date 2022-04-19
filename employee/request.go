package employee

import "encoding/json"

type EmployeeRequest struct {
	Name     string      `json:"name" binding:"required"`
	Position string      `json:"position" binding:"required"`
	Salary   json.Number `json:"salary" binding:"required,number"`
}
