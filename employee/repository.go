package employee

import "gorm.io/gorm"

type Repository interface {
	Create(employee Employee) (Employee, error)
	GetAll() ([]Employee, error)
	GetID(ID int) (Employee, error)
	Update(employee Employee) (Employee, error)
	Delete(employee Employee) (Employee, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepositoy(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) Create(employee Employee) (Employee, error) {
	err := r.db.Create(&employee).Error

	return employee, err
}

func (r *repository) GetAll() ([]Employee, error) {
	var employees []Employee

	err := r.db.Find(&employees).Error

	return employees, err
}

func (r *repository) GetID(ID int) (Employee, error) {
	var employee Employee

	err := r.db.Find(&employee, ID).Error

	return employee, err
}

func (r *repository) Update(employee Employee) (Employee, error) {
	err := r.db.Save(&employee).Error

	return employee, err
}

func (r *repository) Delete(employee Employee) (Employee, error) {
	err := r.db.Delete(employee).Error

	return employee, err
}
