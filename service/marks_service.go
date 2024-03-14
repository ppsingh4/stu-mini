package service
import (
    "smod/entity"
    "smod/dto/transport"
)
// MarksService defines the interface for marks-related operations.
type MarksService interface {
    CreateMarks( marksDTO *transport.Mark) error
    GetMarks(studentID string) (*entity.Mark, error)
    UpdateMarks(studentID string, marksDTO *transport.Mark) error
    DeleteMarks(studentID string) error
}