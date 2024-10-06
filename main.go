package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
)

// Creating Student Management API

type Student struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Email  string `json:"email"`
	Course string `json:"course"`
}

var students []Student
var nextID = 1

func createStudent(w http.ResponseWriter, r *http.Request) {
	var newStudent Student
	err := json.NewDecoder(r.Body).Decode(&newStudent)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	newStudent.ID = nextID
	nextID++
	students = append(students, newStudent)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(newStudent)
}

func getAllStudents(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(students)
}

func getStudentByID(w http.ResponseWriter, r *http.Request) {
	// Extract the ID from the URL path
	idStr := r.URL.Path[len("/students/"):]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid student ID", http.StatusBadRequest)
		return
	}

	// Search for the student
	for _, student := range students {
		if student.ID == id {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(student)
			return
		}
	}

	http.Error(w, "Student not found", http.StatusNotFound)
}

func updateStudent(w http.ResponseWriter, r *http.Request) {
	// Extract the ID from the URL path
	idStr := r.URL.Path[len("/students/"):]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid student ID", http.StatusBadRequest)
		return
	}

	// Decode the updated student data
	var updatedStudent Student
	err = json.NewDecoder(r.Body).Decode(&updatedStudent)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Find and update the student
	for i, student := range students {
		if student.ID == id {
			updatedStudent.ID = id // Ensure the ID remains the same
			students[i] = updatedStudent

			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(updatedStudent)
			return
		}
	}

	http.Error(w, "Student not found", http.StatusNotFound)
}

func deleteStudent(w http.ResponseWriter, r *http.Request) {
	// Extract the ID from the URL path
	idStr := r.URL.Path[len("/students/"):]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid student ID", http.StatusBadRequest)
		return
	}

	// Find and delete the student
	for i, student := range students {
		if student.ID == id {
			// Remove the student from the slice
			students = append(students[:i], students[i+1:]...)
			w.WriteHeader(http.StatusNoContent) // 204 No Content
			return
		}
	}

	http.Error(w, "Student not found", http.StatusNotFound)
}

func main() {
	http.HandleFunc("/students", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			getAllStudents(w, r)
		case http.MethodPost:
			createStudent(w, r)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	http.HandleFunc("/students/", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			getStudentByID(w, r)
		case http.MethodPut:
			updateStudent(w, r)
		case http.MethodDelete:
			deleteStudent(w, r)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	fmt.Println("Server is running on port 7881...")
	log.Fatal(http.ListenAndServe(":7881", nil))
}
