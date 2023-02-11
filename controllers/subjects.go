package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"service/database"

	"github.com/gorilla/mux"
)

func CreateSubject(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Content-Type", "application/json")
	var subject database.Subject
	json.NewDecoder(r.Body).Decode(&subject)
	err := dbconn.Create(&subject)
	if err != nil {
		json.NewEncoder(w).Encode(err.Error)
	} else {
		json.NewEncoder(w).Encode("Subject Created")
	}
}

func GetAllSubjects(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Content-Type", "application/json")
	var subject []database.Subject
	model := dbconn.Find(&subject)
	paginated := pg.Response(model, r, &[]database.Subject{})
	json.NewEncoder(w).Encode(&paginated)
}

func GetSubjectByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	var subject database.Subject
	err := dbconn.Preload("Faculty").Where("id = ?", params["id"]).First(&subject).Error
	fmt.Println(subject)
	if err != nil {
		json.NewEncoder(w).Encode("Invalid ID")
	} else {
		json.NewEncoder(w).Encode(&subject)
	}
}

func DeleteSubject(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Content-Type", "application/json")
	var subject database.Subject
	params := mux.Vars(r)
	err := dbconn.Where("id = ?", params["id"]).First(&subject).Error
	if err != nil {
		json.NewEncoder(w).Encode("Invalid ID")
	} else {
		dbconn.Delete(&subject)
		json.NewEncoder(w).Encode("Student Deleted")
	}
}

func UpdateSubject(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	var subject database.Subject
	err := dbconn.Where("id = ?", params["id"]).First(&subject).Error
	if err != nil {
		json.NewEncoder(w).Encode("Invalid ID")
	}
	json.NewDecoder(r.Body).Decode(&subject)
	dbconn.Save(&subject)
	json.NewEncoder(w).Encode(&subject)
}
