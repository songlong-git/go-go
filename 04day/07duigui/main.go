package main

import "fmt"

func f(n uint64) uint64 {
	if n <= 1 {
		return 1
	}
	return n * f(n-1)
}
func main() {
	ret := f(7)
	fmt.Println(ret)
}
