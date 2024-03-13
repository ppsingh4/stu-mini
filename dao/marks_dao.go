package dao
import (
    "smod/entity"
    "smod/db"
)

type MarksDAO interface {
    Create(marks *entity.Marks) error
    GetByStudentID(studentID string) (*entity.Marks, error)
    UpdateByStudentID(studentID string, marks *entity.Marks) error
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
func (d *MarksDAOImpl) Create(marks *entity.Marks) error {
    err := d.db.Create(marks).Error
    if err != nil {
        return err
    }
    return nil
}
// GetByStudentID retrieves marks for a student from the database by student ID.
func (d *MarksDAOImpl) GetByStudentID(studentID string) (*entity.Marks, error) {
    var marks entity.Marks
    err := d.db.Where("student_id = ?", studentID).First(&marks).Error
    if err != nil {
        return nil, err
    }
    return &marks, nil
}
// UpdateByStudentID updates marks for a student in the database.
func (d *MarksDAOImpl) UpdateByStudentID(studentID string, marks *entity.Marks) error {
    err := d.db.Model(&entity.Marks{}).Where("student_id = ?", studentID).Updates(marks).Error
    if err != nil {
        return err
    }
    return nil
}
// DeleteByStudentID deletes marks for a student from the database by student ID.
func (d *MarksDAOImpl) DeleteByStudentID(studentID string) error {
    err := d.db.Where("student_id = ?", studentID).Delete(&entity.Marks{}).Error
    if err != nil {
        return err
    }
    return nil
}