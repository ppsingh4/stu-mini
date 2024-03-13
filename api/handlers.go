package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"smod/dto/transport"
	"smod/service"

	"github.com/gorilla/mux"
)

// StudentHandler handles HTTP requests related to students.
type StudentHandler struct {
	studentService *service.StudentService
}

// NewStudentHandler creates a new StudentHandler.
func NewStudentHandler(studentService *service.StudentService) *StudentHandler {
	return &StudentHandler{studentService: studentService}
}

// Implement handler functions for student endpoints (e.g., Create, Update, Delete, Get).
func (h *StudentHandler) CreateStudentHandler(w http.ResponseWriter, r *http.Request) {
	// Implement logic to create marks using HTTP
	// Example:
	fmt.Println("creating marks")
	decoder := json.NewDecoder(r.Body)
	var s transport.Student
	err := decoder.Decode(&s)
	if err != nil {
		http.Error(w, "invalid input", http.StatusBadRequest)
		return
	}

	result := h.studentService.CreateStudent(&s)
	if result.Error != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, "Marks created successfully!")
}

// GetMarksByIDHandler retrieves marks by ID
func (h *StudentHandler) GetStudentByIDHandler(w http.ResponseWriter, r *http.Request) {
	// Implement logic to retrieve marks by ID using HTTP
	// Example:
	vars := mux.Vars(r)
	studentID := vars["id"]
	//var s transport.StudentDTO
	result := h.studentService.GetStudent(studentID)
	if result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusInternalServerError)
		return
	}
	response, err := json.Marshal(result)
	if err != nil {
		http.Error(w, "Failed to marshal response", http.StatusInternalServerError)
		return
	}
	// Return marks details as JSON response
	w.Header().Set("the response is", " ")
	w.Write(response)
}

// UpdateMarksHandler updates marks by ID
func (h *StudentHandler)UpdateStudentHandler(w http.ResponseWriter, r *http.Request) {
	// Implement logic to update marks by ID using HTTP
	// Example:
	vars := mux.Vars(r)
	studentID := vars["id"]
	// Retrieve the existing marks
	var existingStudent transport.Student
	result := db.First(&existingStudent, studentID)
	if result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusInternalServerError)
		return
	}
	// Decode the request body to get updated marks details
	decoder := json.NewDecoder(r.Body)
	var updatedStudent transport.Student
	err := decoder.Decode(&updatedStudent)
	if err != nil {
		http.Error(w, "invalid input", http.StatusBadRequest)
		return
	}
	// Update the existing marks
	result = h.studentService.UpdateStudent(studentID, &updatedStudent)
	if result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Student updated successfully!")
}

// DeleteMarksHandler deletes marks by ID
func (h *StudentHandler) DeleteStudentHandler(w http.ResponseWriter, r *http.Request) {
	// Implement logic to delete marks by ID using HTTP
	// Example:
	vars := mux.Vars(r)
	studentID := vars["id"]
	// Delete the marks by ID
	result := h.studentService.DeleteStudent(studentID)
	if result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Marks deleted successfully!")
}

// MarksHandler handles HTTP requests related to marks.
type MarksHandler struct {
	marksService *service.MarksService
}

// NewMarksHandler creates a new MarksHandler.
func NewMarksHandler(marksService *service.MarksService) *MarksHandler {
	return &MarksHandler{marksService: marksService}
}

// Implement handler functions for marks endpoints (e.g., Create, Update, Delete, Get).

// CreateMarksHandler creates a new set of marks
func (h *MarksHandler) CreateMarksHandler(w http.ResponseWriter, r *http.Request) {
	// Implement logic to create marks using HTTP
	// Example:
	params := mux.Vars(r)
	studentID := params["stu_id"]
	fmt.Println("creating marks")
	decoder := json.NewDecoder(r.Body)
	var m transport.MarksDTO
	err := decoder.Decode(&m)
	if err != nil {
		http.Error(w, "invalid input", http.StatusBadRequest)
		return
	}
	result := h.marksService.CreateMarks(studentID, &m)
	if result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, "Marks created successfully!")
}

// GetMarksByIDHandler retrieves marks by ID
func (h *MarksHandler) GetMarksByIDHandler(w http.ResponseWriter, r *http.Request) {
	// Implement logic to retrieve marks by ID using HTTP
	// Example:
	params := mux.Vars(r)
	studentID := params["stu_id"]
	// var m transport.MarksDTO
	result := h.marksService.GetMarks(studentID)
	if result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusInternalServerError)
		return
	}
	response, err := json.Marshal(result)
	if err != nil {
		http.Error(w, result.Error.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("the response is", " ")
	w.Write(response)
}

// UpdateMarksHandler updates marks by ID
func (h *MarksHandler) UpdateMarksHandler(w http.ResponseWriter, r *http.Request) {
	// Implement logic to update marks by ID using HTTP
	// Example:
	vars := mux.Vars(r)
	studentID := vars["id"]
	// Retrieve the existing marks
	// var existingMarks transport.MarksDTO
	// result := db.First(&existingMarks, studentID)
	// if result.Error != nil {
	// 	http.Error(w, result.Error.Error(), http.StatusInternalServerError)
	// 	return
	// }
	// Decode the request body to get updated marks details
	decoder := json.NewDecoder(r.Body)
	var updatedMarks transport.MarksDTO
	err := decoder.Decode(&updatedMarks)
	if err != nil {
		http.Error(w, "invalid input", http.StatusBadRequest)
		return
	}
	// Update the existing marks
	result := h.marksService.UpdateMarks(studentID, &updatedMarks)
	if result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Marks updated successfully!")
}

// DeleteMarksHandler deletes marks by ID
func (h *MarksHandler) DeleteMarksHandler(w http.ResponseWriter, r *http.Request) {
	// Implement logic to delete marks by ID using HTTP
	// Example:
	vars := mux.Vars(r)
	studentID := vars["id"]
	// Delete the marks by ID
	result := h.marksService.DeleteMark(studentID)
	if result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Marks deleted successfully!")
}
