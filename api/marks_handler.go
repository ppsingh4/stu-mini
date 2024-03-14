package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"github.com/ppsingh4/stu-mini/dto/transport"
	"github.com/ppsingh4/stu-mini/service"
	log "github.com/sirupsen/logrus"
	"github.com/gorilla/mux"
)

type MarksHandler struct {
	marksService service.MarksService
}

// NewMarksHandler creates a new MarksHandler.
func NewMarksHandler(marksService service.MarksService) MarksHandler {
	return MarksHandler{marksService: marksService}
}

// Implement handler functions for marks endpoints (e.g., Create, Update, Delete, Get).

// CreateMarksHandler creates a new set of marks
func (h *MarksHandler) CreateMarksHandler(w http.ResponseWriter, r *http.Request) {
	// Implement logic to create marks using HTTP
	// Example:
	// params := mux.Vars(r)
	// studentID := params["stu_id"]
	log.Println("creating marks")
	decoder := json.NewDecoder(r.Body)
	var m transport.Mark
	err := decoder.Decode(&m)
	if err != nil {
		http.Error(w, "invalid input", http.StatusBadRequest)
		return
	}
	result := h.marksService.CreateMarks(&m)
	if result != nil {
		http.Error(w, "result.Error.Error()", http.StatusInternalServerError)
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
	studentID := params["id"]
	// var m transport.MarksDTO
	result, err := h.marksService.GetMarks(studentID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	response, err := json.Marshal(result)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
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
	var updatedMarks transport.Mark
	err := decoder.Decode(&updatedMarks)
	if err != nil {
		http.Error(w, "invalid input", http.StatusBadRequest)
		return
	}
	// Update the existing marks
	result := h.marksService.UpdateMarks(studentID, &updatedMarks)
	if result != nil {
		http.Error(w, "result.Error.Error()", http.StatusInternalServerError)
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
	result := h.marksService.DeleteMarks(studentID)
	if result != nil {
		http.Error(w, "result.Error.Error()", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Marks deleted successfully!")
}
