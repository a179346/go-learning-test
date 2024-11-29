package after_ocp

import "github.com/a179346/go-learning-tests/test/solid"

type ReportFormat int

const (
	ReportFormatText ReportFormat = iota
	ReportFormatJSON
)

type StudentReporter interface {
	GenerateReport(format ReportFormat) string
}

func StudentReporterFactory(student solid.Student, format ReportFormat) StudentReporter {
	switch format {
	case ReportFormatText:
		return NewStudentTextReporter(student)
	case ReportFormatJSON:
		return NewStudentJsonReporter(student)
	default:
		return nil
	}
}
