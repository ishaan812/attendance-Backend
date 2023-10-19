package routes

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"service/controllers"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func InitializeRouter() {

	r := mux.NewRouter()
	headers := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization", "Set-Cookie"})
	methods := handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "OPTIONS"})
	origins := handlers.AllowedOrigins([]string{os.Getenv("FRONTEND_URL"), "http://localhost:3000"})
	credentials := handlers.AllowCredentials()

	r.Handle("/student", http.HandlerFunc(controllers.CreateStudent)).Methods("POST", "OPTIONS")
	r.Handle("/getAllStudents", http.HandlerFunc(controllers.GetAllStudents)).Methods("GET")
	r.Handle("/getAllStudentsBySubject/{subject_code}", http.HandlerFunc(controllers.GetAllStudentsBySubject)).Methods("GET")
	r.Handle("/student/{id}", http.HandlerFunc(controllers.GetStudentByID)).Methods("GET")
	r.Handle("/student/{id}", http.HandlerFunc(controllers.UpdateStudent)).Methods("PUT")
	r.Handle("/student/{id}", http.HandlerFunc(controllers.DeleteStudent)).Methods("DELETE")

	r.Handle("/timetableentry", http.HandlerFunc(controllers.CreateTimeTableEntry)).Methods("POST", "OPTIONS")
	r.Handle("/getAllTimeTableEntries", http.HandlerFunc(controllers.GetAllTimeTableEntries)).Methods("GET")
	r.Handle("/getAllTimeTableEntries/{id}", http.HandlerFunc(controllers.GetAllTimeTableEntriesforFaculty)).Methods("GET")
	r.Handle("/timetableentry/{id}", http.HandlerFunc(controllers.GetTimeTableEntryByID)).Methods("GET")
	r.Handle("/timetableentry/{id}", http.HandlerFunc(controllers.UpdateTimeTableEntry)).Methods("PUT")
	r.Handle("/timetableentry/{id}", http.HandlerFunc(controllers.DeleteTimeTableEntry)).Methods("DELETE")

	r.Handle("/faculty", http.HandlerFunc(controllers.CreateFaculty)).Methods("POST", "OPTIONS")
	r.Handle("/getAllFaculties", http.HandlerFunc(controllers.GetAllFaculties)).Methods("GET")
	r.Handle("/faculty/{id}", http.HandlerFunc(controllers.GetFacultyByID)).Methods("GET")
	r.Handle("/faculty/{id}", http.HandlerFunc(controllers.UpdateFaculty)).Methods("PUT")
	r.Handle("/faculty/{id}", http.HandlerFunc(controllers.DeleteFaculty)).Methods("DELETE")

	r.Handle("/subject", http.HandlerFunc(controllers.CreateSubject)).Methods("POST", "OPTIONS")
	r.Handle("/getAllSubjects", http.HandlerFunc(controllers.GetAllSubjects)).Methods("GET")
	r.Handle("/subject/{id}", http.HandlerFunc(controllers.GetSubjectByID)).Methods("GET")
	r.Handle("/subject/{id}", http.HandlerFunc(controllers.UpdateSubject)).Methods("PUT")
	r.Handle("/subject/{id}", http.HandlerFunc(controllers.DeleteSubject)).Methods("DELETE")
	r.Handle("/subject/{code}", http.HandlerFunc(controllers.GetSubjectBySubjectCode)).Methods("GET")

	r.Handle("/lecture", http.HandlerFunc(controllers.CreateLecture)).Methods("POST", "OPTIONS")
	r.Handle("/lecture/{subject_code}", http.HandlerFunc(controllers.CreateLecturewithSubjectCode)).Methods("POST", "OPTIONS")
	r.Handle("/getAllLectures", http.HandlerFunc(controllers.GetAllLectures)).Methods("GET")
	r.Handle("/lecture/{id}", http.HandlerFunc(controllers.GetLectureByID)).Methods("GET")
	r.Handle("/lecture/{id}", http.HandlerFunc(controllers.UpdateLecture)).Methods("PUT")
	r.Handle("/lecture/{id}", http.HandlerFunc(controllers.DeleteLecture)).Methods("DELETE")
	r.Handle("/getLecturesBySubject/{subject_code}", http.HandlerFunc(controllers.GetLecturesBySubject)).Methods("GET")
	r.Handle("/getLecturesByFaculty/{id}", http.HandlerFunc(controllers.GetLecturesByFaculty)).Methods("GET")
	r.Handle("/fetchLecture", http.HandlerFunc(controllers.FetchLecture)).Methods("POST", "OPTIONS")

	r.HandleFunc("/register", controllers.Register).Methods("POST", "OPTIONS")
	r.HandleFunc("/login", controllers.Login).Methods("POST", "OPTIONS")
	r.HandleFunc("/logout", controllers.Logout).Methods("GET")
	// r.HandleFunc("/refresh", controllers.Refresh).Methods("GET")

	r.Handle("/markAttendance", http.HandlerFunc(controllers.MarkAttendance)).Methods("PUT")
	r.Handle("/getLectureAttendance/{id}", http.HandlerFunc(controllers.GetLectureAttendance)).Methods("GET")
	r.Handle("/getSubjectsbyFaculty/{id}", http.HandlerFunc(controllers.GetSubjectsByFaculty)).Methods("GET")
	// r.Handle("/getStudentAttendance", http.HandlerFunc(controllers.GetAttendanceBySAPID)).Methods("POST", "OPTIONS")
	r.Handle("/getClassAttendance", http.HandlerFunc(controllers.GetAttendanceByYearandDivision)).Methods("POST", "OPTIONS")

	fmt.Println("Server running on localhost:" + os.Getenv("PORT") + "\n")
	serverErr := http.ListenAndServe(":"+os.Getenv("PORT"), handlers.CORS(headers, methods, origins, credentials)(r))
	if serverErr != nil {
		log.Fatal(serverErr)
	}
}
