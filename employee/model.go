package employee

import (
	"time"
)

type Employee struct {
	ID        int
	Name      string
	Position  string
	Salary    int
	CreatedAt time.Time
	UpdatedAt time.Time
}
