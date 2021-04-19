/*
This is the package for the main Server of the services. Services are treated as extensions to the server.

The SERVER expects the following environment variables to be present
	- PORT 		- Port to serve on
	- DB_URL	- URL of a Postgres protocol supporting database

*/
package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/gorilla/mux"
	"github.com/iiitr-services/auth"
	"github.com/iiitr-services/studentdata"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var db *gorm.DB
var err error

// requestHandler make subrouters for extensions
func requestHandler() *mux.Router {
	handler := mux.NewRouter()

	handler.HandleFunc("/", home).Methods("GET")
	auth.Handler(handler.PathPrefix("/auth").Subrouter(), db)
	studentdata.Handler(handler.PathPrefix("/studentdata").Subrouter(), db)
	return handler
}

func main() {

	c := make(chan os.Signal, 1) /* Just a fun thing to do */
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		db.Close()
		fmt.Println("\nShutting down server.\nBye Bye...")
		os.Exit(0)
	}()

	dbInit()

	fmt.Println("Starting Up server on port " + os.Getenv("PORT"))
	log.Fatal(http.ListenAndServe(":"+os.Getenv("PORT"), requestHandler()))

}

// Home serves the homepage of the server
func home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "This is home. Here is the <a href=\"https://www.github.com/iiitr-services\">Github</a> url for consideration.")
}

// dbInit initializes the database
func dbInit() {
	addr := os.Getenv("DB_URL")
	db, err = gorm.Open("postgres", addr)

	if err != nil {
		log.Fatal(err)
		log.Fatal("DB Error")
	}

	db.AutoMigrate(&auth.Student{})
	db.AutoMigrate(&studentdata.AIMSAcademicData{})
}
