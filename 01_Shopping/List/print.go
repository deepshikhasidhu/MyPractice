package list

import "fmt"

var x = 0

func increment() int {
	x++
	return x
}

func PrintVar() {
	fmt.Println(increment(), ".", a)
	fmt.Println(increment(), b)
	fmt.Println(increment(), c)
	fmt.Println(increment(), d)
	fmt.Println(increment(), e)
	fmt.Println(increment(), f)
	fmt.Println(increment(), g)
	fmt.Println(increment(), h)
	fmt.Println(increment(), i)
	fmt.Println(increment(), j)
	fmt.Println(increment(), k)

}
