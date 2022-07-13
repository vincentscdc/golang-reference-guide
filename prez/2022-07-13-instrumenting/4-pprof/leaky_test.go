package main

import (
	"flag"
	"os"
	"testing"

	"go.uber.org/goleak"
)

func TestMain(m *testing.M) {
	leak := flag.Bool("leak", false, "use leak detector")
	flag.Parse()

	if *leak {
		goleak.VerifyTestMain(m)

		return
	}

	code := m.Run()

	os.Exit(code)
}

func Benchmark_leakyFunction(b *testing.B) {
	for i := 0; i < b.N; i++ {
		done := leakyFunction()
		<-done
	}
}

func Benchmark_nonleakyFunction(b *testing.B) {
	for i := 0; i < b.N; i++ {
		done := nonleakyFunction()
		<-done
	}
}

func Benchmark_nonleakyoptimFunction(b *testing.B) {
	for i := 0; i < b.N; i++ {
		done := nonleakyoptimFunction()
		<-done
	}
}

func Test_leakyFunction(t *testing.T) {
	t.Run("happy path", func(t *testing.T) {
		done := leakyFunction()
		<-done
	})
}

func Test_nonleakyFunction(t *testing.T) {
	t.Run("happy path", func(t *testing.T) {
		done := nonleakyFunction()
		<-done
	})
}

func Test_nonleakyoptimFunction(t *testing.T) {
	t.Run("happy path", func(t *testing.T) {
		done := nonleakyoptimFunction()
		<-done
	})
}
