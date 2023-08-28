package database

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// DB is a global variable to hold db connection

var DB *gorm.DB

func InitialMigration(DNS string) *gorm.DB {
	DNS = "postgres://wtqvvruq:JzaV0KL0JuldyFmKie-06OXkkSeNdLRV@rain.db.elephantsql.com/wtqvvruq"
	DB, err := gorm.Open(postgres.Open(DNS), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database")
	}
	if err := DB.SetupJoinTable(&Student{}, "Lectures", &StudentLecture{}); err != nil {
		println(err.Error())
		panic("Failed to setup join table: StudentLecture")
	}
	DB.Exec("CREATE EXTENSION IF NOT EXISTS \"uuid-ossp\"")
	DB.AutoMigrate(Student{}, Faculty{}, Subject{}, Lecture{}, TimeTableEntry{})
	return DB
}
