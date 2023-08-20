package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"service/database"

	"github.com/gorilla/mux"
)

func CreateTimeTableEntry(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Content-Type", "application/json")
	var timetableentry database.TimeTableEntry
	json.NewDecoder(r.Body).Decode(&timetableentry)
	err := dbconn.Create(&timetableentry)
	if err != nil {
		json.NewEncoder(w).Encode(err.Error)
	} else {
		json.NewEncoder(w).Encode("TimeTableEntry Created")
	}
}

func GetAllTimeTableEntriesforFaculty(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Content-Type", "application/json")
	var timetableentries []database.TimeTableEntry
	var timetableresponse []TimeTableResponse
	params := mux.Vars(r)
	err := dbconn.Preload("Subject").Where("faculty_id = ?", params["id"]).Find(&timetableentries).Error
	if err != nil {
		json.NewEncoder(w).Encode("No Faculties Found")
	}
	for _, entry := range timetableentries {
		var subjectentry TimeTableResponse
		subjectentry.SubjectCode = entry.Subject.ID
		subjectentry.Day = entry.Day
		subjectentry.Type = entry.Type
		subjectentry.StartTime = entry.StartTime
		subjectentry.EndTime = entry.EndTime
		subjectentry.SubjectName = entry.Subject.Name
		subjectentry.Division = entry.Division
		subjectentry.Batch = entry.Batch
		subjectentry.Year = entry.Subject.Year
		timetableresponse = append(timetableresponse, subjectentry)
	}
	json.NewEncoder(w).Encode(&timetableresponse)
}

func GetAllTimeTableEntries(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Content-Type", "application/json")
	var timetableentry []database.TimeTableEntry
	err := dbconn.Preload("Subjects").Find(&timetableentry).Error
	if err != nil {
		json.NewEncoder(w).Encode("No Faculties Found")
	}
	json.NewEncoder(w).Encode(&timetableentry)
}

func GetTimeTableEntryByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	var timetableentry database.TimeTableEntry
	err := dbconn.Where("id = ?", params["id"]).First(&timetableentry).Error
	fmt.Println(timetableentry)
	if err != nil {
		json.NewEncoder(w).Encode("Invalid ID")
	} else {
		json.NewEncoder(w).Encode(&timetableentry)
	}
}

func DeleteTimeTableEntry(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Content-Type", "application/json")
	var timetableentry database.TimeTableEntry
	params := mux.Vars(r)
	err := dbconn.Where("id = ?", params["id"]).First(&timetableentry).Error
	if err != nil {
		json.NewEncoder(w).Encode("Invalid ID")
	} else {
		dbconn.Delete(&timetableentry)
		json.NewEncoder(w).Encode("Student Deleted")
	}
}

func UpdateTimeTableEntry(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	var timetableentry database.TimeTableEntry
	err := dbconn.Where("id = ?", params["id"]).First(&timetableentry).Error
	if err != nil {
		json.NewEncoder(w).Encode("Invalid ID")
	}
	json.NewDecoder(r.Body).Decode(&timetableentry)
	dbconn.Save(&timetableentry)
	json.NewEncoder(w).Encode(&timetableentry)
}
