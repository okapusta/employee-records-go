package main

import "os"

import "github.com/okapusta/employee-records-go/pkg/employees"

func main() {
  employees.Run(os.Args)
}
