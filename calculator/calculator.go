package calculator

type Calculator struct {
	grades []Grade
	gpa    float64
}

var legend = map[string]float64{
	"A":  4,
	"A-": 3.7,
}

func (calc *Calculator) InsertGrade(grade Grade) {
	calc.grades = append(calc.grades, grade)
	calc.calculate()
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

/*
func (calc Calculator) String() string {
	return fmt.Sprintf("%v", calc.Grades)
}
*/
