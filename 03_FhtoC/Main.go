package main

import "fmt"

func main() {
	var fh float64
	str1 := "Enter Temperature: "
	fmt.Println(str1)
	fmt.Scan(&fh)
	celsius := (fh - 32) * 5 / 9
	fmt.Println(fh, "Fahrenheit is", celsius, "celsius")

}
