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
	origins := handlers.AllowedOrigins([]string{"http://localhost:3000"})
	credentials := handlers.AllowCredentials()

	r.Handle("/student", jwtMiddleware(http.HandlerFunc(controllers.CreateStudent))).Methods("POST", "OPTIONS")
	r.Handle("/getAllStudents", jwtMiddleware(http.HandlerFunc(controllers.GetAllStudents))).Methods("GET")
	r.Handle("/getAllStudentsBySubject/{subject_code}", jwtMiddleware(http.HandlerFunc(controllers.GetAllStudentsBySubject))).Methods("GET")
	r.Handle("/student/{id}", jwtMiddleware(http.HandlerFunc(controllers.GetStudentByID))).Methods("GET")
	r.Handle("/student/{id}", jwtMiddleware(http.HandlerFunc(controllers.UpdateStudent))).Methods("PUT")
	r.Handle("/student/{id}", jwtMiddleware(http.HandlerFunc(controllers.DeleteStudent))).Methods("DELETE")

	r.Handle("/timetableentry", jwtMiddleware(http.HandlerFunc(controllers.CreateTimeTableEntry))).Methods("POST", "OPTIONS")
	r.Handle("/getAllTimeTableEntries", jwtMiddleware(http.HandlerFunc(controllers.GetAllTimeTableEntries))).Methods("GET")
	r.Handle("/getAllTimeTableEntries/{id}", jwtMiddleware(http.HandlerFunc(controllers.GetAllTimeTableEntriesforFaculty))).Methods("GET")
	r.Handle("/timetableentry/{id}", jwtMiddleware(http.HandlerFunc(controllers.GetTimeTableEntryByID))).Methods("GET")
	r.Handle("/timetableentry/{id}", jwtMiddleware(http.HandlerFunc(controllers.UpdateTimeTableEntry))).Methods("PUT")
	r.Handle("/timetableentry/{id}", jwtMiddleware(http.HandlerFunc(controllers.DeleteTimeTableEntry))).Methods("DELETE")

	r.Handle("/faculty", jwtMiddleware(http.HandlerFunc(controllers.CreateFaculty))).Methods("POST", "OPTIONS")
	r.Handle("/getAllFaculties", jwtMiddleware(http.HandlerFunc(controllers.GetAllFaculties))).Methods("GET")
	r.Handle("/faculty/{id}", jwtMiddleware(http.HandlerFunc(controllers.GetFacultyByID))).Methods("GET")
	r.Handle("/faculty/{id}", jwtMiddleware(http.HandlerFunc(controllers.UpdateFaculty))).Methods("PUT")
	r.Handle("/faculty/{id}", jwtMiddleware(http.HandlerFunc(controllers.DeleteFaculty))).Methods("DELETE")

	r.Handle("/subject", jwtMiddleware(http.HandlerFunc(controllers.CreateSubject))).Methods("POST", "OPTIONS")
	r.Handle("/getAllSubjects", jwtMiddleware(http.HandlerFunc(controllers.GetAllSubjects))).Methods("GET")
	r.Handle("/subject/{id}", jwtMiddleware(http.HandlerFunc(controllers.GetSubjectByID))).Methods("GET")
	r.Handle("/subject/{id}", jwtMiddleware(http.HandlerFunc(controllers.UpdateSubject))).Methods("PUT")
	r.Handle("/subject/{id}", jwtMiddleware(http.HandlerFunc(controllers.DeleteSubject))).Methods("DELETE")
	r.Handle("/subject/{code}", jwtMiddleware(http.HandlerFunc(controllers.GetSubjectBySubjectCode))).Methods("GET")

	r.Handle("/lecture", jwtMiddleware(http.HandlerFunc(controllers.CreateLecture))).Methods("POST", "OPTIONS")
	r.Handle("/lecture/{subject_code}", jwtMiddleware(http.HandlerFunc(controllers.CreateLecturewithSubjectCode))).Methods("POST", "OPTIONS")
	r.Handle("/getAllLectures", jwtMiddleware(http.HandlerFunc(controllers.GetAllLectures))).Methods("GET")
	r.Handle("/lecture/{id}", jwtMiddleware(http.HandlerFunc(controllers.GetLectureByID))).Methods("GET")
	r.Handle("/lecture/{id}", jwtMiddleware(http.HandlerFunc(controllers.UpdateLecture))).Methods("PUT")
	r.Handle("/lecture/{id}", jwtMiddleware(http.HandlerFunc(controllers.DeleteLecture))).Methods("DELETE")
	r.Handle("/getLecturesBySubject/{subject_code}", jwtMiddleware(http.HandlerFunc(controllers.GetLecturesBySubject))).Methods("GET")
	r.Handle("/getLecturesByFaculty/{id}", jwtMiddleware(http.HandlerFunc(controllers.GetLecturesByFaculty))).Methods("GET")
	r.Handle("/fetchLecture", jwtMiddleware(http.HandlerFunc(controllers.FetchLecture))).Methods("POST", "OPTIONS")

	r.HandleFunc("/register", controllers.Register).Methods("POST", "OPTIONS")
	r.HandleFunc("/login", controllers.Login).Methods("POST", "OPTIONS")
	r.HandleFunc("/logout", controllers.Logout).Methods("GET")
	// r.HandleFunc("/refresh", controllers.Refresh).Methods("GET")

	r.Handle("/markAttendance", jwtMiddleware(http.HandlerFunc(controllers.MarkAttendance))).Methods("PUT")
	r.Handle("/getLectureAttendance/{id}", jwtMiddleware(http.HandlerFunc(controllers.GetLectureAttendance))).Methods("GET")
	r.Handle("/getSubjectsbyFaculty/{id}", jwtMiddleware(http.HandlerFunc(controllers.GetSubjectsByFaculty))).Methods("GET")
	r.Handle("/getStudentAttendance", jwtMiddleware(http.HandlerFunc(controllers.GetAttendanceBySAPID))).Methods("POST", "OPTIONS")
	r.Handle("/getClassAttendance", jwtMiddleware(http.HandlerFunc(controllers.GetAttendanceByYearandDivision))).Methods("POST", "OPTIONS")

	fmt.Println("Server running on localhost:" + os.Getenv("PORT") + "\n")
	serverErr := http.ListenAndServe(":"+os.Getenv("PORT"), handlers.CORS(headers, methods, origins, credentials)(r))
	// serverErr := http.ListenAndServe("192.168.155.165:9000", handlers.CORS(headers, methods, origins)(r))
	if serverErr != nil {
		log.Fatal(serverErr)
	}
}
