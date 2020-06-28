package employees

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
)

const EMPLOYEES_URL = "http://dummy.restapiexample.com/api/v1/employees"
const EMPLOYEE_URL = "http://dummy.restapiexample.com/api/v1/employee/%s"

var validArgs = []string{"list", "show"}

func Run(args []string) {
	parsedArgs, err := parseArgs(args)
	if err != nil {
		fmt.Printf("[Error] %v\n", err)
	}
	switch parsedArgs.Command {
	case "list":
		if resp, err := listEmployees(); err != nil {
			fmt.Printf("[Error] %v\n", err)
		} else {
			resp.Print()
		}
	case "show":
		if resp, err := showEmployee(parsedArgs.EmployeeID); err != nil {
			fmt.Printf("[Error] %v\n", err)
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

func listEmployees() (EmployeeList, error) {
	resp, err := http.Get(EMPLOYEES_URL)
	if err != nil {
		return EmployeeList{}, err
	}
	responseData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return EmployeeList{}, err
	}
	var apiResponse EmployeesApiResponse
	json.Unmarshal(responseData, &apiResponse)
	if apiResponse.Status != "success" {
		return EmployeeList{}, fmt.Errorf("Invalid API response: %s", apiResponse.Message)
	}
	return EmployeeList{Employees: apiResponse.Data}, nil
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
		return Employee{}, fmt.Errorf("Invalid API response: %s", apiResponse.Message)
	}
	return apiResponse.Data, nil
}

func (employees EmployeeList) Print() {
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
