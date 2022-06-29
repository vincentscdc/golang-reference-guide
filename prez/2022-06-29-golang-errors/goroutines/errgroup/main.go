package main

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"time"

	"golang.org/x/sync/errgroup"
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
	rand.NewSource(time.Now().UnixNano())

	cI := make(chan int)
	ints := make([]int, 0)
	go func() {
		for v := range cI {
			ints = append(ints, v)
		}
	}()

	errs, _ := errgroup.WithContext(context.Background())

	for i := 1; i <= 5; i++ {
		i := i

		errs.Go(func() error {
			r, err := worker(i)
			if err == nil {
				cI <- r
			}

			return err
		})
	}

	if err := errs.Wait(); err != nil {
		log.Println(err)
	}

	close(cI)

	for _, res := range ints {
		log.Printf("%v\n", res)
	}
}

// END OMIT
