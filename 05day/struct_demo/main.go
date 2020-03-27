package main
import "fmt"
type person struct {
	name   string
	age    int
	gender string
	hoppy  []string
}

func main() {
	var p person
	p.name = "宋龙"
	p.age = 18
	p.gender = "男"
	p.hoppy =[]string{"篮球","足球"}
	fmt.Println(p)
	fmt.Printf("%T\n",p)
	fmt.Println(p.name)

}
