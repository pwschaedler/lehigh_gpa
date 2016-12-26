package calculator

import "fmt"

type Calculator struct {
	grades []Grade
	gpa    float64
}

var legend = map[string]float64{
	"A":  4,
	"A-": 3.7,
}

func NewCalc() *Calculator {
	var calc Calculator
	return &calc
}

func (calc *Calculator) InsertGrade(grade Grade) {
	calc.grades = append(calc.grades, grade)
	calc.calculate()
	fmt.Println("Not yet implemented, but I added one grade!")
}

func (calc *Calculator) calculate() {
	var totalQualPoints float64
	var totalCredits int

	for _, grade := range calc.grades {
		totalCredits += grade.Credits
		totalQualPoints += legend[grade.Grade] * float64(grade.Credits)
	}

	calc.gpa = totalQualPoints / float64(totalCredits)
}

func (calc *Calculator) ImportFromFile() {
	fmt.Println("Not yet implemented.")
}

func (calc *Calculator) ExportToFile() {
	fmt.Println("Not yet implemented.")
}

func (calc *Calculator) PrintGPA() {
	fmt.Println(calc.gpa)
}

func (calc *Calculator) PrintTranscript() {
	fmt.Println("Not yet implemented, but here's the struct!")
	fmt.Println(calc)
}

/*
func (calc Calculator) String() string {
	return fmt.Sprintf("%v", calc.Grades)
}
*/
