package main

import (
	"errors"
	"log"
)

var errBadInt = errors.New("bad int")

func main() {
	log.Println(errBadInt)
}
