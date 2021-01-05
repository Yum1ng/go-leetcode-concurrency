package main

import (
	"fmt"
	"time"
)

func main() {
	counter := Counter{
		oneReady: make(chan int, 1),
		twoReady: make(chan int, 1),
	}
	count([]int{1, 2, 3}, counter)
	count([]int{1, 3, 2}, counter)
	count([]int{2, 1, 3}, counter)
	count([]int{2, 3, 1}, counter)
	count([]int{3, 1, 2}, counter)
	count([]int{3, 2, 1}, counter)
}

func count(input []int, counter Counter) {
	if len(input) != 3 {
		panic("invalid input")
	}
	fmt.Println("Begin")
	for _, i := range input {
		switch i {
		case 1:
			counter.First()
		case 2:
			counter.Second()
		case 3:
			counter.Third()
		default:
			panic("invalid input, must be 1, 2 or 3")
		}
	}
	time.Sleep(time.Second)
}

type Counter struct {
	oneReady chan int
	twoReady chan int
}

func (c Counter) First() {
	fmt.Println(1)
	c.oneReady <- 1
}

func (c Counter) Second() {
	go func() {
		for {
			select {
			case <-c.oneReady:
				fmt.Println(2)
				c.twoReady <- 2
				return
			default:
				// do nothing
			}
		}
	}()
}

func (c Counter) Third() {
	go func() {
		for {
			select {
			case <-c.twoReady:
				fmt.Println(3)
				return
			default:
				// do nothing
			}
		}
	}()
}
