package main

import (
	"log"
	"net/http"
	"smod/api"
	"smod/dao"
	"smod/db"
	"smod/service"
	"time"

	"gorm.io/gorm"
)

func main() {
    // Initialize the database
	time.Sleep(10*time.Second)
	dsn := "host=postgres user=hp password=Harshith dbname=prithvi port=5432 sslmode=disable TimeZone=Asia/Shanghai"
    ndatabase, err := db.Initialize(dsn)
    if err != nil {
        log.Fatal("Failed to initialize database:", err)
    }
	//defer ndatabase.Close()
	database := ConvertToDatabase(ndatabase)

    // Close the database connection when main function exits
	
	studentDAO := dao.NewStudentDAO(database)
    studentService := service.NewStudentService(studentDAO)

	MarksDAO := dao.NewMarksDAO(database)
    MarksService := service.NewMarksService(MarksDAO)
    // Initialize the handlers
    studentHandler := api.NewStudentHandler(studentService)
    marksHandler := api.NewMarksHandler(MarksService)

    // Setup the router
    router := api.SetupRouter(studentHandler, marksHandler)

    // Start the HTTP server
    log.Println("Server listening on port 8080")
    err = http.ListenAndServe(":8080", router)
    if err != nil {
        log.Fatal("Failed to start server:", err)
    }
}
func ConvertToDatabase(gormDB *gorm.DB) *db.Database {
    return &db.Database{DB: gormDB}
}