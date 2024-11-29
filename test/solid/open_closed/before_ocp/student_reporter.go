package before_ocp

import (
	"encoding/json"
	"fmt"

	"github.com/a179346/go-learning-tests/test/solid"
)

/**
 * Open–closed principle:
 *
 * Software entities (classes, modules, functions, etc.) should be open for extension, but closed for modification.
 */

/**
 * In this example, every time we want to add a new report format, we need to modify the GenerateReport method.
 */
type ReportFormat int

const (
	ReportFormatText ReportFormat = iota
	ReportFormatJSON
)

type StudentReporter struct {
	student solid.Student
}

func NewStudentReporter(student solid.Student) *StudentReporter {
	return &StudentReporter{student: student}
}

func (sr StudentReporter) GenerateReport(format ReportFormat) string {
	switch format {
	case ReportFormatText:
		return fmt.Sprintf(
			"Name: %s, English: %v, Math: %v, Science: %v",
			sr.student.Name,
			sr.student.EnglishScore,
			sr.student.MathScore,
			sr.student.ScienceScore,
		)

	case ReportFormatJSON:
		bytes, _ := json.Marshal(map[string]interface{}{
			"name":    sr.student.Name,
			"english": sr.student.EnglishScore,
			"math":    sr.student.MathScore,
			"science": sr.student.ScienceScore,
		})
		return string(bytes)

	default:
		return ""
	}
}
