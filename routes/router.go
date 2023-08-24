package routes

import (
	"fmt"
	"log"
	"net/http"
	"service/controllers"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func InitializeRouter() {

	r := mux.NewRouter()
	headers := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"})
	methods := handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "OPTIONS"})
	origins := handlers.AllowedOrigins([]string{"*"})

	r.HandleFunc("/student", controllers.CreateStudent).Methods("POST", "OPTIONS")
	r.HandleFunc("/getAllStudents", controllers.GetAllStudents).Methods("GET")
	r.HandleFunc("/getAllStudentsBySubject/{subject_code}", controllers.GetAllStudentsBySubject).Methods("GET")
	r.HandleFunc("/student/{id}", controllers.GetStudentByID).Methods("GET")
	r.HandleFunc("/student/{id}", controllers.UpdateStudent).Methods("PUT")
	r.HandleFunc("/student/{id}", controllers.DeleteStudent).Methods("DELETE")

	r.HandleFunc("/timetableentry", controllers.CreateTimeTableEntry).Methods("POST", "OPTIONS")
	r.HandleFunc("/getAllTimeTableEntries", controllers.GetAllTimeTableEntries).Methods("GET")
	r.HandleFunc("/getAllTimeTableEntries/{id}", controllers.GetAllTimeTableEntriesforFaculty).Methods("GET")
	r.HandleFunc("/timetableentry/{id}", controllers.GetTimeTableEntryByID).Methods("GET")
	r.HandleFunc("/timetableentry/{id}", controllers.UpdateTimeTableEntry).Methods("PUT")
	r.HandleFunc("/timetableentry/{id}", controllers.DeleteTimeTableEntry).Methods("DELETE")

	r.HandleFunc("/faculty", controllers.CreateFaculty).Methods("POST", "OPTIONS")
	r.HandleFunc("/getAllFaculties", controllers.GetAllFaculties).Methods("GET")
	r.HandleFunc("/faculty/{id}", controllers.GetFacultyByID).Methods("GET")
	r.HandleFunc("/faculty/{id}", controllers.UpdateFaculty).Methods("PUT")
	r.HandleFunc("/faculty/{id}", controllers.DeleteFaculty).Methods("DELETE")

	r.HandleFunc("/subject", controllers.CreateSubject).Methods("POST", "OPTIONS")
	r.HandleFunc("/getAllSubjects", controllers.GetAllSubjects).Methods("GET")
	r.HandleFunc("/subject/{id}", controllers.GetSubjectByID).Methods("GET")
	r.HandleFunc("/subject/{id}", controllers.UpdateSubject).Methods("PUT")
	r.HandleFunc("/subject/{id}", controllers.DeleteSubject).Methods("DELETE")
	r.HandleFunc("/subject/{code}", controllers.GetSubjectBySubjectCode).Methods("GET")

	r.HandleFunc("/lecture", controllers.CreateLecture).Methods("POST", "OPTIONS")
	r.HandleFunc("/lecture/{subject_code}", controllers.CreateLecturewithSubjectCode).Methods("POST", "OPTIONS")
	r.HandleFunc("/getAllLectures", controllers.GetAllLectures).Methods("GET")
	r.HandleFunc("/lecture/{id}", controllers.GetLectureByID).Methods("GET")
	r.HandleFunc("/lecture/{id}", controllers.UpdateLecture).Methods("PUT")
	r.HandleFunc("/lecture/{id}", controllers.DeleteLecture).Methods("DELETE")
	r.HandleFunc("/getLecturesBySubject/{subject_code}", controllers.GetLecturesBySubject).Methods("GET")
	r.HandleFunc("/getLecturesByFaculty/{id}", controllers.GetLecturesByFaculty).Methods("GET")
	r.HandleFunc("/fetchLecture", controllers.FetchLecture).Methods("POST", "OPTIONS")

	r.HandleFunc("/register", controllers.Register).Methods("POST", "OPTIONS")
	r.HandleFunc("/login", controllers.Login).Methods("POST", "OPTIONS")
	r.HandleFunc("/logout", controllers.Logout).Methods("GET")
	r.HandleFunc("/refresh", controllers.Refresh).Methods("GET")

	r.HandleFunc("/markAttendance", controllers.MarkAttendance).Methods("PUT")
	r.HandleFunc("/getLectureAttendance/{id}", controllers.GetLectureAttendance).Methods("POST", "OPTIONS")
	r.HandleFunc("/getSubjectbyFacultyID/{id}", controllers.GetSubjectsByFaculty).Methods("GET")
	r.HandleFunc("/getStudentAttendance", controllers.GetAttendanceBySAPID).Methods("POST", "OPTIONS")
	r.HandleFunc("/getClassAttendance", controllers.GetAttendanceByYearandDivision).Methods("POST", "OPTIONS")

	fmt.Print("Server running on localhost:9000\n")
	serverErr := http.ListenAndServe("localhost:9000", handlers.CORS(headers, methods, origins)(r))
	// serverErr := http.ListenAndServe("192.168.155.165:9000", handlers.CORS(headers, methods, origins)(r))
	if serverErr != nil {
		log.Fatal(serverErr)
	}
}
