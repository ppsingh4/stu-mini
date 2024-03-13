package service
import (
    "smod/dao"
    "smod/entity"
    "smod/dto/transport"
)
// MarksServiceImpl is the implementation of MarksService.
type MarksServiceImpl struct {
    marksDAO dao.MarksDAO
}
// NewMarksService creates a new instance of MarksService.
func NewMarksService(marksDAO dao.MarksDAO) MarksService {
    return &MarksServiceImpl{marksDAO: marksDAO}
}
// CreateMarks creates marks for a student.
func (s *MarksServiceImpl) CreateMarks(studentID string, marksDTO *transport.Marks) error {
    // Convert DTO to entity
    marks := entity.Marks{
        ID: marksDTO.ID,
        Maths:  marksDTO.Maths,
        Social:  marksDTO.Social,
        Science:  marksDTO.Science,
		English: marksDTO.English,
		Hindi: marksDTO.Hindi,
    }
    // Call DAO to create marks
    err := s.marksDAO.Create(&marks)
    if err != nil {
        return err
    }
    return nil
}
// GetMarks retrieves marks for a student.
func (s *MarksServiceImpl) GetMarks(studentID string) (*entity.Marks, error) {
    // Call DAO to get marks for a student
    marks, err := s.marksDAO.GetByStudentID(studentID)
    if err != nil {
        return nil, err
    }
    return marks, nil
}
// UpdateMarks updates marks for a student.
func (s *MarksServiceImpl) UpdateMarks(studentID string, marksDTO *transport.Marks) error {
    // Convert DTO to entity
    updatedMarks := entity.Marks{
        ID: marksDTO.ID,
        Maths:  marksDTO.Maths,
        Social:  marksDTO.Social,
        Science:  marksDTO.Science,
		English: marksDTO.English,
		Hindi: marksDTO.Hindi,
    }
    // Call DAO to update marks
    err := s.marksDAO.UpdateByStudentID(studentID, &updatedMarks)
    if err != nil {
        return err
    }
    return nil
}
// DeleteMarks deletes marks for a student.
func (s *MarksServiceImpl) DeleteMarks(studentID string) error {
    // Call DAO to delete marks for a student
    err := s.marksDAO.DeleteByStudentID(studentID)
    if err != nil {
        return err
    }
    return nil
}