package after_ocp

import (
	"fmt"

	"github.com/a179346/go-learning-tests/test/solid"
)

type StudentTextReporter struct {
	student solid.Student
}

func NewStudentTextReporter(student solid.Student) *StudentTextReporter {
	return &StudentTextReporter{student: student}
}

func (sr StudentTextReporter) GenerateReport(format ReportFormat) string {
	return fmt.Sprintf(
		"Name: %s, English: %v, Math: %v, Science: %v",
		sr.student.Name,
		sr.student.EnglishScore,
		sr.student.MathScore,
		sr.student.ScienceScore,
	)
}
