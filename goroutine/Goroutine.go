package goroutine

import (
	"fmt"
	"time"
)

func GoChannel() {
	c1 := make(chan string)
	c2 := make(chan string)

	go func() {
		for {
			c1 <- "Every 500ms"
			time.Sleep(time.Millisecond * 500)
		}
	}()

	go func() {
		for {
			c2 <- "Every Two Seconds"
			time.Sleep(time.Second * 2)
		}
	}()

	count := 0
	for {
		select {
		case msg1 := <-c1:
			fmt.Println(msg1)
		case msg2 := <-c2:
			fmt.Println(msg2)
		}

		if count > 10 {
			break
		}
		count++

	}
}

func GoWorker() {
	jobs := make(chan int, 10)
	results := make(chan int, 10)

	go worker(jobs, results)
	go worker(jobs, results)
	go worker(jobs, results)
	go worker(jobs, results)

	// sending i to jobs
	for i := 0; i < 10; i++ {
		jobs <- i
	}
	close(jobs)

	// get results 
	for j := 0; j < 10; j++ {
		fmt.Println(<-results)
	}
}

func worker(jobs <-chan int, results chan<- int) {
	for n := range jobs {
		results <- fib(n)
	}
}

func fib(n int) int {
	if n <= 1 {
		return n
	}

	return fib(n-1) + fib(n-2)
}
