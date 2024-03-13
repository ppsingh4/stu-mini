package service
import (
    "smod/entity"
    "smod/dto/transport"
)
// MarksService defines the interface for marks-related operations.
type MarksService interface {
    CreateMarks(studentID string, marksDTO *transport.Marks) error
    GetMarks(studentID string) (*entity.Marks, error)
    UpdateMarks(studentID string, marksDTO *transport.Marks) error
    DeleteMarks(studentID string) error
}