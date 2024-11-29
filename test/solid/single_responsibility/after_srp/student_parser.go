package after_srp

import (
	"errors"
	"strconv"
	"strings"

	singleresponsibility "github.com/a179346/go-learning-tests/test/solid/single_responsibility"
)

type StudentParser struct{}

func (sp StudentParser) Parse(studentInfo string) (*singleresponsibility.Student, error) {

	info := strings.Split(studentInfo, ",")
	if len(info) != 4 {
		return nil, errors.New("The length of student info should equals to 4")
	}

	name := info[0]
	englishScore, err := strconv.Atoi(info[1])
	if err != nil {
		return nil, err
	}
	mathScore, err := strconv.Atoi(info[2])
	if err != nil {
		return nil, err
	}
	scienceScore, err := strconv.Atoi(info[3])
	if err != nil {
		return nil, err
	}

	return &singleresponsibility.Student{
		Name:         name,
		EnglishScore: englishScore,
		MathScore:    mathScore,
		ScienceScore: scienceScore,
	}, nil
}
