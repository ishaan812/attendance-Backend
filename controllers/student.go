package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"service/database"

	"github.com/gorilla/mux"
)

func CreateStudent(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	var student database.Student
	json.NewDecoder(r.Body).Decode(&student)
	err := dbconn.Create(&student)
	if err != nil {
		json.NewEncoder(w).Encode(err.Error)
	} else {
		json.NewEncoder(w).Encode("Student Created")
	}
}

func GetAllStudents(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	var student []database.Student
	year := r.URL.Query().Get("year")
	division := r.URL.Query().Get("division")
	batch := r.URL.Query().Get("batch")
	department := r.URL.Query().Get("department")
	if year == "" && division == "" && batch == "" && department == "" {
		dbconn.Find(&student)
		json.NewEncoder(w).Encode(&student)
	} else {
		querystring := ""
		if year != "" {
			querystring += "year = " + year
		}
		if division != "" {
			if querystring != "" {
				querystring = querystring + " AND "
			}
			querystring = querystring + "division = '" + division + "'"
		}
		if batch != "" {
			if querystring != "" {
				querystring = querystring + " AND "
			}
			querystring = querystring + "batch = " + batch
		}
		if department != "" {
			if querystring != "" {
				querystring = querystring + " AND "
			}
			querystring = querystring + "department = '" + department + "'"
		}
		fmt.Println(querystring)
		dbconn.Find(&student, querystring)
		json.NewEncoder(w).Encode(&student)
	}
}

func GetAllStudentsBySubject(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	var students []database.Student
	params := mux.Vars(r)
	year := r.URL.Query().Get("year")
	division := r.URL.Query().Get("division")
	batch := r.URL.Query().Get("batch")
	if year == "" && division == "" && batch == "" {
		dbconn.Find(&students)
		json.NewEncoder(w).Encode(&students)
	} else {
		querystring := "'" + params["subject_code"] + "' = any(subjects)"
		if year != "" {
			if querystring != "" {
				querystring = querystring + " AND "
			}
			querystring += "year = '" + year + "'"
		}
		if division != "" {
			if querystring != "" {
				querystring = querystring + " AND "
			}
			querystring = querystring + "division = '" + division + "'"
		}
		if batch != "" {
			if querystring != "" {
				querystring = querystring + " AND "
			}
			querystring = querystring + "batch = '" + batch + "'"
		}
		fmt.Println(querystring)
		err := dbconn.Preload("Lectures").Find(&students, querystring).Error
		if err != nil {
			json.NewEncoder(w).Encode(err)
		}
		json.NewEncoder(w).Encode(&students)
	}
}

func GetStudentByID(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	var student database.Student
	err := dbconn.Where("id = ?", params["id"]).First(&student).Error
	fmt.Println(student)
	if err != nil {
		json.NewEncoder(w).Encode("Invalid ID")
	} else {
		json.NewEncoder(w).Encode(&student)
	}
}

func DeleteStudent(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	var student database.Student
	params := mux.Vars(r)
	err := dbconn.Where("id = ?", params["id"]).First(&student).Error
	if err != nil {
		json.NewEncoder(w).Encode("Invalid ID")
	} else {
		dbconn.Delete(&student)
		json.NewEncoder(w).Encode("Student Deleted")
	}
}

func UpdateStudent(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	var student database.Student
	err := dbconn.Where("id = ?", params["id"]).First(&student).Error
	if err != nil {
		json.NewEncoder(w).Encode("Invalid ID")
	}
	json.NewDecoder(r.Body).Decode(&student)
	dbconn.Save(&student)
	json.NewEncoder(w).Encode(&student)
}
