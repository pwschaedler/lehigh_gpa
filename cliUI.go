package main

import (
	"fmt"
	"os"

	"bufio"

	"github.com/pwschaedler/lehighgpa/calculator"
)

type CliUI struct{}

func NewUI() *CliUI {
	return new(CliUI)
}

func (ui CliUI) PrintOptions() {
	fmt.Println("\nPlease select an option.")
	fmt.Println("a - Add Grade (not yet implemented)")
	fmt.Println("e - Export Grades To File (not yet implemented)")
	fmt.Println("g - Print GPA")
	fmt.Println("i - Import Grades From File (not yet implemented)")
	fmt.Println("p - Print Transcript (not yet implemented)")
	fmt.Println("q - Quit")
}

func (ui CliUI) displayUI() {
	var calc = calculator.NewCalc()
	var r = bufio.NewReader(os.Stdin)

	fmt.Println("Welcome to the Lehigh GPA Calculator!")
	ui.PrintOptions()

	for {
		// Print prompt
		fmt.Print("> ")

		// Select an option
		op, _, _ := r.ReadRune()
		bytesLeft := r.Buffered()
		r.Discard(bytesLeft)

		switch op {
		case 'a':
			calc.InsertGrade(calculator.Grade{ClassName: "CSE 17", Credits: 4, Grade: "A"})
		case 'e':
			calc.ExportToFile()
		case 'g':
			calc.PrintGPA()
		case 'h':
			ui.PrintOptions()
		case 'i':
			calc.ImportFromFile()
		case 'p':
			calc.PrintTranscript()
		case 'q':
			return
		}
	}
}
