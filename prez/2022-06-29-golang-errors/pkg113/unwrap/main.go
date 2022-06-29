package main

import (
	"errors"
	"fmt"
)

func main() {
	err := errors.New("error wrapped")
	errW := fmt.Errorf("errW: %w", err)

	fmt.Printf("%v\n", err)
	fmt.Printf("%v\n", errW)
	fmt.Printf("%v\n", errors.Unwrap(errW))
	fmt.Printf("%v\n", errors.Unwrap(err))
}
