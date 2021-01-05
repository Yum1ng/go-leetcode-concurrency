package main

import "testing"

func TestPrintInOrder(t *testing.T) {
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
