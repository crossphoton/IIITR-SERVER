/*
This application is an extension to [iiitr-services/Server](https://github.com/iiitr-services/Server)

StudentData package implements methods and endpoints related to student data.

Currently supported methods
	- Store AIMS data 			(for students)
	- Get AIMS data 			(for students)

This extension assumes these environment variables to be available (besides from main Server)
	- JWT_SIGNING_KEY			- For JWT formation
*/
package studentdata

import (
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB
var err error

// Handler handles http requests
func Handler(r *mux.Router, database *gorm.DB) {
	db = database
	db.AutoMigrate(AIMSAcademicData{})
	r.HandleFunc("/AIMSData", updateAIMSData).Methods("POST")
	r.HandleFunc("/AIMSData", getAIMSData).Methods("GET")
}
