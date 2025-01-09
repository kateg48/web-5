package main

import "fmt"

func removeDuplicates(inputStream, outputStream chan string) {
	var tmp string = ""
	for v := range inputStream {
		if v != tmp {
			outputStream <- v
		}
		tmp = v
	}
	close(outputStream)
}

func main() {
	inputStream := make(chan string)
	outputStream := make(chan string)
	var s string = "фыфыфыфыфыфыфы"
	go removeDuplicates(inputStream, outputStream)
	go func() {
		for v := range outputStream {
			fmt.Print(v)
		}
	}()

	for _, a := range s {
		inputStream <- string(a)
	}

}
