package calculator

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"os/user"
)

type Calculator struct {
	Transcript Transcript
	Gpa        float64
}

// Constants
var legend = map[string]float64{
	"A":  4,
	"A-": 3.7,
}
var usr, _ = user.Current()
var defaultFilepath = fmt.Sprintf("%s/Documents/transcript.json", usr.HomeDir)

func NewCalc() *Calculator {
	var calc Calculator
	return &calc
}

func (calc *Calculator) InsertGrade(class Class) {
	calc.Transcript.Classes = append(calc.Transcript.Classes, class)
	calc.calculate()
	fmt.Println("Not yet implemented, but I added one grade!")
}

func (calc *Calculator) calculate() {
	var totalQualPoints float64
	var totalCredits int

	for _, class := range calc.Transcript.Classes {
		totalCredits += class.Credits
		totalQualPoints += legend[class.Grade] * float64(class.Credits)
	}

	calc.Gpa = totalQualPoints / float64(totalCredits)
}

func (calc *Calculator) ImportFromFile() {
	var f, err = os.Open(defaultFilepath)
	if err != nil {
		fmt.Println("Could not open file.")
		fmt.Println(err)
		return
	}

	var r = bufio.NewReader(f)
	var decoder = json.NewDecoder(r)

	err = decoder.Decode(&calc.Transcript)
	if err != nil {
		fmt.Println("Reading error on import.")
		fmt.Println(err)
	}

	calc.calculate()
}

func (calc *Calculator) ExportToFile() {
	var f, err = os.Create(defaultFilepath)
	if err != nil {
		fmt.Println("Could not write to file.")
		fmt.Println(err)
		return
	}

	var w = bufio.NewWriter(f)
	var encoder = json.NewEncoder(w)

	err = encoder.Encode(calc.Transcript)
	if err != nil {
		fmt.Println("Writing error on export.")
		fmt.Println(err)
	}

	err = w.Flush()
	if err != nil {
		fmt.Println("Writing error on export.")
		fmt.Println(err)
	}
}

func (calc *Calculator) PrintGPA() {
	fmt.Println(calc.Gpa)
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
