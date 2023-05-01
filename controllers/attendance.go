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
		err = dbconn.FirstOrCreate(&StudentLecture).Error
		if err != nil {
			json.NewEncoder(w).Encode(err)
		} else {
			dbconn.Save(&StudentLecture)
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

	var Report StudentAttendanceReport
	var SubAttendances []float64
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
		if SubAttendance.TotalLectures == 0 {
			json.NewEncoder(w).Encode("No Lectures for this subject")
		} else {
			if SubAttendance.TotalLectures != 0 {
				SubAttendance.Attendance = (float64(SubAttendance.AttendedLectures) / float64(SubAttendance.TotalLectures)) * 100
			} else {
				SubAttendance.Attendance = 0
			}
			SubAttendances = append(SubAttendances, SubAttendance.Attendance)
			Report.SubjectAttendance = append(Report.SubjectAttendance, SubAttendance)
		}
	}
	var res float64
	for i := 0; i < len(SubAttendances); i++ {
		res += SubAttendances[i]
	}
	if len(Subjects) != 0 {
		Report.GrandAttendance = res / float64(len(Subjects))
		if Report.GrandAttendance < 75 {
			Report.Status = "Defaulter"
		} else {
			Report.Status = "Eligible"
		}
		json.NewEncoder(w).Encode(&Report)
	} else {
		json.NewEncoder(w).Encode("No Subjects")
	}
}

//input: year and division
//output: list of students with their attendance in different subjects

func GetAttendanceByYearandDivision(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Content-Type", "application/json")
	var ClassAttendanceRequest ClassAttendanceReq
	json.NewDecoder(r.Body).Decode(&ClassAttendanceRequest)
	var Students []database.Student
	var Subjects []database.Subject
	//get list of students
	err := dbconn.Where("year = ? AND division = ?", ClassAttendanceRequest.Year, ClassAttendanceRequest.Division).Find(&Students).Error
	if err != nil {
		json.NewEncoder(w).Encode("Wrong year or division")
	}
	//get list of subjects
	err = dbconn.Where("year = ?", ClassAttendanceRequest.Year).Find(&Subjects).Error
	if err != nil {
		json.NewEncoder(w).Encode("Wrong year")
	}
	var Report DivisionReport
	Report.Year = ClassAttendanceRequest.Year
	Report.Division = ClassAttendanceRequest.Division
	//get attendance of each student in each subject
	for i := 0; i < len(Students); i++ {
		var StudentReport StudentAttendanceReport
		var SubAttendances []float64
		StudentReport.SAPID = Students[i].SAPID
		StudentReport.StudentName = Students[i].Name
		for j := 0; j < len(Subjects); j++ {
			var SubAttendance SubjectAttendance
			SubAttendance.SubjectName = Subjects[j].Name
			SubAttendance.SubjectCode = Subjects[j].SubjectCode
			var TotalLectures []database.StudentLecture
			var AttendedLectures []database.StudentLecture
			err := dbconn.Preload("Lecture").Select("DISTINCT lecture_id").Where("subject_id = ?", Subjects[j].ID).Find(&TotalLectures).Error
			if err != nil {
				fmt.Println(err)
			}

			SubAttendance.TotalLectures = len(TotalLectures)
			err = dbconn.Preload("Lecture").Where("student_id = ? AND subject_id = ?", Students[i].ID, Subjects[j].ID).Find(&AttendedLectures).Error
			if err != nil {
				fmt.Println(err)
			}
			SubAttendance.AttendedLectures = len(AttendedLectures)
			if SubAttendance.TotalLectures == 0 {
				json.NewEncoder(w).Encode("No Lectures for this subject")
			} else {
				if SubAttendance.TotalLectures != 0 {
					SubAttendance.Attendance = (float64(SubAttendance.AttendedLectures) / float64(SubAttendance.TotalLectures)) * 100
				} else {
					SubAttendance.Attendance = 0
				}
				SubAttendances = append(SubAttendances, SubAttendance.Attendance)
				StudentReport.SubjectAttendance = append(StudentReport.SubjectAttendance, SubAttendance)
			}
		}

		var res float64
		for i := 0; i < len(SubAttendances); i++ {

			res += SubAttendances[i]

		}
		if len(Subjects) != 0 {
			StudentReport.GrandAttendance = res / float64(len(Subjects))
			if StudentReport.GrandAttendance < 75 {
				StudentReport.Status = "Defaulter"
			} else {
				StudentReport.Status = "Eligible"
			}
		} else {
			json.NewEncoder(w).Encode("No Subjects")
		}
		Report.AttendanceList = append(Report.AttendanceList, StudentReport)
	}
	json.NewEncoder(w).Encode(&Report)
}
