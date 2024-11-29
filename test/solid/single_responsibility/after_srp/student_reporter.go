package after_srp

import (
	"fmt"

	singleresponsibility "github.com/a179346/go-learning-tests/test/solid/single_responsibility"
)

type StudentReporter struct{}

func (sr StudentReporter) GenerateReport(student singleresponsibility.Student) string {
	return fmt.Sprintf(
		"Name: %s, English: %v, Math: %v, Science: %v",
		student.Name,
		student.EnglishScore,
		student.MathScore,
		student.ScienceScore,
	)
}
