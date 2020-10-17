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
	r := io.LimitReader(f, 100)
	res, err := ioutil.ReadAll(r)
	if err != nil {
		log.Panic(err)
	}
	log.Printf("b=%d \n", len(res))
}
