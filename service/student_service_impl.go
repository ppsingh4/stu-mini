package service
import (
    "smod/dao"
    "smod/entity"
    "smod/dto/transport"
)
// StudentServiceImpl is the implementation of StudentService.
type StudentServiceImpl struct {
    studentDAO dao.StudentDAO
}
// NewStudentService creates a new instance of StudentService.
func NewStudentService(studentDAO dao.StudentDAO) StudentService {
    return &StudentServiceImpl{studentDAO: studentDAO}
}
// CreateStudent creates a new student.
func (s *StudentServiceImpl) CreateStudent(studentDTO *transport.Student) (string, error) {
    // Validation and business logic can be performed here
    // Convert DTO to entity
    student := entity.Student{
        FirstName: studentDTO.FirstName,
        LastName:  studentDTO.LastName,
        Grade:     studentDTO.Grade,
        Phno:   studentDTO.Phno,
        EmailID:   studentDTO.EmailID,
        Address:   studentDTO.Address,
    }
    // Call DAO to create student
    studentID, err := s.studentDAO.Create(&student)
    if err != nil {
        return "", err
    }
    return studentID, nil
}
// GetStudent retrieves a student by ID.
func (s *StudentServiceImpl) GetStudent(studentID string) (*entity.Student, error) {
    // Call DAO to get student by ID
    student, err := s.studentDAO.GetByID(studentID)
    if err != nil {
        return nil, err
    }
    return student, nil
}
// UpdateStudent updates an existing student.
func (s *StudentServiceImpl) UpdateStudent(studentID string, studentDTO *transport.Student) error {
    // Validation and business logic can be performed here
    // Convert DTO to entity
    updatedStudent := entity.Student{
        FirstName: studentDTO.FirstName,
        LastName:  studentDTO.LastName,
        Grade:     studentDTO.Grade,
        Phno:   studentDTO.Phno,
        EmailID:   studentDTO.EmailID,
        Address:   studentDTO.Address,
    }
    // Call DAO to update student
    err := s.studentDAO.Update(studentID, &updatedStudent)
    if err != nil {
        return err
    }
    return nil
}
// DeleteStudent deletes an existing student by ID.
func (s *StudentServiceImpl) DeleteStudent(studentID string) error {
    // Call DAO to delete student by ID
    err := s.studentDAO.Delete(studentID)
    if err != nil {
        return err
    }
    return nil
}