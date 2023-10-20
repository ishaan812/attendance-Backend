package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"service/database"
	"time"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

func MarkAttendance(w http.ResponseWriter, r *http.Request) {

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
		StudentLecture.SubjectID = AttendanceQuery.SubjectID
		err := dbconn.Where("s_api_d = ?", AttendanceQuery.Attendance[i]).First(&Student).Error
		if err != nil {
			json.NewEncoder(w).Encode("Invalid SAP ID")
		}
		StudentLecture.StudentID = Student.ID
		// StudentLecture.ID = Lecture.ID
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

	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	var StudentLectures []database.StudentLecture
	json.NewDecoder(r.Body).Decode(&StudentLectures)
	dbconn.Select("student_id").Preload("Student").Where("lecture_id = ?", params["id"]).Find(&StudentLectures)
	var StudentAttendance []int
	for i := 0; i < len(StudentLectures); i++ {
		StudentAttendance = append(StudentAttendance, StudentLectures[i].Student.SAPID)
	}
	json.NewEncoder(w).Encode(StudentAttendance)
}

// Add report generation
// func GetAttendanceBySAPID(w http.ResponseWriter, r *http.Request) {

// 	w.Header().Set("Content-Type", "application/json")
// 	//Destructuring Request
// 	var StudentAttendanceRequest StudentAttendanceReq
// 	var Student database.Student
// 	json.NewDecoder(r.Body).Decode(&StudentAttendanceRequest)
// 	//Find StudentID by SAPID
// 	err := dbconn.Where("s_api_d = ?", StudentAttendanceRequest.SAPID).First(&Student).Error
// 	if err != nil {
// 		json.NewEncoder(w).Encode("Wrong SAPID")
// 	}
// 	//Get Subjects for Student based on year
// 	var Subjects []database.Subject
// 	err = dbconn.Where("year = ?", Student.Year).Find(&Subjects).Error
// 	if err != nil {
// 		json.NewEncoder(w).Encode("Wrong year")
// 	}

// 	var Report StudentAttendanceReport
// 	var SubAttendances []float64
// 	Report.SAPID = Student.SAPID
// 	Report.StudentName = Student.Name

// 	for i := 0; i < len(Subjects); i++ {
// 		var SubAttendance SubjectAttendance
// 		SubAttendance.SubjectName = Subjects[i].Name
// 		SubAttendance.SubjectCode = Subjects[i].ID
// 		var TotalLectures []database.StudentLecture
// 		var AttendedLectures []database.StudentLecture
// 		err := dbconn.Preload("Lecture").Where("subject_id = ?", Subjects[i].ID).Find(&TotalLectures).Error
// 		if err != nil {
// 			fmt.Println(err)
// 		}
// 		SubAttendance.TotalLectures = len(TotalLectures)
// 		err = dbconn.Preload("Lecture").Where("student_id = ? AND subject_id = ? AND ", Student.ID, Subjects[i].ID).Find(&AttendedLectures).Error
// 		if err != nil {
// 			fmt.Println(err)
// 		}
// 		SubAttendance.AttendedLectures = len(AttendedLectures)
// 		if SubAttendance.TotalLectures == 0 {
// 			json.NewEncoder(w).Encode("No Lectures for this subject")
// 		} else {
// 			if SubAttendance.TotalLectures != 0 {
// 				SubAttendance.Attendance = (float64(SubAttendance.AttendedLectures) / float64(SubAttendance.TotalLectures)) * 100
// 			} else {
// 				SubAttendance.Attendance = 0
// 			}
// 			SubAttendances = append(SubAttendances, SubAttendance.Attendance)
// 			Report.SubjectAttendance = append(Report.SubjectAttendance, SubAttendance)
// 		}
// 	}
// 	var res float64
// 	for i := 0; i < len(SubAttendances); i++ {
// 		res += SubAttendances[i]
// 	}
// 	if len(Subjects) != 0 {
// 		Report.GrandAttendance = res / float64(len(Subjects))
// 		if Report.GrandAttendance < 75 {
// 			Report.Status = "Defaulter"
// 		} else {
// 			Report.Status = "Eligible"
// 		}
// 		json.NewEncoder(w).Encode(&Report)
// 	} else {
// 		json.NewEncoder(w).Encode("No Subjects")
// 	}
// }

// input: year and division
// output: list of students with their attendance in different subjects
func GetAttendanceByYearandDivision(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var err error
	var ClassAttendanceRequest ClassAttendanceReq
	json.NewDecoder(r.Body).Decode(&ClassAttendanceRequest)
	var Students []database.Student
	var Subjects []database.Subject

	//get list of students
	err = dbconn.Where("year = ? AND division = ?", ClassAttendanceRequest.Year, ClassAttendanceRequest.Division).Find(&Students).Error
	if err != nil {
		json.NewEncoder(w).Encode("Wrong year or division")
	}
	//get list of subjects
	err = dbconn.Where("year = ?", ClassAttendanceRequest.Year).Find(&Subjects).Error
	if err != nil {
		json.NewEncoder(w).Encode("Wrong year")
	}
	SubjectCodes := []string{}
	for i := 0; i < len(Subjects); i++ {
		SubjectCodes = append(SubjectCodes, Subjects[i].ID)
	}
	SubjectNames := []string{}
	for i := 0; i < len(Subjects); i++ {
		SubjectNames = append(SubjectNames, Subjects[i].Name)
	}
	var Report DivisionReport
	Report.Year = ClassAttendanceRequest.Year
	Report.Division = ClassAttendanceRequest.Division
	Report.Subjects = SubjectNames
	Report.StartDate, err = time.Parse("2006-01-02", ClassAttendanceRequest.StartDate)
	if err != nil {
		http.Error(w, "Invalid start date format", http.StatusBadRequest)
		return
	}
	Report.EndDate, err = time.Parse("2006-01-02", ClassAttendanceRequest.EndDate)
	if err != nil {
		http.Error(w, "Invalid end date format", http.StatusBadRequest)
		return
	}
	//get attendance of each student in each subject
	for i := 0; i < len(Students); i++ {
		var StudentReport StudentAttendanceReport
		var SubAttendances []float64
		StudentReport.SAPID = Students[i].SAPID
		StudentReport.StudentName = Students[i].Name
		StudentReport.Subjects = Students[i].Subjects

		for j := 0; j < len(Subjects); j++ {
			var SubAttendance SubjectAttendance
			SubAttendance.SubjectName = Subjects[j].Name
			SubAttendance.SubjectCode = Subjects[j].ID
			var TotalTheoryLectures []database.StudentLecture
			var TotalPracticalLectures []database.StudentLecture
			var TheoryLectures []database.StudentLecture
			var PracticalLectures []database.StudentLecture

			err := dbconn.Table("student_lectures").
				Joins("JOIN lectures ON student_lectures.lecture_id = lectures.id").
				Select("DISTINCT student_lectures.lecture_id").
				Where("student_lectures.subject_id = ? AND lectures.type = ? AND lectures.date_of_lecture BETWEEN ? AND ?", Subjects[j].ID, "theory", Report.StartDate, Report.EndDate).
				Find(&TotalTheoryLectures).Error

			if err != nil {
				fmt.Println(err)
			}

			err1 := dbconn.Table("student_lectures").
				Joins("JOIN lectures ON student_lectures.lecture_id = lectures.id").
				Select("DISTINCT student_lectures.lecture_id").
				Where("student_lectures.subject_id = ? AND lectures.type = ? AND lectures.batch = ? AND lectures.date_of_lecture BETWEEN ? AND ?", Subjects[j].ID, "practical", Students[i].Batch, Report.StartDate, Report.EndDate).
				Find(&TotalPracticalLectures).Error

			if err1 != nil {
				fmt.Println(err1)
			}

			SubAttendance.TotalTheoryLectures = len(TotalTheoryLectures)
			SubAttendance.TotalPracticalLectures = len(TotalPracticalLectures)

			err = dbconn.Table("student_lectures").
				Joins("JOIN lectures ON student_lectures.lecture_id = lectures.id").
				Where("student_lectures.attendance = true  AND student_lectures.student_id = ? AND lectures.type = ? AND student_lectures.subject_id = ? AND lectures.date_of_lecture BETWEEN ? AND ?", Students[i].ID, "theory", Subjects[j].ID, Report.StartDate, Report.EndDate).
				Find(&TheoryLectures).Error
			err2 := dbconn.Table("student_lectures").
				Joins("JOIN lectures ON student_lectures.lecture_id = lectures.id").
				Where("student_lectures.attendance = true  AND student_lectures.student_id = ? AND lectures.type = ? AND student_lectures.subject_id = ? AND lectures.date_of_lecture BETWEEN ? AND ?", Students[i].ID, "practical", Subjects[j].ID, Report.StartDate, Report.EndDate).
				Find(&PracticalLectures).Error

			if err != nil {
				fmt.Println(err)
			}

			if err2 != nil {
				fmt.Println(err)
			}

			SubAttendance.TheoryLectures = len(TheoryLectures)
			SubAttendance.PracticalLectures = len(PracticalLectures)
			if SubAttendance.TotalTheoryLectures == 0 {

				SubAttendance.TheoryAttendance = 0
			} else {

				SubAttendance.TheoryAttendance = (float64(SubAttendance.TheoryLectures) / float64(SubAttendance.TotalTheoryLectures)) * 100
			}

			if SubAttendance.TotalPracticalLectures == 0 {
				SubAttendance.PracticalAttendance = 0
			} else {
				SubAttendance.PracticalAttendance = (float64(SubAttendance.PracticalLectures) / float64(SubAttendance.TotalPracticalLectures)) * 100
			}

			if SubAttendance.TotalTheoryLectures == 0 && SubAttendance.TotalPracticalLectures == 0 {
				SubAttendances = append(SubAttendances, 0)
			} else if SubAttendance.TotalTheoryLectures == 0 && SubAttendance.TotalPracticalLectures != 0 {
				SubAttendances = append(SubAttendances, SubAttendance.PracticalAttendance)
			} else if SubAttendance.TotalTheoryLectures != 0 && SubAttendance.TotalPracticalLectures == 0 {
				SubAttendances = append(SubAttendances, SubAttendance.TheoryAttendance)
			} else {
				// SubAttendances = append(SubAttendances, (((float64(SubAttendance.TheoryLectures)/float64(SubAttendance.TotalTheoryLectures))+(float64(SubAttendance.PracticalLectures)/float64(SubAttendance.TotalPracticalLectures)))/(float64(SubAttendance.TotalTheoryLectures)+float64(SubAttendance.TotalPracticalLectures)))*100)
				SubAttendances = append(SubAttendances, ((float64(SubAttendance.TheoryLectures+SubAttendance.PracticalLectures))/float64(SubAttendance.TotalTheoryLectures+SubAttendance.TotalPracticalLectures))*100)
			}

			StudentReport.SubjectAttendance = append(StudentReport.SubjectAttendance, SubAttendance)

		}
		var res float64
		var SubjectMap []int

		// Array of Subattendences
		// fmt.Println(Students[i].Subjects)
		for z := 0; z < len(Students[i].Subjects); z++ {
			for y := 0; y < len(SubjectCodes); y++ {
				if Students[i].Subjects[z] == SubjectCodes[y] {
					SubjectMap = append(SubjectMap, y)

				}
			}
		}
		fmt.Println(SubjectMap)
		for k := 0; k < len(SubjectMap); k++ {

			res += SubAttendances[SubjectMap[k]]
			fmt.Println(SubAttendances)
		}

		ActiveSubjects := 0
		for m := 0; m < len(SubjectMap); m++ {
			if StudentReport.SubjectAttendance[SubjectMap[m]].TotalTheoryLectures != 0 || StudentReport.SubjectAttendance[SubjectMap[m]].TotalPracticalLectures != 0 {
				ActiveSubjects++
			}
		}
		if len(Subjects) != 0 && ActiveSubjects != 0 {
			fmt.Println(float64(ActiveSubjects))
			StudentReport.GrandAttendance = res / float64(ActiveSubjects)
			if StudentReport.GrandAttendance < 75 {
				StudentReport.Status = "Defaulter"
			} else {
				StudentReport.Status = "Eligible"
			}
		} else {
			StudentReport.GrandAttendance = -1
			StudentReport.Status = "-"
		}
		Report.AttendanceList = append(Report.AttendanceList, StudentReport)
	}
	// json.NewEncoder(w).Encode(&Students)
	// json.NewEncoder(w).Encode(&Subjects)
	json.NewEncoder(w).Encode(&Report)
}
