package employee

type Service interface {
	Create(employeeRequest EmployeeRequest) (Employee, error)
	GetAll() ([]Employee, error)
	GetID(ID int) (Employee, error)
	Update(ID int, employeeRequest EmployeeRequest) (Employee, error)
	Delete(ID int) (Employee, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) Create(employeeRequest EmployeeRequest) (Employee, error) {
	salary, _ := employeeRequest.Salary.Int64()

	employee := Employee{
		Name:     employeeRequest.Name,
		Position: employeeRequest.Position,
		Salary:   int(salary),
	}

	return s.repository.Create(employee)
}

func (s *service) GetAll() ([]Employee, error) {
	return s.repository.GetAll()
}

func (s *service) GetID(ID int) (Employee, error) {
	return s.repository.GetID(ID)
}

func (s *service) Update(ID int, employeeRequest EmployeeRequest) (Employee, error) {
	employee, _ := s.repository.GetID(ID)
	salary, _ := employeeRequest.Salary.Int64()

	employee.Name = employeeRequest.Name
	employee.Position = employeeRequest.Position
	employee.Salary = int(salary)

	return s.repository.Update(employee)
}

func (s *service) Delete(ID int) (Employee, error) {
	employee, _ := s.repository.GetID(ID)

	return s.repository.Delete(employee)
}
