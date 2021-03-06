package main

import (
	"encoding/json"
	"fmt"
	"net"
	"net/http"
	"time"
)

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

// ArticleHandler is a function handler
func ArticleHandler(w http.ResponseWriter, r *http.Request) {
	// mux.Vars returns all path parameters as a map
	vars := mux.Vars(r)
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Category is: %v\n", vars["category"])
	fmt.Fprintf(w, "ID is: %v\n", vars["id"])
}

func traceIP() {
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
