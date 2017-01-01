package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/pwschaedler/lehighgpa/calculator"
)

type CliUI struct{}

func NewUI() *CliUI {
	return new(CliUI)
}

func (ui CliUI) printOptions() {
	fmt.Println("\nPlease select an option.")
	fmt.Println("a - Add Grade (not yet implemented)     h - Print Prompt Options")
	fmt.Println("e - Export Grades To File               i - Import Grades From File")
	fmt.Println("g - Print GPA                           p - Print Transcript (not yet implemented)")
	fmt.Println("q - Quit")
}

func (ui CliUI) DisplayUI() {
	var calc = calculator.NewCalc()
	var r = bufio.NewReader(os.Stdin)

	fmt.Println("Welcome to the Lehigh GPA Calculator!")
	ui.printOptions()

	for {
		// Print prompt
		fmt.Print("> ")

		// Select an option
		op, _, _ := r.ReadRune()
		bytesLeft := r.Buffered()
		r.Discard(bytesLeft)

		switch op {
		case 'a':
			calc.InsertGrade(calculator.Class{ClassName: "CSE 17", Credits: 4, Grade: "A"})
		case 'e':
			calc.ExportToFile()
		case 'g':
			calc.PrintGPA()
		case 'h':
			ui.printOptions()
		case 'i':
			calc.ImportFromFile()
		case 'p':
			calc.PrintTranscript()
		case 'q':
			return
		}
	}
}
