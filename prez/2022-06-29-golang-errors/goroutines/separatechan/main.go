package main

import (
	"fmt"
	"log"
	"math/rand"
	"sync"
	"time"
)

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
	cI := make(chan int)
	cErr := make(chan error)

	go func() {
		for v := range cErr {
			log.Printf("error: %v\n", v)
		}
	}()

	go func() {
		for v := range cI {
			log.Printf("ok: %v\n", v)
		}
	}()

	var waitGroup sync.WaitGroup

	for i := 1; i <= 5; i++ {
		waitGroup.Add(1)

		i := i

		go func() {
			defer waitGroup.Done()

			r, err := worker(i)
			if err != nil {
				cErr <- err

				return
			}

			cI <- r
		}()
	}

	waitGroup.Wait()

	close(cI)
	close(cErr)
}

// END OMIT
