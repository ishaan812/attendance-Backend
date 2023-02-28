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
	Email      string         `json:"email"`
	Year       int            `json:"year"`
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
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"-"`
	ID          uuid.UUID      `gorm:"primarykey;type:uuid;default:uuid_generate_v4()"`
	SubjectCode string         `json:"subject_code"`
	Name        string         `json:"name"`
	Year        int            `json:"year"`
	Department  string         `json:"department"`
	Semester    int            `json:"semester"`
	FacultyID   uuid.UUID      `json:"faculty_id"`
	Faculty     *Faculty       `gorm:"foreignkey:FacultyID" json:"faculty"`
}

type Lecture struct {
	CreatedAt     time.Time
	UpdatedAt     time.Time
	DeletedAt     gorm.DeletedAt `gorm:"index" json:"-"`
	ID            uuid.UUID      `gorm:"primarykey;type:uuid;default:uuid_generate_v4()"`
	DateOfLecture string         `json:"date_of_lecture" gorm:"type:date"`
	StartTime     string         `json:"start_time"`
	EndTime       string         `json:"end_time"`
	SubjectID     uuid.UUID      `json:"subject_id"`
	Subject       *Subject       `gorm:"foreignkey:SubjectID" json:"subject"`
	Type          string         `json:"type"`
	Year          string         `json:"year_of_graduation"`
	Division      string         `json:"division"`
	Batch         int            `json:"batch"`
	FacultyID     uuid.UUID      `json:"faculty_id"`
	Faculty       *Faculty       `gorm:"foreignkey:FacultyID" json:"faculty"`
}

type StudentLecture struct {
	CreatedAt  time.Time
	UpdatedAt  time.Time
	DeletedAt  gorm.DeletedAt `gorm:"index"`
	ID         uuid.UUID      `gorm:"primarykey;type:uuid;default:uuid_generate_v4()"`
	StudentID  uuid.UUID      `json:"student_id"`
	Student    Student        `gorm:"foreignkey:StudentID" json:"student,omitempty"`
	SubjectID  uuid.UUID      `json:"subject_id"`
	Subject    Subject        `gorm:"foreignkey:SubjectID" json:"subject,omitempty"`
	LectureID  uuid.UUID      `json:"lecture_id"`
	Lecture    Lecture        `gorm:"foreignkey:LectureID" json:"lecture,omitempty"`
	Attendance bool           `json:"attendance"`
}
