package main

import "fmt"

func main() {
	s := []int{}
	fmt.Printf("len=%d cap =%d %v\n", len(s), cap(s), s)

	x := []int{10, 20, 30, 40, 50, 60}
	fmt.Printf("len=%d cap =%d %v\n", len(x), cap(x), x)

	fmt.Printf("len=%d cap =%d %v\n", len(x[:0]), cap(x[:0]), x[:0])

	fmt.Printf("len=%d cap =%d %v\n", len(x[:3]), cap(x[:3]), x[:3])

	// dropa o capacity se ignora do começo
	fmt.Printf("len=%d cap =%d %v\n", len(x[3:]), cap(x[3:]), x[3:])

	// dobra o capacity
	x = append(x, 70)
	fmt.Printf("len=%d cap =%d %v\n", len(x), cap(x), x)
}
