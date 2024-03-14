package dao

import (
	"errors"
	"github.com/ppsingh4/stu-mini/db"
	"github.com/ppsingh4/stu-mini/entity"

	"gorm.io/gorm"
)

type MarksDAO interface {
	Create(marks *entity.Mark) error
	GetByStudentID(studentID string) (*entity.Mark, error)
	UpdateByStudentID(studentID string, marks *entity.Mark) error
	DeleteByStudentID(studentID string) error
}

// MarksDAOImpl is the implementation of MarksDAO.
type MarksDAOImpl struct {
	db *db.Database
}

// NewMarksDAO creates a new instance of MarksDAO.
func NewMarksDAO(database *db.Database) MarksDAO {
	return &MarksDAOImpl{db: database}
}

// Create creates marks for a student in the database.
func (d *MarksDAOImpl) Create(marks *entity.Mark) error {
	err := d.db.Create(marks).Error
	if err != nil {
		return err
	}
	return nil
}

// GetByStudentID retrieves marks for a student from the database by student ID.
func (d *MarksDAOImpl) GetByStudentID(studentID string) (*entity.Mark, error) {
	var marks entity.Mark
	err := d.db.Where("id = ?", studentID).First(&marks).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, errors.New("record not found for that student")
	} else if err != nil {
		return nil, err
	}
	return &marks, nil
}

// UpdateByStudentID updates marks for a student in the database.
func (d *MarksDAOImpl) UpdateByStudentID(studentID string, marks *entity.Mark) error {
	err := d.db.Model(&entity.Mark{}).Where("id = ?", studentID).Updates(marks).Error
	if err != nil {
		return err
	}
	return nil
}

// DeleteByStudentID deletes marks for a student from the database by student ID.
func (d *MarksDAOImpl) DeleteByStudentID(studentID string) error {
	err := d.db.Where("id = ?", studentID).Delete(&entity.Mark{}).Error
	if err != nil {
		return err
	}
	return nil
}
