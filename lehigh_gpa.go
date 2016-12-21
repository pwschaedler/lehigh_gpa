package main

import (
	"fmt"

	"github.com/pwschaedler/lehigh_gpa/calculator"
)

func main() {
	var calc calculator.Calculator
	calc.InsertGrade(calculator.Grade{ClassName: "CSE 17", Credits: 4, Grade: "A"})
	fmt.Println(calc)
}
