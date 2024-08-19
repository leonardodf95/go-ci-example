package main

import "fmt"

func main() {
	x := 1
	y := 2
	fmt.Println(fmt.Sprintf("%d + %d = %d", x, y, sum(x, y)))
}

func sum(x, y int) int {
	return x + y
}
