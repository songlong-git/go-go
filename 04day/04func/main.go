package main

import "fmt"

func main() {
	add := func(x, y int) {
		fmt.Println(x + y)
	}
	add(10, 20)
	func(x,y int){
		fmt.Println(x + y)
	}(10,20)
}
