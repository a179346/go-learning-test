package after_srp

import (
	"fmt"

	"github.com/a179346/go-learning-tests/test/solid"
)

type StudentReporter struct{}

func (sr StudentReporter) GenerateReport(student solid.Student) string {
	return fmt.Sprintf(
		"Name: %s, English: %v, Math: %v, Science: %v",
		student.Name,
		student.EnglishScore,
		student.MathScore,
		student.ScienceScore,
	)
}
