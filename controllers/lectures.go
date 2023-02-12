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
	fmt.Println(params["id"])
	subject_id, err := uuid.Parse(params["id"])
	if err != nil {
		json.NewEncoder(w).Encode(err)
	}
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
