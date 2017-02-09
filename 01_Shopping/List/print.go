package list

import "fmt"

var x = 0

func increment() int {
	x++
	return x
}

func PrintVar() { // func to print the list
	fmt.Println(Name)
	fmt.Println(increment(), "%v \n", a)
	fmt.Println(increment(), "%v \n", b)
	fmt.Println(increment(), "%v \n", c)
	fmt.Println(increment(), "%v \n", d)
	fmt.Println(increment(), "%v \n", e)
	fmt.Println(increment(), "%v \n", f)
	fmt.Println(increment(), "%v \n", g)
	fmt.Println(increment(), "%v \n", h)
	fmt.Println(increment(), "%v \n", i)
	fmt.Println(increment(), "%v \n", j)
	fmt.Println(increment(), "%v \n", k)

}
