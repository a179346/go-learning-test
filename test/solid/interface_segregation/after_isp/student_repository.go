package after_isp

import "github.com/a179346/go-learning-tests/test/solid"

/**
 * We change the PrintStudentsWhoPassed function to depend on the ReadableStudentRepository interface instead of the StudentRepository interface.
 *
 * This adheres to the interface segregation principle.
 */

func PrintStudentsWhoPassed(sr ReadableStudentRepository) {
	students := sr.GetAll()
	for _, student := range students {
		if student.EnglishScore >= 60 && student.MathScore >= 60 && student.ScienceScore >= 60 {
			println(student.Name)
		}
	}
}

type ReadableStudentRepository interface {
	GetAll() []solid.Student
}

type WriteableStudentRepository interface {
	AddStudent(student solid.Student)
}

type StudentRepositoryImpl struct {
	students []solid.Student
}

func NewStudentRepository() *StudentRepositoryImpl {
	return &StudentRepositoryImpl{
		students: []solid.Student{
			{Name: "Alice", EnglishScore: 100, MathScore: 90, ScienceScore: 80},
			{Name: "Bob", EnglishScore: 90, MathScore: 80, ScienceScore: 70},
		},
	}
}

func (sr StudentRepositoryImpl) GetAll() []solid.Student {
	return sr.students
}

func (sr *StudentRepositoryImpl) AddStudent(student solid.Student) {
	sr.students = append(sr.students, student)
}
