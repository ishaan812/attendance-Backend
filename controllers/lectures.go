package controllers

import (
	"encoding/json"
	"net/http"
	"service/database"

	"github.com/gorilla/mux"
)

func CreateLecture(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Content-Type", "application/json")
	var lecture database.Lecture
	json.NewDecoder(r.Body).Decode(&lecture)
	err := dbconn.Create(&lecture)
	if err != nil {
		json.NewEncoder(w).Encode(err.Error)
	} else {
		json.NewEncoder(w).Encode(&lecture)
	}
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
