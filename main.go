package main

import (
	"net"
	"time"

	"github.com/ppsingh4/stu-mini/api"
	"github.com/ppsingh4/stu-mini/dao"
	"github.com/ppsingh4/stu-mini/db"
	"github.com/ppsingh4/stu-mini/service"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"gorm.io/gorm"

    pb "github.com/ppsingh4/stu-mini/api/pb"
    
)

func main() {
	// Initialize the database
	time.Sleep(10 * time.Second)
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
	// studentHandler := api.NewStudentHandler(studentService)
	// marksHandler := api.NewMarksHandler(MarksService)

	// // Setup the router
	// router := api.SetupRouter(studentHandler, marksHandler)

	// Start the HTTP server

	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()

    studentServer := api.NewStudentServer(studentService)
    marksServer := api.NewMarksServer(MarksService)

    pb.RegisterStudentServiceServer(s, studentServer)
    pb.RegisterMarksServiceServer(s, marksServer)
    
	// log.Println("Server listening on port 8080")
	// err = http.ListenAndServe(":8080", router)
	// if err != nil {
	//     log.Fatal("Failed to start server:", err)
	// }

	log.Println("gRPC server started")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
func ConvertToDatabase(gormDB *gorm.DB) *db.Database {
	return &db.Database{DB: gormDB}
}
