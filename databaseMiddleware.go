package main

import (
	"crypto/sha256"
	"crypto/sha512"
	b64 "encoding/base64"
	"fmt"
	"math/rand"
	"time"
	//net or https for another module.
	//filepath, directories, reading/writing
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	fmt.Println("YO ... Line one: ")
	//timeTest()
	//randomNumbers()
	//shaHashes()

	//base64Encoding()
	collective := 69
	fmt.Println(collective, yaw)
}
