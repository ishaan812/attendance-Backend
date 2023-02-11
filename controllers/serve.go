package controllers

import (
	"github.com/morkid/paginate"
	"gorm.io/gorm"
)

var dbconn *gorm.DB
var pg = paginate.New()

func GetDB(db *gorm.DB) {
	dbconn = db
}



