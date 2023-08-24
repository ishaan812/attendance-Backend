package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"service/database"

	"github.com/gorilla/mux"
)

func CreateFaculty(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Content-Type", "application/json")
	var faculty database.Faculty
	json.NewDecoder(r.Body).Decode(&faculty)
	err := dbconn.Create(&faculty)
	if err != nil {
		json.NewEncoder(w).Encode(err.Error)
	} else {
		json.NewEncoder(w).Encode("Faculty Created")
	}
}

func GetAllFaculties(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Content-Type", "application/json")
	var faculty []database.Faculty
	err := dbconn.Find(&faculty).Error
	if err != nil {
		json.NewEncoder(w).Encode("No Faculties Found")
	}
	json.NewEncoder(w).Encode(&faculty)
}

func GetFacultyByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	var faculty database.Faculty
	err := dbconn.Where("id = ?", params["id"]).First(&faculty).Error
	fmt.Println(faculty)
	if err != nil {
		json.NewEncoder(w).Encode("Invalid ID")
	} else {
		json.NewEncoder(w).Encode(&faculty)
	}
}

func DeleteFaculty(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Content-Type", "application/json")
	var faculty database.Faculty
	params := mux.Vars(r)
	err := dbconn.Where("id = ?", params["id"]).First(&faculty).Error
	if err != nil {
		json.NewEncoder(w).Encode("Invalid ID")
	} else {
		dbconn.Delete(&faculty)
		json.NewEncoder(w).Encode("Student Deleted")
	}
}

func UpdateFaculty(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	var faculty database.Faculty
	err := dbconn.Where("id = ?", params["id"]).First(&faculty).Error
	if err != nil {
		json.NewEncoder(w).Encode("Invalid ID")
	}
	json.NewDecoder(r.Body).Decode(&faculty)
	dbconn.Save(&faculty)
	json.NewEncoder(w).Encode(&faculty)
}

func GetSubjectsByFaculty(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	var faculty database.Faculty
	err := dbconn.Where("id = ?", params["id"]).First(&faculty).Error
	if err != nil {
		json.NewEncoder(w).Encode("Invalid ID")
	} else {
		for i := 0; i < len(faculty.Subjects); i++ {
			dbconn.Where("id = ?", faculty.Subjects[i]).First(&faculty.Subjects[i])
		}
	}
}
