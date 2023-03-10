package controllers

import (
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type Claims struct {
	SAPID      string    `json:"sap_id"`
	UserID     int       `json:"user_id"`
	Email      string    `json:"email"`
	Name       string    `json:"name"`
	Department string    `json:"department"`
	Expires    time.Time `json:"expires"`
	jwt.RegisteredClaims
}

// Attendance
type AttendanceQuery struct {
	LectureID  string `json:"lecture_id"`
	Attendance []int  `json:"attendance"`
}

// Student

type StudentAttendanceReq struct {
	SAPID int `json:"sap_id"`
}

// Attendance Report
type StudentAttendanceReport struct {
	SAPID             int                 `json:"student_id"`
	StudentName       string              `json:"student_name"`
	SubjectAttendance []SubjectAttendance `json: "subject_attendance"`
	GrandAttendance   int                 `json: "grand_attendance"`
}

type SubjectAttendance struct {
	SubjectName      string `json: "subject_name"`
	SubjectCode      string `json:"subject_code"`
	TotalLectures    int    `json: "total_lectures"`
	AttendedLectures int    `json: "attended_lectures"`
	Attendance       int    `json: "attendance"`
}

// Lecture
type FetchLectureReq struct {
	DateOfLecture string `json:"date_of_lecture"`
	Type          string `json:"type"`
	Division      string `json:"division"`
	Batch         int    `json:"batch"`
	FacultyID     string `json:"faculty_id"`
}
