package main

import (
	"fmt"
	"log"
)

// START OMIT
type badValueError struct {
	key   string
	value interface{}
}

func (bv badValueError) Error() string {
	return fmt.Sprintf("bad value %v for key %s", bv.value, bv.key)
}

func main() {
	log.Println(badValueError{key: "k", value: 1})
}

// END OMIT
