package controllers

import (
	"encoding/json"
	"fmt"
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
		json.NewEncoder(w).Encode("Lecture Created")
	}
}

func GetAllLectures(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Content-Type", "application/json")
	var lecture []database.Lecture
	model := dbconn.Find(&lecture)
	paginated := pg.Response(model, r, &[]database.Lecture{})
	json.NewEncoder(w).Encode(&paginated)
}

func GetLectureByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	var lecture database.Lecture
	err := dbconn.Preload("Subject").Where("id = ?", params["id"]).First(&lecture).Error
	fmt.Println(lecture)
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
