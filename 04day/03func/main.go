package main

import "fmt"

func test(x, y int) {
	fmt.Println(x, y)
	if x > 0 {
		z := 100
		fmt.Println(z)
	}

}
func main() {
	test(10,20)
}
