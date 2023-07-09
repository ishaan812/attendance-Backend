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
	SubjectCode string         `json:"subject_code,omitempty"`
	Name        string         `json:"name,omitempty"`
	Year        int            `json:"year,omitempty"`
	Department  string         `json:"department,omitempty"`
	FacultyID   uuid.UUID      `json:"faculty_id,omitempty"`
	Faculty     *Faculty       `gorm:"foreignkey:FacultyID" json:"faculty,omitempty"`
}

type Lecture struct {
	CreatedAt     time.Time
	UpdatedAt     time.Time
	DeletedAt     gorm.DeletedAt `gorm:"index" json:"-"`
	ID            uuid.UUID      `gorm:"primarykey;type:uuid;default:uuid_generate_v4()"`
	DateOfLecture string         `json:"date_of_lecture,omitempty" gorm:"type:date"`
	StartTime     string         `json:"start_time,omitempty"`
	EndTime       string         `json:"end_time,omitempty"`
	SubjectID     uuid.UUID      `json:"subject_id,omitempty"`
	Subject       *Subject       `gorm:"foreignkey:SubjectID" json:"subject,omitempty"`
	Type          string         `json:"type,omitempty"`
	Year          string         `json:"year_of_graduation,omitempty"`
	Division      string         `json:"division,omitempty"`
	Batch         int            `json:"batch,omitempty"`
	FacultyID     uuid.UUID      `json:"faculty_id,omitempty"`
	Faculty       *Faculty       `gorm:"foreignkey:FacultyID" json:"faculty,omitempty"`
}

type StudentLecture struct {
	CreatedAt  time.Time
	UpdatedAt  time.Time
	DeletedAt  gorm.DeletedAt `gorm:"index"`
	ID         uuid.UUID      `gorm:"index;type:uuid;default:uuid_generate_v4()"`
	StudentID  uuid.UUID      `gorm:"primarykey" json:"student_id"`
	Student    Student        `gorm:"foreignkey:StudentID" json:"student,omitempty"`
	SubjectID  uuid.UUID      `json:"subject_id"`
	Subject    Subject        `gorm:"foreignkey:SubjectID" json:"subject,omitempty"`
	LectureID  uuid.UUID      `gorm:"primarykey" json:"lecture_id"`
	Lecture    Lecture        `gorm:"foreignkey:LectureID" json:"lecture,omitempty"`
	Attendance bool           `json:"attendance"`
}
