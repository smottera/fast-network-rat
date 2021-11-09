package main

import (
	"fmt"
	"net"
	"time"
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
