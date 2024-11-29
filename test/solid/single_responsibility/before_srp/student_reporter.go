package before_srp

import (
	"errors"
	"fmt"
	"strconv"
	"strings"

	singleresponsibility "github.com/a179346/go-learning-tests/test/solid/single_responsibility"
)

type StudentReporter struct {
	student *singleresponsibility.Student
}

func (sr *StudentReporter) Read(studentInfo string) error {
	info := strings.Split(studentInfo, ",")
	if len(info) != 4 {
		return errors.New("The length of student info should equals to 4")
	}

	name := info[0]
	englishScore, err := strconv.Atoi(info[1])
	if err != nil {
		return err
	}
	mathScore, err := strconv.Atoi(info[2])
	if err != nil {
		return err
	}
	scienceScore, err := strconv.Atoi(info[3])
	if err != nil {
		return err
	}

	sr.student = &singleresponsibility.Student{
		Name:         name,
		EnglishScore: englishScore,
		MathScore:    mathScore,
		ScienceScore: scienceScore,
	}

	return nil
}

func (sr StudentReporter) GenerateReport() string {
	return fmt.Sprintf(
		"Name: %s, English: %v, Math: %v, Science: %v",
		sr.student.Name,
		sr.student.EnglishScore,
		sr.student.MathScore,
		sr.student.ScienceScore,
	)
}
