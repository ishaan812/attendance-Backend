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
	var err error
	json.NewDecoder(r.Body).Decode(&AttendanceQuery)
	LectureID, err := uuid.Parse(AttendanceQuery.LectureID)
	if err != nil {
		json.NewEncoder(w).Encode("Invalid Lecture ID")
	}
	for i := 0; i < len(AttendanceQuery.Attendance); i++ {
		var StudentLecture database.StudentLecture
		var Student database.Student
		StudentLecture.LectureID = LectureID
		err := dbconn.Where("s_api_d = ?", AttendanceQuery.Attendance[i]).First(&Student).Error
		if err != nil {
			json.NewEncoder(w).Encode("Invalid SAP ID")
		}
		StudentLecture.StudentID = Student.ID
		StudentLecture.Attendance = true
		err = dbconn.Create(&StudentLecture).Error
		if err != nil {
			json.NewEncoder(w).Encode(err)
		} else {
			json.NewEncoder(w).Encode("Attendance Marked")
		}
	}
}
