package employees

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
)

type Args struct {
	Command    string
	EmployeeID int64
}

type EmployeesApiResponse struct {
	Status string     `json:"status"`
	Data   []Employee `json:"data"`
}

type EmployeeApiResponse struct {
	Status string   `json:"status"`
	Data   Employee `json:"data"`
}

type Employee struct {
	ID             string `json:"id"`
	EmployeeName   string `json:"employee_name"`
	EmployeeSalary string `json:"employee_salary"`
	EmployeeAge    string `json:"employee_age"`
	ProfileImage   string `json:"profile_image"`
}

type Employees struct {
	Employees []Employee
}

const EMPLOYEES_URL = "http://dummy.restapiexample.com/api/v1/employees"
const EMPLOYEE_URL = "http://dummy.restapiexample.com/api/v1/employee/%s"

var validArgs = []string{"list", "show"}

func Run(args []string) {
	parsedArgs, err := parseArgs(args)
	if err != nil {
		fmt.Printf("[Error] %v", err)
	}
	switch parsedArgs.Command {
	case "list":
		if resp, err := listEmployees(); err != nil {
			fmt.Printf("[Error] %v", err)
		} else {
			resp.Print()
		}
	case "show":
		if resp, err := showEmployee(parsedArgs.EmployeeID); err != nil {
			fmt.Printf("[Error] %v", err)
		} else {
			resp.Print()
		}
	}
}

func parseArgs(args []string) (Args, error) {
	if len(args) < 2 {
		return Args{}, fmt.Errorf("You must provide arguments")
	}
	valid := false
	command := args[1]
	var employeeID int64
	for _, arg := range validArgs {
		if command == arg {
			valid = true
		}
	}
	if !valid {
		return Args{}, fmt.Errorf("Invalid Command: %s", command)
	}
	if command == "show" {
		if len(args) < 3 {
			return Args{}, fmt.Errorf("You must provide employee ID")
		}
		var err error
		employeeID, err = strconv.ParseInt(args[2], 10, 64)
		if err != nil {
			return Args{}, fmt.Errorf("Invalid argument - must be a number %#v", args[2])
		}
	}
	return Args{
		Command:    command,
		EmployeeID: employeeID,
	}, nil
}

func listEmployees() (Employees, error) {
	resp, err := http.Get(EMPLOYEES_URL)
	if err != nil {
		return Employees{}, err
	}
	responseData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return Employees{}, err
	}
	var apiResponse EmployeesApiResponse
	json.Unmarshal(responseData, &apiResponse)
	if apiResponse.Status != "success" {
		return Employees{}, fmt.Errorf("Invalid API response")
	}
	return Employees{Employees: apiResponse.Data}, nil
}

func showEmployee(employeeID int64) (Employee, error) {
	employeeUrl := fmt.Sprintf(EMPLOYEE_URL, strconv.FormatInt(employeeID, 10))
	resp, err := http.Get(employeeUrl)
	if err != nil {
		return Employee{}, nil
	}
	responseData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return Employee{}, err
	}
	var apiResponse EmployeeApiResponse
	json.Unmarshal(responseData, &apiResponse)
	if apiResponse.Status != "success" {
		return Employee{}, fmt.Errorf("Invalid API response")
	}
	return apiResponse.Data, nil
}

func (employees Employees) Print() {
	for _, employee := range employees.Employees {
		employee.Print()
		fmt.Printf("--------------------------------\n\n")
	}
}

func (employee Employee) Print() {
	fmt.Printf("Employee ID: %s\n", employee.ID)
	fmt.Printf("Employee Name: %s\n", employee.EmployeeName)
	fmt.Printf("Employee Salary: %s\n", employee.EmployeeSalary)
	fmt.Printf("Employee Age: %s\n", employee.EmployeeAge)
	fmt.Printf("Profile Image: %s\n", employee.ProfileImage)
}
