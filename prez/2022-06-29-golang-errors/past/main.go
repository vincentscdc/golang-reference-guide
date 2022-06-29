package main

import (
	"errors"
	"fmt"
)

// START OMIT
var (
	err1 = errors.New("error1")
	err2 = errors.New("error2")
)

func this1() error {
	err := this2()
	if err != nil {
		return fmt.Errorf("this1: %v", err)
	}

	return err1
}

func this2() error {
	return err2
}

func main() {
	err := this1()
	if err != nil {
		fmt.Println(err)
	}
}

// END OMIT
