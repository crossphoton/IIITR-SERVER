package auth

import (
	"log"
	"os"

	"example.com/studentdata"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB
var err error

func saveUser(s studentdata.Student) {
	err = db.DB().Ping()
	if err != nil {
		dbInit()
		log.Println("Reconnecting to DB")
	}

	db.Save(&s)
}

func dbInit() {
	addr := os.Getenv("DB_URL")
	db, err = gorm.Open("postgres", addr)

	if err != nil {
		log.Fatal(err)
		log.Fatal("DB Error")
	}

	db.AutoMigrate(&studentdata.Student{})
	db.AutoMigrate(&studentdata.AIMSAcademicData{})
}