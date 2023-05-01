package controllers

import (
	"service/database"
)

func GetStudentAttendance(Student *database.Student) (*StudentAttendanceReport, error) {
	// var StudentLecture []database.StudentLecture
	var Result StudentAttendanceReport
	err := dbconn.Where("s_api_d = ?", Student.SAPID).First(&Student).Error
	if err != nil {
		return nil, err
	}
	Result.SAPID = Student.SAPID
	Result.StudentName = Student.Name
	return &Result, nil
}
