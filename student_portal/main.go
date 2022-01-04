package main

import (
	"log"
	"net/http"
	"student_portal/controllers"
	"student_portal/database"
	"student_portal/entity"

	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/mysql" //Required for MySQL dialect
)

func main() {
	initDB()
	log.Println("Starting the HTTP server on port 8090")

	router := mux.NewRouter().StrictSlash(true)
	initaliseHandlers(router)
	log.Fatal(http.ListenAndServe(":8090", router))
}

func initaliseHandlers(router *mux.Router) {
	router.HandleFunc("/create", controllers.CreateStudent).Methods("POST")
	router.HandleFunc("/get", controllers.GetAllStudent).Methods("GET")
	router.HandleFunc("/get/{yearOfJoining}", controllers.GetAllStudentByYearOfJoining).Methods("GET")
	router.HandleFunc("/get/{id}", controllers.GetStudentByID).Methods("GET")
	router.HandleFunc("/update/{id}", controllers.UpdateMobileNumberByID).Methods("PUT")
	router.HandleFunc("/delete/{id}", controllers.DeletStudentByID).Methods("DELETE")
}

func initDB() {
	config :=
		database.Config{
			ServerName: "localhost:3306",
			User:       "root",
			Password:   "root",
			DB:         "student_details",
		}

	connectionString := database.GetConnectionString(config)
	err := database.Connect(connectionString)
	if err != nil {
		panic(err.Error())
	}
	database.Migrate(&entity.Student{})
}
