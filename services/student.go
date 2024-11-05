package services

import (
	"errors"
	"student-api/api/models"
	"sync"
)

var (
	students   = make(map[int]models.Student)
	lastID     = 0
	studentMux sync.Mutex
)

func GetAllStudents() []models.Student {
	studentMux.Lock()
	defer studentMux.Unlock()

	result := make([]models.Student, 0, len(students))
	for _, student := range students {
		result = append(result, student)
	}
	return result
}

func CreateStudent(student models.Student) models.Student {
	studentMux.Lock()
	defer studentMux.Unlock()

	lastID++
	student.ID = lastID
	students[student.ID] = student
	return student
}

func GetStudentByID(id int) (models.Student, error) {
	studentMux.Lock()
	defer studentMux.Unlock()

	student, exists := students[id]
	if !exists {
		return models.Student{}, errors.New("student not found")
	}
	return student, nil
}

func UpdateStudent(id int, updatedStudent models.Student) (models.Student, error) {
	studentMux.Lock()
	defer studentMux.Unlock()

	student, exists := students[id]
	if !exists {
		return models.Student{}, errors.New("student not found")
	}
	updatedStudent.ID = student.ID
	students[id] = updatedStudent
	return updatedStudent, nil
}

func DeleteStudent(id int) error {
	studentMux.Lock()
	defer studentMux.Unlock()

	if _, exists := students[id]; !exists {
		return errors.New("student not found")
	}
	delete(students, id)
	return nil
}
