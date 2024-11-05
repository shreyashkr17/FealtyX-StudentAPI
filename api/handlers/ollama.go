package handlers

import (
	"net/http"
	"student-api/services"
)

func GetStudentSummary(w http.ResponseWriter, r *http.Request, id int) {
	summary, err := services.GenerateStudentSummary(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Write([]byte(summary))
}
