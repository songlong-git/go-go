package main

import "fmt"

func sum(x int, y int) (ret int) {
	return x + y
}
func main() {
	r := sum(10, 20)
	fmt.Println(r)
}
