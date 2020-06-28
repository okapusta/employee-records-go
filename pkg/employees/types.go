package employees

type Args struct {
	Command    string
	EmployeeID int64
}

type EmployeesApiResponse struct {
	Status  string     `json:"status"`
	Data    []Employee `json:"data"`
	Message string     `json:message`
}

type EmployeeApiResponse struct {
	Status  string   `json:"status"`
	Data    Employee `json:"data"`
	Message string   `json:message`
}

type Employee struct {
	ID             string `json:"id,omitempty"`
	EmployeeName   string `json:"employee_name,omitempty"`
	EmployeeSalary string `json:"employee_salary,omitempty"`
	EmployeeAge    string `json:"employee_age,omitempty"`
	ProfileImage   string `json:"profile_image,omitempty"`
}

type EmployeeList struct {
	Employees []Employee
}
