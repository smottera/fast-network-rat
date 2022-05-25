//versions of go, gRPC, protobuffers need to be compatible
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

func timeTest() {
	now := time.Now()

	fmt.Println("locay :", now.Location())
	fmt.Println("time.Now(): ", now)
	fmt.Println("now dot nanosecond: ", now.Nanosecond())
	fmt.Println("computer clock in seconds: ", now.Unix())
	fmt.Println("computer clock in nanoseconds: ", now.UnixNano())

	diff := time.Now().Sub(now)

	fmt.Println("diff :", diff, time.Now().After(now))
	//you can also format time in any fashion
}

func randomNumbers() {
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)

	fmt.Println(r1.Float32())
	fmt.Println(r1.Float64())
	fmt.Println(r1.Intn(1000000000000))
}

func shaHashes() {
	s := "Yo Pierre you wanna come out and play?!"
	h := sha256.New()
	h.Write([]byte(s))
	byteSlice := h.Sum(nil)

	fmt.Println(s)
	fmt.Printf("%x\n", byteSlice)

	s2 := "Yo Pierre you wanna come out and play?!"
	h2 := sha512.New()
	h2.Write([]byte(s2))
	byteSlice2 := h2.Sum(nil)

	fmt.Println(s2)
	fmt.Printf("%x\n", byteSlice2)
}

func base64Encoding() {
	data := "abc123!?$*&()'-=@~"
	sEnc := b64.StdEncoding.EncodeToString([]byte(data))
	fmt.Println(sEnc)

	sDec, _ := b64.StdEncoding.DecodeString(sEnc)
	fmt.Println(string(sDec))
	fmt.Println()

	uEnc := b64.URLEncoding.EncodeToString([]byte(data))
	fmt.Println(uEnc)
	uDec, _ := b64.URLEncoding.DecodeString(uEnc)
	fmt.Println(string(uDec))

}

func main() {
	fmt.Println("YO ... Line one: ")
	//timeTest()
	//randomNumbers()
	//shaHashes()

	//base64Encoding()

}
