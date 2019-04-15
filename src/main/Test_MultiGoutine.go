package main

import (
	"fmt"
	"runtime"
	"time"
)

func sum(a []int, c chan int) {
	total := 0
	for _, v := range a {
		total += v
	}
	c <- total  // send total to c
}

func main() {
	runtime.GOMAXPROCS(1)
	//runtime.GOMAXPROCS(2)
	a := []int{7, 2, 8, -9, 4, 0}

	c := make(chan int)
	go sum(a[:len(a)/2], c)

	go sum(a[len(a)/2:], c)

	time.Sleep(1)
	x, y := <-c, <-c  // receive from c

	fmt.Println(x, y, x + y)
}
