package database

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Student struct {
	CreatedAt  time.Time
	UpdatedAt  time.Time
	DeletedAt  gorm.DeletedAt `gorm:"index" json:"-"`
	ID         uuid.UUID      `gorm:"primarykey;type:uuid;default:uuid_generate_v4()"`
	SAPID      int            `json:"sap_id"`
	Name       string         `json:"name"`
	Year       int            `json:"year"`
	Email      string         `json:"email"`
	Department string         `json:"department"`
	Division   string         `json:"division"`
	Batch      int            `json:"batch"`
	Lectures   []*Lecture     `gorm:"many2many:student_lectures;foreignKey:ID;joinForeignKey:StudentID;" json:"lectures"`
}

type Faculty struct {
	CreatedAt  time.Time
	UpdatedAt  time.Time
	DeletedAt  gorm.DeletedAt `gorm:"index" json:"-"`
	ID         uuid.UUID      `gorm:"primarykey;type:uuid;default:uuid_generate_v4()"`
	SAPID      int            `json:"sap_id"`
	Password   string         `json:"password"`
	Name       string         `json:"name"`
	Email      string         `json:"email"`
	Department string         `json:"department"`
}

type Subject struct {
	CreatedAt  time.Time
	UpdatedAt  time.Time
	DeletedAt  gorm.DeletedAt `gorm:"index" json:"-"`
	ID         uuid.UUID      `gorm:"primarykey;type:uuid;default:uuid_generate_v4()"`
	FacultyID  uuid.UUID      `json:"faculty_id"`
	Faculty    *Faculty       `gorm:"foreignkey:FacultyID" json:"faculty"`
	Name       string         `json:"name"`
	Year       int            `json:"year"`
	Department string         `json:"department"`
	Semester   int            `json:"semester"`
}

type Lecture struct {
	CreatedAt     time.Time
	UpdatedAt     time.Time
	DeletedAt     gorm.DeletedAt `gorm:"index" json:"-"`
	ID            uuid.UUID      `gorm:"primarykey;type:uuid;default:uuid_generate_v4()"`
	SubjectID     uuid.UUID      `json:"subject_id"`
	Subject       Subject        `gorm:"foreignkey:SubjectID" json:"subject"`
	Type          string         `json:"type"`
	DateOfLecture time.Time      `json:"date_of_lecture"`
	Division      string         `json:"division"`
	Batch         int            `json:"batch"`
}

type StudentLecture struct {
	CreatedAt  time.Time
	UpdatedAt  time.Time
	DeletedAt  gorm.DeletedAt `gorm:"index"`
	ID         uuid.UUID      `gorm:"primarykey;type:uuid;default:uuid_generate_v4()"`
	StudentID  uuid.UUID      `json:"student_id"`
	Student    Student        `gorm:"foreignkey:StudentID" json:"user"`
	LectureID  uuid.UUID      `json:"lecture_id"`
	Lecture    Lecture        `gorm:"foreignkey:LectureID" json:"lecture"`
	Attendance bool           `json:"attendance"`
}
