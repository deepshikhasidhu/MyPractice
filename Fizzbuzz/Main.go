package main

import "fmt"

func main() {
	for i := 1; i <= 100; i++ {
		a := i % 3
		b := i % 5
		if a == 0 && b != 0 {
			fmt.Println("Fizz")
		} else if a != 0 && b == 0 {
			fmt.Println("Buzz")
		} else if a == 0 && b == 0 {
			fmt.Println("FizzBuzz")
		} else {
			fmt.Println(i)
		}
	}
}
