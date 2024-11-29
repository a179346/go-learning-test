package before_srp_test

import (
	"testing"

	"github.com/a179346/go-learning-tests/test/solid/single_responsibility/before_srp"
)

func TestStudentReporter(t *testing.T) {
	t.Run("Normal case", func(t *testing.T) {
		studentInfo := "John,90,85,93"

		studentReporter := before_srp.StudentReporter{}
		err := studentReporter.Read(studentInfo)
		if err != nil {
			t.Error(err)
		}

		report := studentReporter.GenerateReport()
		expectedReport := "Name: John, English: 90, Math: 85, Science: 93"
		if report != expectedReport {
			t.Errorf("Report: %s\nExpected: %s", report, expectedReport)
		}
	})

	t.Run("Student info wrong length", func(t *testing.T) {
		studentInfo := "John,90,85"

		studentReporter := before_srp.StudentReporter{}
		err := studentReporter.Read(studentInfo)
		if err == nil {
			t.Errorf("Should throw error: wrong info length")
		}
	})

	t.Run("Student info wrong ensligh score", func(t *testing.T) {
		studentInfo := "John,90.1,85,93"

		studentReporter := before_srp.StudentReporter{}
		err := studentReporter.Read(studentInfo)
		if err == nil {
			t.Errorf("Should throw error: wrong ensligh score")
		}
	})
}
