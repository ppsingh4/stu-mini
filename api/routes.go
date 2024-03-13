package api

import (
    "net/http"
    "github.com/gorilla/mux"
)

func SetupRouter(studentHandler StudentHandler, marksHandler MarksHandler) http.Handler {
    router := mux.NewRouter()

    // Student endpoints
    router.HandleFunc("/students", studentHandler.CreateStudentHandler).Methods("POST")
    router.HandleFunc("/students/{id}", studentHandler.GetStudentByIDHandler).Methods("GET")
    router.HandleFunc("/students/{id}", studentHandler.UpdateStudentHandler).Methods("PUT")
    router.HandleFunc("/students/{id}", studentHandler.DeleteStudentHandler).Methods("DELETE")

    // Marks endpoints
    router.HandleFunc("/marks/{stu_id}", marksHandler.CreateMarksHandler).Methods("POST")
    router.HandleFunc("/marks/{stu_id}", marksHandler.GetMarksByIDHandler).Methods("GET")
    router.HandleFunc("/marks/{id}", marksHandler.UpdateMarksHandler).Methods("PUT")
    router.HandleFunc("/marks/{id}", marksHandler.DeleteMarksHandler).Methods("DELETE")

    return router
}
