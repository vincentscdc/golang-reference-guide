package main

import (
	"fmt"
	"log"
	"math/rand"
	"sync"
	"time"
)

type Result struct {
	ID    int
	Error error
}

func worker(id int) (int, error) {
	log.Printf("Worker %d starting\n", id)

	rnd := rand.New(rand.NewSource(time.Now().UnixNano()))

	r := rnd.Intn(100)

	time.Sleep(time.Duration(r) * time.Millisecond)

	log.Printf("Worker %d done\n", id)

	if r > 50 {
		return 0, fmt.Errorf("error in worker %d", id)
	}

	return r, nil
}

// START OMIT
func main() {
	cRes := make(chan Result)

	go func() {
		for v := range cRes {
			if v.Error != nil {
				log.Printf("error: %v\n", v)
			} else {
				log.Printf("ok: %v\n", v)
			}
		}
	}()

	var waitGroup sync.WaitGroup

	for i := 1; i <= 5; i++ {
		waitGroup.Add(1)

		i := i

		go func() {
			defer waitGroup.Done()

			r, err := worker(i)

			cRes <- Result{ID: r, Error: err}
		}()
	}

	waitGroup.Wait()

	close(cRes)
}

// END OMIT
