package main

import "fmt"

//构造体
type dog struct {
	name string
}

// 构造函数
func newDog(name string) dog {
	return dog{
		name: name,
	}
}
func (d dog) wang() {
	fmt.Printf("%s:汪汪汪", d.name)
}
func main() {
	d1 := newDog("周琳")
	d1.wang()
 
}
