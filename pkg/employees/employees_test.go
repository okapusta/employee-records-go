package employees

import (
  "fmt"
  "testing"
  "github.com/stretchr/testify/assert"
)

func Test_parseArgs_ValidArgs(t *testing.T) {
  var result, expected Args
  var err error

  listValidArgs := []string{`/file/path`, `list`}
  result, err = parseArgs(listValidArgs)
  assert.Equal(t, err, nil, "Error should be nil")
  expected = Args{Command: "list"}
  assert.Equal(t, result, expected, "Parses args into struct")

  validShowArgs := []string{`/fie/path`, `show`, `1`}
  result, err = parseArgs(validShowArgs)
  assert.Equal(t, err, nil, "Error should be nil")
  expected = Args{Command: "show", EmployeeID: 1}
  assert.Equal(t, result, expected, "Parses args into struct")
}

func Test_parseArgs_InvalidCommand(t *testing.T) {
  invaidCommand := []string{`/file/path`, `invalid`}
  _, err := parseArgs(invaidCommand)
  assert.Equal(t, err, fmt.Errorf("Invalid Command: invalid"))
}

func Test_parseArgs_ShowNoID(t *testing.T) {
  invaidCommand := []string{`/file/path`, `show`}
  _, err := parseArgs(invaidCommand)
  assert.Equal(t, err, fmt.Errorf("You must provide employee ID"))
}

func Test_listEmplyees_ValidResponse(t *testing.T) {

}

func Test_listEmplyees_InvalidResponse(t *testing.T) {

}

func Test_showEmployee_ValidResponse(t *testing.T) {

}

func Test_showEmployee_InValidResponse(t *testing.T) {
}
