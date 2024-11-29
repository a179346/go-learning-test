package before_isp

import "github.com/a179346/go-learning-tests/test/solid"

/**
 * Interface segregation principle:
 *
 * A client should never be forced to implement an interface that it doesn't use or clients shouldn't be forced to depend on methods they do not use.
 */

/**
 * In this example, the PrintStudentsWhoPassed function only depends on the GetAll method of the StudentRepository interface.
 *
 * However, the StudentRepository interface has an AddStudent method that is not used by the PrintStudentsWhoPassed function.
 *
 * This violates the interface segregation principle.
 */
func PrintStudentsWhoPassed(sr StudentRepository) {
	students := sr.GetAll()
	for _, student := range students {
		if student.EnglishScore >= 60 && student.MathScore >= 60 && student.ScienceScore >= 60 {
			println(student.Name)
		}
	}
}

type StudentRepository interface {
	GetAll() []solid.Student
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
