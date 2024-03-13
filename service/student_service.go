package service

import (
	"smod/dto/transport"
	"smod/entity"
)

type StudentService interface {
	CreateStudent(studentDTO *transport.Student) (uint, error)
	GetStudent(studentID string) (*entity.Student, error)
	UpdateStudent(studentID string, studentDTO *transport.Student) error
	DeleteStudent(studentID string) error
}
