package controllers

import (
	"encoding/json"
	"net/http"
	"service/database"

	"github.com/google/uuid"
)

func MarkAttendance(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Content-Type", "application/json")
	var AttendanceQuery AttendanceQuery
	var Lecture database.Lecture
	var err error
	json.NewDecoder(r.Body).Decode(&AttendanceQuery)
	LectureID, err := uuid.Parse(AttendanceQuery.LectureID)
	if err != nil {
		json.NewEncoder(w).Encode("Invalid Lecture ID")
	}
	dbconn.Preload("Subject").Where("id = ?", LectureID).First(&Lecture)
	for i := 0; i < len(AttendanceQuery.Attendance); i++ {
		var StudentLecture database.StudentLecture
		var Student database.Student
		StudentLecture.LectureID = LectureID
		err := dbconn.Where("s_api_d = ?", AttendanceQuery.Attendance[i]).First(&Student).Error
		if err != nil {
			json.NewEncoder(w).Encode("Invalid SAP ID")
		}
		StudentLecture.StudentID = Student.ID
		StudentLecture.SubjectID = Lecture.SubjectID
		StudentLecture.Attendance = true
		err = dbconn.Create(&StudentLecture).Error
		if err != nil {
			json.NewEncoder(w).Encode(err)
		} else {
			json.NewEncoder(w).Encode("Attendance Marked")
		}
	}
}

// Add report generation
func GetStudentAttendanceReport(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Content-Type", "application/json")
	var Student database.Student
	var StudentLecture []database.StudentLecture
	// set := make(map[string]struct{})
	// var exists = struct{}{}
	json.NewDecoder(r.Body).Decode(&Student)
	err := dbconn.Where("s_api_d = ?", Student.SAPID).First(&Student).Error
	if err != nil {
		json.NewEncoder(w).Encode("Invalid SAP ID")
	}
	err = dbconn.Preload("Lecture").Preload("Subject").Where("student_id = ?", Student.ID).Find(&StudentLecture).Error
	if err != nil {
		json.NewEncoder(w).Encode("Error Fetching Attendance")
	}
	json.NewEncoder(w).Encode(&StudentLecture)
}
