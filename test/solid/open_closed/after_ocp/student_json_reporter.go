package after_ocp

import (
	"encoding/json"

	"github.com/a179346/go-learning-tests/test/solid"
)

type StudentJsonReporter struct {
	student solid.Student
}

func NewStudentJsonReporter(student solid.Student) *StudentJsonReporter {
	return &StudentJsonReporter{student: student}
}

func (sr StudentJsonReporter) GenerateReport(format ReportFormat) string {
	bytes, _ := json.Marshal(map[string]interface{}{
		"name":    sr.student.Name,
		"english": sr.student.EnglishScore,
		"math":    sr.student.MathScore,
		"science": sr.student.ScienceScore,
	})
	return string(bytes)
}
