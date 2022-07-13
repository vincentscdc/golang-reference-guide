package main

import (
	"net/http"
)

const concatsCount = 1000000

func Leaky() http.HandlerFunc {
	return http.HandlerFunc(func(respW http.ResponseWriter, r *http.Request) {
		done := leakyFunction()
		<-done
	})
}

// no closing of pandas channel
func leakyFunction() chan struct{} {
	done := make(chan struct{})

	pandas := make(chan []string)
	go func() {
		s := make([]string, 1)
		for i := 0; i < concatsCount; i++ {
			s = append(s, "magical pandas")
		}

		done <- struct{}{}

		pandas <- s
	}()

	return done
}

func NonLeaky() http.HandlerFunc {
	return http.HandlerFunc(func(respW http.ResponseWriter, r *http.Request) {
		done := nonleakyFunction()
		<-done
	})
}

func NonLeakyOptim() http.HandlerFunc {
	return http.HandlerFunc(func(respW http.ResponseWriter, r *http.Request) {
		done := nonleakyoptimFunction()
		<-done
	})
}

func nonleakyFunction() chan struct{} {
	done := make(chan struct{})

	pandas := make(chan []string)
	go func() {
		s := make([]string, 1)
		for i := 0; i < concatsCount; i++ {
			s = append(s, "magical pandas")
		}

		done <- struct{}{}
		close(pandas)
	}()

	return done
}

func nonleakyoptimFunction() chan struct{} {
	done := make(chan struct{})

	pandas := make(chan []string)
	go func() {
		s := make([]string, 0, concatsCount)
		for i := 0; i < concatsCount; i++ {
			s = append(s, "magical pandas")
		}

		done <- struct{}{}
		close(pandas)
	}()

	return done
}
