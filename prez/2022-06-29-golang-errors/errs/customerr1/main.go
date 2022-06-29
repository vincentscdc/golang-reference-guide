package main

import (
	"log"
)

// START OMIT
type marshalError string

func (me marshalError) Error() string {
	return string(me)
}

const ErrBadInt = marshalError("bad int")

func main() {
	log.Println(ErrBadInt)
}

// END OMIT
