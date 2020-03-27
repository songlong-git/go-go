package main

import "fmt"

func intSum(x, y int) int {
	return x + y
}

func sayHello() {
	fmt.Println("Hello 龙帝")
}

func calc(x, y int) (sum , sub int) {
	sum = x + y
	sub = x - y
	return sum, sub
}

func main() {
	//sayHello()
	//a, b := calc(10, 20)
a,b :=calc(10, 20)
	fmt.Printf("c = %d,d =%d\n",a,b)
}
