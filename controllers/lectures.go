package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"service/database"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

func CreateLecture(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Content-Type", "application/json")
	var lecture database.Lecture
	json.NewDecoder(r.Body).Decode(&lecture)
	err := dbconn.Create(&lecture).Error
	if err != nil {
		json.NewEncoder(w).Encode(err.Error)
	}
	json.NewEncoder(w).Encode(&lecture)
}

func CreateLecturewithSubjectCode(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "*")
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	var lecture database.Lecture
	json.NewDecoder(r.Body).Decode(&lecture)
	fmt.Println(params["subject_code"])
	var subject database.Subject
	err := dbconn.Where("subject_code = ?", params["subject_code"]).First(&subject).Error
	if err != nil {
		json.NewEncoder(w).Encode(err.Error)
	}
	lecture.SubjectID = subject.ID
	err = dbconn.Create(&lecture).Error
	if err != nil {
		json.NewEncoder(w).Encode(err.Error)
	}
	json.NewEncoder(w).Encode(&lecture)
}

func GetAllLectures(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Content-Type", "application/json")
	var lectures []database.Lecture

	err := dbconn.Preload("Subject").Find(&lectures).Error
	if err != nil {
		json.NewEncoder(w).Encode(err)
	}
	json.NewEncoder(w).Encode(&lectures)
}

func GetLectureByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	var lecture database.Lecture
	err := dbconn.Preload("Subject").Preload("Faculty").Where("id = ?", params["id"]).First(&lecture).Error
	if err != nil {
		json.NewEncoder(w).Encode("Invalid ID")
	} else {
		json.NewEncoder(w).Encode(&lecture)
	}
}

func DeleteLecture(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Content-Type", "application/json")
	var lecture database.Lecture
	params := mux.Vars(r)
	err := dbconn.Where("id = ?", params["id"]).First(&lecture).Error
	if err != nil {
		json.NewEncoder(w).Encode("Invalid ID")
	} else {
		dbconn.Delete(&lecture)
		json.NewEncoder(w).Encode("Student Deleted")
	}
}

func UpdateLecture(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	var lecture database.Lecture
	err := dbconn.Where("id = ?", params["id"]).First(&lecture).Error
	if err != nil {
		json.NewEncoder(w).Encode("Invalid ID")
	}
	json.NewDecoder(r.Body).Decode(&lecture)
	dbconn.Save(&lecture)
	json.NewEncoder(w).Encode(&lecture)
}

func GetLecturesBySubject(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	var subject database.Subject
	err := dbconn.Where("id = ?", params["subject_code"]).First(&subject).Error
	if err != nil {
		json.NewEncoder(w).Encode(err)
	}
	subject_id := subject.ID
	var lectures []database.Lecture
	err = dbconn.Preload("Subject").Preload("Faculty").Where("subject_id = ?", subject_id).Find(&lectures).Error
	if err != nil {
		json.NewEncoder(w).Encode(err)
	}
	json.NewEncoder(w).Encode(&lectures)
}

func GetLecturesByFaculty(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	faculty_id, err := uuid.Parse(params["id"])
	if err != nil {
		json.NewEncoder(w).Encode(err)
	}
	var lectures []database.Lecture
	err = dbconn.Preload("Subject").Preload("Faculty").Where("faculty_id = ?", faculty_id).Find(&lectures).Error
	if err != nil {
		json.NewEncoder(w).Encode(err)
	}
	json.NewEncoder(w).Encode(&lectures)
}

func FetchLecture(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Content-Type", "application/json")
	var lectures database.Lecture
	var FetchLectureRequest FetchLectureReq
	json.NewDecoder(r.Body).Decode(&FetchLectureRequest)
	date_of_lecture := FetchLectureRequest.DateOfLecture
	type1 := FetchLectureRequest.Type
	division := FetchLectureRequest.Division
	batch := FetchLectureRequest.Batch
	querystring := ""
	if date_of_lecture != "" {
		querystring += "date_of_lecture = '" + date_of_lecture + "'"
	}
	if type1 != "" {
		if querystring != "" {
			querystring = querystring + " AND "
		}
		querystring = querystring + "type = '" + type1 + "'"
	}
	if division != "" {
		if querystring != "" {
			querystring = querystring + " AND "
		}
		querystring = querystring + "division = " + division
	}
	if batch != 0 {
		if querystring != "" {
			querystring = querystring + " AND "
		}
		batchstring := fmt.Sprintf("%d", batch)
		querystring = querystring + "batch = '" + batchstring + "'"
	}
	fmt.Println(date_of_lecture, type1, division, batch)
	err := dbconn.Preload("Subject").Preload("Faculty").Find(&lectures, querystring).Error
	if err != nil {
		json.NewEncoder(w).Encode(err)
	}
	json.NewEncoder(w).Encode(&lectures)
}
