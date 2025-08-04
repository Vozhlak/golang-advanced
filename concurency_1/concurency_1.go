package concurency_1

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func generator(ch chan int) {
	defer close(ch)
	for i := 0; i < 10; i++ {
		r := rand.New(rand.NewSource(time.Now().UnixNano()))
		num := r.Intn(101)
		ch <- num
	}
}

func squarer(input chan int, output chan int) {
	for num := range input {
		result := num * num

		output <- result
	}
}

func Run() {
	ch1 := make(chan int)
	ch2 := make(chan int)

	result := make([]int, 0, 10)

	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()
		generator(ch1)
	}()
	go func() {
		defer wg.Done()
		squarer(ch1, ch2)
	}()

	go func() {
		defer close(ch2)
		wg.Wait()
	}()

	for num := range ch2 {
		result = append(result, num)
	}

	fmt.Println(result)
}
