package main

import (
	"fmt"
	"log"
)

func main() {
	log.Println(fmt.Errorf("bad int %d", 3))
}
