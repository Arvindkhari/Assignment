package controllers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"
	"student_portal/database"
	"student_portal/entity"

	"github.com/gorilla/mux"
)

//CreateStudent creates student
func CreateStudent(w http.ResponseWriter, r *http.Request) {
	requestBody, _ := ioutil.ReadAll(r.Body)
	var student entity.Student
	json.Unmarshal(requestBody, &student)

	database.Connector.Create(student)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(student)
}

//GetAllStudent to get all student data
func GetAllStudent(w http.ResponseWriter, r *http.Request) {
	var students []entity.Student
	database.Connector.Find(&students)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(students)
}

//GetStudentByID returns Student with specific ID
func GetStudentByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["id"]

	var student entity.Student
	database.Connector.First(&student, key)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(student)
}

//
//GetStudentByYearOfJoining returns Student with specific students
func GetAllStudentByYearOfJoining(w http.ResponseWriter, r *http.Request) {
	var students []entity.Student
	database.Connector.Find(&students)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(students)
}

//UpdateMobileNumberByID will updates student with respective ID
func UpdateMobileNumberByID(w http.ResponseWriter, r *http.Request) {
	requestBody, _ := ioutil.ReadAll(r.Body)
	var student entity.Student
	json.Unmarshal(requestBody, &student)
	database.Connector.Save(&student)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(student)
}

//DeleteStudentByID will delete's student with specific ID
func DeletStudentByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["id"]

	var student entity.Student
	id, _ := strconv.ParseInt(key, 10, 64)
	database.Connector.Where("id = ?", id).Delete(&student)
	w.WriteHeader(http.StatusNoContent)
}
