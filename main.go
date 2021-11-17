package main

import (
	"fmt"
	"net"
	"time"
	_"github.com/gorilla/mux"
)

func DNSlookups(url string, pointer string) {

	ipRec, _ := net.LookupIP(url)
	for _, ip := range ipRec {
		fmt.Println("IP address is: ", ip)
	}

	canonicalName, _ := net.LookupCNAME(url)
	fmt.Println(canonicalName)

	ptr, _ := net.LookupAddr(pointer)
	for _, ptrvalue := range ptr {
		fmt.Println(ptrvalue)
	}

	nameServer, _ := net.LookupNS(url)
	for _, ns := range nameServer {
		fmt.Println(ns)
	}

	mxRecords, _ := net.LookupMX(url)
	for _, mx := range mxRecords {
		fmt.Println(mx.Host, mx.Pref)
	}
}

func main() {

	t1 := time.Now()

	t2 := time.Now()
	diff := t2.Sub(t1)
	fmt.Println("testMeee", diff)

	DNSlookups("gmail.com", "8.8.4.4")

}

////
func GetAllEmployees(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	employees := []model.Employee{}
	db.Find(&employees)
	respondJSON(w, http.StatusOK, employees)
}

func CreateEmployee(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	employee := model.Employee{}

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&employee); err != nil {
		respondError(w, http.StatusBadRequest, err.Error())
		return
	}
	defer r.Body.Close()

	if err := db.Save(&employee).Error; err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(w, http.StatusCreated, employee)
}

func GetEmployee(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	name := vars["name"]
	employee := getEmployeeOr404(db, name, w, r)
	if employee == nil {
		return
	}
	respondJSON(w, http.StatusOK, employee)
}
