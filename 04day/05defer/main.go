package main

import "fmt"

func f1() int {
	x := 5
	defer func() {
		x++
	}()
	return x
}
func main() {
	fmt.Println(f1())
}
