package dao

import (
	"github.com/ppsingh4/stu-mini/db"
	"github.com/ppsingh4/stu-mini/entity"
)

type StudentDAO interface {
	Create(student *entity.Student) (uint, error)
	GetByID(id string) (*entity.Student, error)
	Update(id string, student *entity.Student) error
	Delete(id string) error
}

// StudentDAOImpl is the implementation of StudentDAO.
type StudentDAOImpl struct {
	db *db.Database
}

// NewStudentDAO creates a new instance of StudentDAO.
func NewStudentDAO(database *db.Database) StudentDAO {
	return &StudentDAOImpl{db: database}
}

// Create creates a new student record in the database.
func (d *StudentDAOImpl) Create(student *entity.Student) (uint, error) {
	err := d.db.Create(student).Error
	if err != nil {
		return 0, err
	}
	return student.ID, nil
}

// GetByID retrieves a student record from the database by ID.
func (d *StudentDAOImpl) GetByID(id string) (*entity.Student, error) {
	var student entity.Student
	err := d.db.Where("id = ?", id).First(&student).Error
	if err != nil {
		return nil, err
	}
	return &student, nil
}

// Update updates an existing student record in the database.
func (d *StudentDAOImpl) Update(id string, student *entity.Student) error {
	err := d.db.Model(&entity.Student{}).Where("id = ?", id).Updates(student).Error
	if err != nil {
		return err
	}
	return nil
}

// Delete deletes a student record from the database by ID.
func (d *StudentDAOImpl) Delete(id string) error {
	err := d.db.Where("id = ?", id).Delete(&entity.Student{}).Error
	if err != nil {
		return err
	}
	return nil
}
