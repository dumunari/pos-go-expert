package main

func main() {
	// ch := make(chan string, 1)
	ch := make(chan string, 2)
	ch <- "Hello"
	ch <- "World"

	println(<-ch)
	// time.Sleep(2 * time.Second)
	println(<-ch)
}
