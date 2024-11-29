package after_srp_test

import (
	"testing"

	"github.com/a179346/go-learning-tests/test/solid"
	"github.com/a179346/go-learning-tests/test/solid/single_responsibility/after_srp"
)

func TestStudentParser(t *testing.T) {
	t.Run("Normal case", func(t *testing.T) {
		studentInfo := "John,90,85,93"

		studentParser := after_srp.StudentParser{}
		student, err := studentParser.Parse(studentInfo)
		if err != nil {
			t.Error(err)
		}
		if student == nil {
			t.Errorf("Should successfully parse the student info")
			return
		}

		expectedStudent := &solid.Student{
			Name:         "John",
			EnglishScore: 90,
			MathScore:    85,
			ScienceScore: 93,
		}

		if *student != *expectedStudent {
			t.Errorf("Student: %+v\nExpected: %+v", student, expectedStudent)
		}
	})

	t.Run("Student info wrong length", func(t *testing.T) {
		studentInfo := "John,90,85"

		studentParser := after_srp.StudentParser{}
		_, err := studentParser.Parse(studentInfo)
		if err == nil {
			t.Errorf("Should throw error: wrong info length")
		}
	})

	t.Run("Student info wrong english score", func(t *testing.T) {
		studentInfo := "John,90.1,85,93"

		studentParser := after_srp.StudentParser{}
		_, err := studentParser.Parse(studentInfo)
		if err == nil {
			t.Errorf("Should throw error: wrong english score")
		}
	})
}

func TestStudentReporter(t *testing.T) {
	t.Run("Should return corresponding report", func(t *testing.T) {
		studentReporter := after_srp.StudentReporter{}

		student := &solid.Student{
			Name:         "John",
			EnglishScore: 90,
			MathScore:    85,
			ScienceScore: 93,
		}

		report := studentReporter.GenerateReport(*student)
		expectedReport := "Name: John, English: 90, Math: 85, Science: 93"
		if report != expectedReport {
			t.Errorf("Report: %s\nExpected: %s", report, expectedReport)
		}
	})
}
