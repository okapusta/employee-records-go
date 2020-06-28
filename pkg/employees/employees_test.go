package employees

import (
	"fmt"
	"testing"

	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
)

func Test_parseArgs_ValidArgs(t *testing.T) {
	var result, expected Args
	var err error

	listValidArgs := []string{`/file/path`, `list`}
	result, err = parseArgs(listValidArgs)
	assert.Equal(t, nil, err, "Error should be nil")
	expected = Args{Command: "list"}
	assert.Equal(t, expected, result, "Parses args into struct")

	validShowArgs := []string{`/fie/path`, `show`, `1`}
	result, err = parseArgs(validShowArgs)
	assert.Equal(t, nil, err, "Error should be nil")
	expected = Args{Command: "show", EmployeeID: 1}
	assert.Equal(t, expected, result, "Parses args into struct")
}

func Test_parseArgs_InvalidCommand(t *testing.T) {
	invaidCommand := []string{`/file/path`, `invalid`}
	_, err := parseArgs(invaidCommand)
	assert.Equal(t, fmt.Errorf("Invalid Command: invalid"), err)
}

func Test_parseArgs_ShowNoID(t *testing.T) {
	invaidCommand := []string{`/file/path`, `show`}
	_, err := parseArgs(invaidCommand)
	assert.Equal(t, fmt.Errorf("You must provide employee ID"), err)
}

func Test_listEmplyees_ValidResponse(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("GET", "http://dummy.restapiexample.com/api/v1/employees",
		httpmock.NewStringResponder(200, `{
      "status": "success",
      "data": [{
        "id": "1",
        "employee_name":"Joe Doe",
        "employee_salary":"100",
        "employee_age":"30",
        "profile_image":""}]}`))
	result, err := listEmployees()
	assert.Equal(t, nil, err, "Error should be nil")
	expected := EmployeeList{
		Employees: []Employee{Employee{
			ID:             "1",
			EmployeeName:   "Joe Doe",
			EmployeeSalary: "100",
			EmployeeAge:    "30",
			ProfileImage:   "",
		}},
	}
	assert.Equal(t, expected, result, "Employee list struct")
}

func Test_listEmplyees_InvalidResponse(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("GET", "http://dummy.restapiexample.com/api/v1/employees",
		httpmock.NewStringResponder(200, `{
      "status": "failed",
      "message": "oops"}`))
	_, err := listEmployees()
	assert.Equal(t, fmt.Errorf("Invalid API response: oops"), err, "Error should not be nil")
}

func Test_showEmployee_ValidResponse(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("GET", "http://dummy.restapiexample.com/api/v1/employee/1",
		httpmock.NewStringResponder(200, `{
      "status": "success",
      "data": {
        "id": "1",
        "employee_name":"Joe Doe",
        "employee_salary":"100",
        "employee_age":"30",
        "profile_image":""}}`))
	result, err := showEmployee(int64(1))
	assert.Equal(t, nil, err, "Error should be nil")
	expected := Employee{
		ID:             "1",
		EmployeeName:   "Joe Doe",
		EmployeeSalary: "100",
		EmployeeAge:    "30",
		ProfileImage:   "",
	}
	assert.Equal(t, expected, result, "Employee struct")
}

func Test_showEmployee_InValidResponse(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("GET", "http://dummy.restapiexample.com/api/v1/employee/1",
		httpmock.NewStringResponder(200, `{
      "status": "failed",
      "message": "Oops! someting issue found to fetch record."}`))
	_, err := showEmployee(int64(1))
	assert.Equal(t, fmt.Errorf("Invalid API response: Oops! someting issue found to fetch record."), err, "Error should not be nil")
}
