package main

import (
	"io"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	f, err := os.Open("/dev/zero")
	if err != nil {
		log.Panic(err)
	}
	defer f.Close()
	r1 := io.LimitReader(f, 100)
	r2 := io.LimitReader(f, 200)
	res, err := ioutil.ReadAll(r1) //nolint
	if err != nil {
		log.Panic(err)
	}
	log.Printf("b=%d \n", len(res))
	res, err = ioutil.ReadAll(r2) // want `ioutil.ReadAll`
	if err != nil {
		log.Panic(err)
	}
	log.Printf("b=%d \n", len(res))
}
