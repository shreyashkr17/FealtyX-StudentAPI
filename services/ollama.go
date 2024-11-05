package services

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"sync"
	"time"
)

var ollamaMux sync.Mutex

func GenerateStudentSummary(id int) (string, error) {
	ollamaMux.Lock()
	defer ollamaMux.Unlock()

	student, err := GetStudentByID(id)
	if err != nil {
		return "", err
	}

	prompt := fmt.Sprintf("Summarize the profile of a student with ID %d, name %s, age %d, and email %s.", student.ID, student.Name, student.Age, student.Email)
	reqBody, err := json.Marshal(map[string]interface{}{
		"model":  "llama3.2",
		"prompt": prompt,
		"stream": false,
	})
	if err != nil {
		return "", err
	}

	ollamaAPIURL := os.Getenv("OLLAMA_API_URL")
	if ollamaAPIURL == "" {
		return "", errors.New("OLLAMA_API_URL not set in environment")
	}

	var summaryBuilder bytes.Buffer

	for {
		resp, err := http.Post(ollamaAPIURL, "application/json", bytes.NewBuffer(reqBody))
		if err != nil {
			return "", err
		}
		defer resp.Body.Close()

		var res map[string]interface{}
		if err := json.NewDecoder(resp.Body).Decode(&res); err != nil {
			return "", err
		}

		log.Printf("Response from llama3.2 API: %+v\n", res)

		responsePart, exists := res["response"].(string)
		if exists {
			summaryBuilder.WriteString(responsePart + " ")
		}

		if done, ok := res["done"].(bool); ok && done {
			break
		}

		time.Sleep(500 * time.Millisecond)
	}

	summary := summaryBuilder.String()
	if summary == "" {
		return "", errors.New("summary generation failed: no content received")
	}
	return summary, nil
}
