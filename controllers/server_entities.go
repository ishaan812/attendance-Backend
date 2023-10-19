package controllers

import (
	"time"

	"github.com/golang-jwt/jwt"
)

type Claims struct {
	jwt.Claims
	SAPID  string `json:"sap_id"`
	UserID int    `json:"user_id"`
	Email  string `json:"email"`
	Name   string `json:"name"`
}

// Attendance
type AttendanceQuery struct {
	LectureID  string `json:"lecture_id"`
	Attendance []int  `json:"attendance"`
	SubjectID  string `json:"subject_code"`
}

// Student

type StudentAttendanceReq struct {
	SAPID int `json:"sap_id"`
}

// Attendance Report
type StudentAttendanceReport struct {
	SAPID             int                 `json:"student_id"`
	StudentName       string              `json:"student_name"`
	Subjects          []string            `json:"subjects"`
	SubjectAttendance []SubjectAttendance `json:"subject_attendance"`
	GrandAttendance   float64             `json:"grand_attendance"`
	Status            string              `json:"defaulter"`
}

type DivisionReport struct {
	Year           string                    `json:"year"`
	Subjects       []string                  `json:"subjects"`
	Division       string                    `json:"division"`
	StartDate      time.Time                 `json:"start_date"`
	EndDate        time.Time                 `json:"end_date"`
	AttendanceList []StudentAttendanceReport `json:"students"`
}

type ClassAttendanceReq struct {
	Year      string `json:"year"`
	Division  string `json:"division"`
	Batch     int    `json:"batch"`
	Type      string `json:"type"`
	StartDate string `json:"start_date"`
	EndDate   string `json:"end_date"`
}

type SubjectAttendance struct {
	SubjectName            string  `json:"subject_name"`
	SubjectCode            string  `json:"subject_code"`
	TotalTheoryLectures    int     `json:"total_theory"`
	TotalPracticalLectures int     `json:"total_practical"`
	TheoryLectures         int     `json:"attended_theoryLectures"`
	PracticalLectures      int     `json:"attended_practicalLectures"`
	TheoryAttendance       float64 `json:"attendance_theory"`
	PracticalAttendance    float64 `json:"attendance_practical"`
}

// Lecture
type FetchLectureReq struct {
	DateOfLecture string `json:"date_of_lecture"`
	Type          string `json:"type"`
	Division      string `json:"division"`
	Batch         int    `json:"batch"`
	FacultyID     string `json:"faculty_id"`
}

// Time Table Resposne
type TimeTableResponse struct {
	SubjectCode string `json:"subject_code"`
	Day         string `json:"day"`
	Type        string `json:"type"`
	StartTime   string `json:"start_time"`
	EndTime     string `json:"end_time"`
	SubjectName string `json:"subject_name"`
	Division    string `json:"division,omitempty"`
	Batch       int    `json:"batch,omitempty"`
	Year        string `json:"year,omitempty"`
}
