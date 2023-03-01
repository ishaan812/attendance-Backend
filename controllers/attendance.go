package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"service/database"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
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

func GetLectureAttendance(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	var StudentLectures []database.StudentLecture
	json.NewDecoder(r.Body).Decode(&StudentLectures)
	dbconn.Preload("Student").Where("lecture_id = ?", params["id"]).Find(&StudentLectures)
	fmt.Println(StudentLectures)
	json.NewEncoder(w).Encode(StudentLectures)
}

// Add report generation
func GetAttendanceBySAPID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Content-Type", "application/json")
	//Destructuring Request
	var StudentAttendanceRequest StudentAttendanceReq
	var Student database.Student
	json.NewDecoder(r.Body).Decode(&StudentAttendanceRequest)
	//Find StudentID by SAPID
	err := dbconn.Where("s_api_d = ?", StudentAttendanceRequest.SAPID).First(&Student).Error
	if err != nil {
		json.NewEncoder(w).Encode("Wrong SAPID")
	}
	//Get Subjects for Student based on year
	var Subjects []database.Subject
	err = dbconn.Where("year = ?", Student.Year).Find(&Subjects).Error
	if err != nil {
		json.NewEncoder(w).Encode("Wrong year")
	}
	// print(Subjects[0].Name)

	var Report StudentAttendanceReport
	var SubAttendances []int
	Report.SAPID = Student.SAPID
	Report.StudentName = Student.Name

	for i := 0; i < len(Subjects); i++ {
		var SubAttendance SubjectAttendance
		SubAttendance.SubjectName = Subjects[i].Name
		SubAttendance.SubjectCode = Subjects[i].SubjectCode
		var TotalLectures []database.StudentLecture
		var AttendedLectures []database.StudentLecture
		err := dbconn.Preload("Lecture").Where("subject_id = ?", Subjects[i].ID).Find(&TotalLectures).Error
		if err != nil {
			fmt.Println(err)
		}
		SubAttendance.TotalLectures = len(TotalLectures)
		err = dbconn.Preload("Lecture").Where("student_id = ? AND subject_id = ?", Student.ID, Subjects[i].ID).Find(&AttendedLectures).Error
		if err != nil {
			fmt.Println(err)
		}
		SubAttendance.AttendedLectures = len(AttendedLectures)
		SubAttendance.Attendance = (SubAttendance.AttendedLectures / SubAttendance.TotalLectures) * 100
		SubAttendances = append(SubAttendances, SubAttendance.Attendance)
		Report.SubjectAttendance = append(Report.SubjectAttendance, SubAttendance)
	}
	var res int
	for i := 0; i < len(SubAttendances); i++ {
		res += SubAttendances[i]
	}
	Report.GrandAttendance = res / len(Subjects)

	json.NewEncoder(w).Encode(&Report)
	// 	// 	append totallectures, attended lectures, attendance% to json array
	// 	//	calculate grand total attendance
	// 	//	append to json
	// 	//	send json
}
