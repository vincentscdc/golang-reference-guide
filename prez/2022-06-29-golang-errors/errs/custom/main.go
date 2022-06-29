package main

import (
	"errors"
	"log"
)

type Error struct {
	Path string
	User string
}

func (e *Error) Error() string {
	return "yes"
}

func (e *Error) Is(target error) bool {
	t, ok := target.(*Error)
	if !ok {
		return false
	}

	return e.User == t.User
}

func (e *Error) As(target any) bool {
	t, ok := target.(*Error)
	if !ok {
		return false
	}

	return e.Path == t.Path
}

func main() {
	err1 := &Error{Path: "123", User: "momo"}
	err2 := &Error{Path: "456", User: "vincent"}
	errRef := &Error{Path: "789", User: "momo"}

	if errors.Is(err1, errRef) {
		log.Printf("err1 is same as &Error{User: vincent}")
	}

	if errors.Is(err2, errRef) {
		log.Printf("err2 is same as &Error{User: vincent}")
	}

	if errors.As(err1, &errRef) {
		log.Printf("err1 as &Error{User: vincent}")
	}
}
