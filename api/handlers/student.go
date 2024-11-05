package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
	"student-api/api/models"
	"student-api/services"
	"sync"
)

var mu sync.Mutex

func StudentHandler(w http.ResponseWriter, r *http.Request) {

	idStr := strings.TrimPrefix(r.URL.Path, "/students/")

	isSummary := strings.HasSuffix(idStr, "/summary")

	if isSummary {
		idStr = strings.TrimSuffix(idStr, "/summary")
	}

	if idStr == "" || r.URL.Path == "/students" {
		switch r.Method {
		case http.MethodGet:
			GetAllStudents(w, r)
		case http.MethodPost:
			CreateStudent(w, r)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
		return
	}

	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	switch r.Method {
	case http.MethodGet:
		if isSummary {
			GetStudentSummary(w, r, id)
		} else {
			GetStudentByID(w, r, id)
		}
	case http.MethodPut:
		UpdateStudent(w, r, id)
	case http.MethodDelete:
		DeleteStudent(w, r, id)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func GetAllStudents(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	defer mu.Unlock()
	students := services.GetAllStudents()
	json.NewEncoder(w).Encode(students)
}

func CreateStudent(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	defer mu.Unlock()
	var student models.Student
	if err := json.NewDecoder(r.Body).Decode(&student); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}
	createdStudent := services.CreateStudent(student)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(createdStudent)
}

func GetStudentByID(w http.ResponseWriter, r *http.Request, id int) {
	mu.Lock()
	defer mu.Unlock()
	student, err := services.GetStudentByID(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(student)
}

func UpdateStudent(w http.ResponseWriter, r *http.Request, id int) {
	mu.Lock()
	defer mu.Unlock()
	var student models.Student
	if err := json.NewDecoder(r.Body).Decode(&student); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}
	updatedStudent, err := services.UpdateStudent(id, student)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(updatedStudent)
}

func DeleteStudent(w http.ResponseWriter, r *http.Request, id int) {
	mu.Lock()
	defer mu.Unlock()
	if err := services.DeleteStudent(id); err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
