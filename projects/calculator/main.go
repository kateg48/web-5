package main

import "fmt"

func main() {
	c1 := make(chan int)
	c2 := make(chan int)
	s := make(chan struct{})

	var r int = 1000000

	go func(r *int) {
		*r = <-calculator(c1, c2, s)
	}(&r)

	//c1 <- 3
	//c2 <- 2
	s <- struct{}{}
	fmt.Print(r)
}

func calculator(firstChan <-chan int, secondChan <-chan int,
	stopChan <-chan struct{}) <-chan int {

	outp := make(chan int)
	go func(outp chan int) {
		defer close(outp)

		select {
		case a := <-firstChan:
			outp <- a * a
		case a := <-secondChan:
			outp <- a * 3
		case <-stopChan:
			return
		}

	}(outp)
	return outp
}
