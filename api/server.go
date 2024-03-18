package api

import (
	"context"
	"fmt"
	"log"

	pb "github.com/ppsingh4/stu-mini/api/pb"
	"github.com/ppsingh4/stu-mini/dto/transport"
	"github.com/ppsingh4/stu-mini/service" // Import your service package
)

// StudentServer implements the StudentService gRPC service
type StudentServer struct {
	pb.UnimplementedStudentServiceServer
	studentService service.StudentService
}

func NewStudentServer(studentService service.StudentService) StudentServer {
	return StudentServer{studentService: studentService}
}

// CreateStudent implements the CreateStudent gRPC method
func (s StudentServer) CreateStudentRPC(ctx context.Context, req *pb.CreateStudentRequest) (*pb.CreateStudentResponse, error) {

	createdStudent := req.GetStudent()
	log.Printf("Creating student: %v", createdStudent)
	dtoStudent := ConvertPBStudentToDTOStudent(createdStudent)
	id, err := s.studentService.CreateStudent(dtoStudent)
	if err != nil {
		return nil, err
	}
	fmt.Println(id)
	return &pb.CreateStudentResponse{Message: "Student created successfully"}, nil
}

func ConvertPBStudentToDTOStudent(pbStudent *pb.StudentRequest) transport.Student {
	return transport.Student{
		// Fill in the fields with corresponding values from pbStudent
		FirstName: pbStudent.FirstName,
		LastName:  pbStudent.LastName,
		Grade:     pbStudent.Grade,
		Phno:      pbStudent.Phno,
		EmailID:   pbStudent.EmailId,
		Address:   pbStudent.Address,
	}
}

// GetStudent implements the GetStudent gRPC method
func (s StudentServer) GetStudentRPC(ctx context.Context, req *pb.GetStudentRequest) (*pb.GetStudentResponse, error) {
	// Implement logic to retrieve a student by ID
	studentID := req.GetId()
	// Retrieve student from service layer
	student, err := s.studentService.GetStudent(studentID)
	if err != nil {
		return nil, err
	}

	pbStudent := &pb.StudentResponse{
		Id:        student.ID,
		FirstName: student.FirstName,
		LastName:  student.LastName,
		Grade:     student.Grade,
		Phno:      student.Phno,
		EmailId:   student.EmailID,
		Address:   student.Address,
	}
	// Return the student fetched
	return &pb.GetStudentResponse{Student: pbStudent}, nil
}

// UpdateStudent implements the UpdateStudent gRPC method
func (s StudentServer) UpdateStudentRPC(ctx context.Context, req *pb.UpdateStudentRequest) (*pb.UpdateStudentResponse, error) {
	// Implement logic to update a student
	studentId := req.GetId()
	newStudent := req.GetStudent()
	dtoNewStudent := ConvertPBStudentToDTOStudent(newStudent)
	if err := s.studentService.UpdateStudent(studentId, dtoNewStudent); err != nil {
		return nil, err
	}
	// For demonstration, let's just echo back the updated student received
	return &pb.UpdateStudentResponse{Message: "Updated student successfully"}, nil
}

// DeleteStudent implements the DeleteStudent gRPC method
func (s StudentServer) DeleteStudentRPC(ctx context.Context, req *pb.DeleteStudentRequest) (*pb.DeleteStudentResponse, error) {
	// Implement logic to delete a student by ID
	studentId := req.GetId()
	if err := s.studentService.DeleteStudent(studentId); err != nil {
		return nil, err
	}
	// For demonstration, let's just return a success message
	return &pb.DeleteStudentResponse{Message: "Student deleted successfully"}, nil
}

// MarksServer implements the MarksService gRPC service
type MarksServer struct {
	pb.UnimplementedMarksServiceServer
	marksService service.MarksService
}

func NewMarksServer(marksService service.MarksService) MarksServer {
	return MarksServer{marksService: marksService}
}

// CreateMarks implements the CreateMarks gRPC method
func (s MarksServer) CreateMarksRPC(ctx context.Context, req *pb.CreateMarksRequest) (*pb.CreateMarksResponse, error) {
	createdMarks := req.GetMarks()
	dtoMarks := ConvertPBMarksToDTOMarks(createdMarks)
	if err := s.marksService.CreateMarks(dtoMarks); err != nil {
		return nil, err
	}
	return &pb.CreateMarksResponse{Result: "Created marks successfully"}, nil
}

// GetMarks implements the GetMarks gRPC method
func (s MarksServer) GetMarksRPC(ctx context.Context, req *pb.GetMarksRequest) (*pb.GetMarksResponse, error) {
	// Implement logic to retrieve marks by student ID
	studentID := req.GetStudentId()
	marks, err := s.marksService.GetMarks(studentID)
	if err != nil {
		return nil, err
	}

	pbMarks := &pb.Marks{
		// ID:         marks.StudentId,
		Sub_1: marks.Sub_1,
		Sub_2: marks.Sub_2,
		Sub_3: marks.Sub_3,
	}
	return &pb.GetMarksResponse{Marks: pbMarks}, nil
}

// UpdateMarks implements the UpdateMarks gRPC method
func (s MarksServer) UpdateMarksRPC(ctx context.Context, req *pb.UpdateMarksRequest) (*pb.UpdateMarksResponse, error) {
	studentID := req.GetStudentId()
	updatedMarks := req.GetMarks()
	dtoUpdatedMarks := ConvertPBMarksToDTOMarks(updatedMarks)
	if err := s.marksService.UpdateMarks(studentID, dtoUpdatedMarks); err != nil {
		return nil, err
	}
	return &pb.UpdateMarksResponse{Message: "Updated marks successfully "}, nil
}

// DeleteMarks implements the DeleteMarks gRPC method
func (s MarksServer) DeleteMarksRPC(ctx context.Context, req *pb.DeleteMarksRequest) (*pb.DeleteMarksResponse, error) {
	// Implement logic to delete marks by student ID
	studentID := req.GetStudentId()
	if err := s.marksService.DeleteMarks(studentID); err != nil {
		return nil, err
	}
	// For demonstration, let's just return a success message
	return &pb.DeleteMarksResponse{Message: "Marks deleted successfully"}, nil
}

func ConvertPBMarksToDTOMarks(pbMarks *pb.Marks) transport.Mark {
	return transport.Mark{
		// Fill in the fields with corresponding values from pbStudent
		ID:    pbMarks.StudentId,
		Sub_1: pbMarks.Sub_1,
		Sub_2: pbMarks.Sub_2,
		Sub_3: pbMarks.Sub_3,
	}
}

// func main() {
//     // Start gRPC server
//     lis, err := net.Listen("tcp", ":50051")
//     if err != nil {
//         log.Fatalf("failed to listen: %v", err)
//     }
//     s := grpc.NewServer()

//     // Initialize services
//     studentService := service.NewStudentService() // Initialize your student service
//     marksService := service.NewMarksService()     // Initialize your marks service

//     // Register StudentService server
//     pb.RegisterStudentServiceServer(s, &StudentServer{studentService: studentService})

//     // Register MarksService server
//     pb.RegisterMarksServiceServer(s, &MarksServer{})

//     log.Println("gRPC server started")
//     if err := s.Serve(lis); err != nil {
//         log.Fatalf("failed to serve: %v", err)
//     }
// }
