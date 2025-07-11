package main

import (
	"fmt"
	"log"
	"sync"
	"time"
)

func worker(jobs chan int, results chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := range jobs {
		fmt.Printf("START: worker job %d\n", i)
		time.Sleep(time.Second)
		fmt.Printf("END: worker job %d\n", i)
		results <- i*10
	}
}
func main() {
	log.Print("Data Migration")

	jobs := make(chan int)
	results := make(chan int)

	var wg sync.WaitGroup
	go func() {
		for r := range results {
			fmt.Printf("Result  = %d\n", r)
		}
	}()

	for j := 1; j <= 7; j += 1 {
		wg.Add(1)
		go worker(jobs, results, &wg)
	}


	for i := 10; i <= 20; i+=1 {
		jobs <- i
	}
	close(jobs)
	wg.Wait()
}
